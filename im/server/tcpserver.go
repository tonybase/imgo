package server

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"imgo/im/common"
	"imgo/im/model"
	"imgo/im/util"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

/*
服务端结构体
*/
type Server struct {
	listener    net.Listener        // 服务端监听器 监听xx端口
	clients     common.ClientTable  // 客户端列表 抽象出来单独维护和入参 更方便管理连接
	joinsniffer chan net.Conn       // 访问连接嗅探器 触发创建客户端连接处理方法
	quitsniffer chan *common.Client // 连接退出嗅探器 触发连接退出处理方法
	insniffer   common.InMessage    // 接收消息嗅探器 触发接收消息处理方法 对应客户端中in属性
}

var ClientMaps common.ClientTable

/*
 IM服务启动方法
*/
func StartIMServer(config util.IMConfig) {
	log.Println("服务端启动中...")
	//初始化服务端
	server := &Server{
		clients:     make(common.ClientTable, model.Config.MaxClients),
		joinsniffer: make(chan net.Conn),
		quitsniffer: make(chan *common.Client),
		insniffer:   make(common.InMessage),
	}
	ClientMaps = server.clients
	// 添加关闭勾子，当关闭服务器时执行
	server.interruptHandler()
	// 启动监听方法(包含各类嗅探器)
	server.listen()
	// 启动服务端端口监听(等待连接)
	server.start()
}

/*
 监听方法
*/
func (this *Server) listen() {
	go func() {
		for {
			select {
			// 接收到了消息
			case message := <-this.insniffer:
				this.receivedHandler(message)
			// 新来了一个连接
			case conn := <-this.joinsniffer:
				this.joinHandler(conn)
			// 退出了一个连接
			case client := <-this.quitsniffer:
				this.quitHandler(client)
			}
		}
	}()
}

/*
 新客户端请求处理方法
*/
func (this *Server) joinHandler(conn net.Conn) {
	//获取UUID作为客户端的key
	key := uuid.New()
	// 创建一个客户端
	client := common.CreateClient(key, conn)
	//给客户端指定key
	this.clients[key] = client
//	log.Printf("新客户端Key:[%s] online:%d", client.Key, len(ClientMaps))
	// 开启协程不断地接收消息
	go func() {
		for {
			// 客户端读取消息
			msg := <-client.In
			// 消息交给嗅探器 触发对应的处理方法
			this.insniffer <- msg
		}
	}()
	// 开启协程一直等待断开
	go func() {
		for {
			//客户端接收断开请求
			conn := <-client.Quit
//			log.Printf("客户端:[%s]退出", client.Key)
			//请求交给嗅探器 触发对应的处理方法
			this.quitsniffer <- conn
		}
	}()
	// 返回客户端的唯一标识
	data := make(map[string]interface{})
	data["key"] = key
	client.PutOut(common.NewIMResponseData(util.SetData("conn", data), common.GET_KEY_RETURN))
}

/*
 客户端退出处理方法
*/
func (this *Server) quitHandler(client *common.Client) {
	if client != nil {
		// 通知在线的好友，我离线了
		if client.Login != nil {
			// 判断要求改变的状态和当前该用户的状态是否一致
			model.DeleteConnByKey(client.Key)
			count, _ := model.CountConnByUserId(client.Login.UserId)
			// 如果没有这用户的连接，同时更新用户状态为离线
			if count == 0 {
				model.UpdateUserStatus(client.Login.UserId, "0")
			}
			keys, err := model.GetBuddiesKeyById(client.Login.UserId)
			if err == nil {
				for i := 0; i < len(keys); i++ {
					//给对应的连接推送好友状态变化的通知
					data := make(map[string]string)
					data["id"] = client.Login.UserId
					data["state"] = "0"
					if (this.clients[keys[i]] != nil) {
						this.clients[keys[i]].PutOut(common.NewIMResponseData(util.SetData("user", data), common.PUSH_STATUS_CHANGE))
					}
				}
			}
		}

		// 调用客户端关闭方法
		client.Close()
		delete(this.clients, client.Key)

//		log.Printf("客户端退出: %s online:%d", client.Key, len(ClientMaps))
	}
}

/*
 接收消息处理方法
*/
func (this *Server) receivedHandler(request common.IMRequest) {
//	log.Println("开始读取数据")
//	log.Println("读取的数据为", request)

	// 获取请求的客户端
	client := request.Client
	// 获取请求数据
	reqData := request.Data
//	log.Printf("客户端:[%s]发送命令:[%s]消息内容:[%s]", client.Key, request.Command, request.Data)

	// 未登录业务处理部分
	switch request.Command {
	case common.GET_CONN:
		token := reqData["user"]["token"]
		if token == "" {
			client.PutOut(common.NewIMResponseSimple(301, "用户令牌不能为空!", common.GET_CONN_RETURN))
			return
		}
		// 校验用户是否登录，把Login数据放在client当中
		login, err := model.GetLoginByToken(token)
		if err != nil {
			client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.GET_CONN_RETURN))
			return
		}
		client.Login = login
//		log.Printf("登录比较：token=%s Login=%s", token, client.Login)
		if !strings.EqualFold(client.Login.Token, token) {
			client.PutOut(common.NewIMResponseSimple(302, "该用户令牌无效!", common.GET_CONN_RETURN))
			return
		}
		if client.Login.Id != "" {
			// 更新在线状态，如果现在已经是在线，然后再设置在线，影响行还是为0
			_, err := model.UpdateUserStatus(client.Login.UserId, "1")
			if err != nil {
				client.PutOut(common.NewIMResponseSimple(304, "设置用户状态失败!", common.GET_CONN_RETURN))
				return
			}
			// 创建或者更新连接信息 这个error不能handler，当没有数据时会为 sql: no rows in result set
			conn, _ := model.GetConnByToken(token)
			if conn != nil {
				_, err := model.UpdateConnByToken(client.Key, client.Login.UserId, token)
				if err != nil {
					client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.GET_CONN_RETURN))
					return
				}
			} else {
				_, err := model.AddConn(client.Key, client.Login.UserId, token)
				if err != nil {
					client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.GET_CONN_RETURN))
					return
				}
			}
			data := make(map[string]interface{})
			data["status"] = 1
			client.PutOut(common.NewIMResponseData(util.SetData("conn", data), common.GET_CONN_RETURN))
			// 通知在线的好友，我上线了
			keys, err := model.GetBuddiesKeyById(client.Login.UserId)
			if err != nil {
				client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.SEND_STATUS_CHANGE))
				return
			}
			for i := 0; i < len(keys); i++ {
				//给对应的连接推送好友状态变化的通知
				data := make(map[string]string)
				data["id"] = client.Login.UserId
				data["state"] = "1"
				if (this.clients[keys[i]] != nil) {
					this.clients[keys[i]].PutOut(common.NewIMResponseData(util.SetData("user", data), common.PUSH_STATUS_CHANGE))
				}
			}
			return
		} else {
			client.PutOut(common.NewIMResponseSimple(303, "用户未登录!", common.GET_CONN_RETURN))
			return
		}
	}
	// 校验连接是已经授权
	if client.Login == nil {
		client.PutOut(common.NewIMResponseSimple(401, "用户未登录!", common.UNAUTHORIZED))
		return
	}
	// 已经登录业务逻辑部分
	switch request.Command {
	case common.GET_BUDDY_LIST:
		// 获取好友分组列表
		categories, err := model.GetCategoriesByUserId(client.Login.UserId)
		if err != nil {
			client.PutOut(common.NewIMResponseSimple(301, "获取好友分类错误!", common.GET_BUDDY_LIST_RETURN))
			return
		}
		categories, err = model.GetBuddiesByCategories(categories)
		if err != nil {
			client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.GET_BUDDY_LIST_RETURN))
			return
		}
		client.PutOut(common.NewIMResponseData(util.SetData("categories", categories), common.GET_BUDDY_LIST_RETURN))
		//初始化好友列表之后 检查该用户有没有未读的好友请求 并推送给用户
		buddyRequests, err := model.GetBuddyRequestsByReceiver(client.Login.UserId)
		if(err != nil) {
			client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.GET_BUDDY_LIST_RETURN))
		}
		if len(buddyRequests) > 0 {
			for _, buddyRequest := range buddyRequests {
				user, _ := model.GetUserById(buddyRequest.Sender)
				data := make(map[string]interface{})
				data["id"] = user.Id
				data["nick"] = user.Nick
				data["status"] = user.Status
				data["sign"] = user.Sign
				data["avatar"] = user.Avatar
				data["buddyRequestId"] = buddyRequest.Id
				client.PutOut(common.NewIMResponseData(util.SetData("user", data), common.PUSH_BUDDY_REQUEST))
			}
		}
	case common.CREATE_SESSION:
		// 创建会话  //{"command":"CREATE_SESSION","data":{"session":{"sender":"xxx","receiver":"xxx","token":"xxxx"}}}
		sender := reqData["session"]["sender"]
		receiver := reqData["session"]["receiver"]

		if sender == "" {
			client.PutOut(common.NewIMResponseSimple(301, "发送者不能为空!", common.CREATE_SESSION_RETURN))
			return
		}
		if receiver == "" {
			client.PutOut(common.NewIMResponseSimple(302, "接收者不能为空!", common.CREATE_SESSION_RETURN))
			return
		}
		conversationId := model.GetSession(sender, receiver).Id
		if conversationId == "" {
			conversationId = model.AddSession(sender, receiver)
		}
		if conversationId == "" {
			client.PutOut(common.NewIMResponseSimple(303, "创建会话失败", common.GET_CONN_RETURN))
			return
		} else {
			data := make(map[string]string)
			data["ticket"] = conversationId
			data["receiver"] = receiver
			client.PutOut(common.NewIMResponseData(util.SetData("session", data), common.CREATE_SESSION_RETURN))
		}

	case common.SEND_MSG:
		ticket := reqData["message"]["ticket"]
		content := reqData["message"]["content"]

		if ticket == "" {
			client.PutOut(common.NewIMResponseSimple(301, "Ticket不能为空!", common.SEND_MSG_RETURN))
			return
		}
		if content == "" {
			client.PutOut(common.NewIMResponseSimple(302, "消息内容不能为空!", common.SEND_MSG_RETURN))
			return
		}
		conversion, err := model.GetSessionById(ticket)
		if err != nil {
			client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.SEND_MSG_RETURN))
			return
		}
		if conversion.Id != "" {
			isSent := false
			keys, err := model.GetReceiverKeyByTicket(ticket)
			if err != nil {
				client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.SEND_MSG_RETURN))
				return
			}
			for _, key := range keys {
				if this.clients[key] == nil {
					// client.PutOut(common.NewIMResponseSimple(402, "对方还未登录!", common.SEND_MSG_RETURN))
					continue
				}
				// 把消息转发给接收者
				data := make(map[string]string)
				data["sender"] = client.Login.UserId
				data["ticket"] = ticket
				data["content"] = content
				log.Println("开始转发给:", key)
				this.clients[key].PutOut(common.NewIMResponseData(util.SetData("message", data), common.PUSH_MSG))
				isSent = true
			}
			if !isSent {
				client.PutOut(common.NewIMResponseSimple(304, "对方不在线!", common.GET_CONN_RETURN))
				return
			}
		} else {
			client.PutOut(common.NewIMResponseSimple(303, "会话已关闭!", common.SEND_MSG_RETURN))
			return
		}

	case common.SEND_STATUS_CHANGE:
		status := reqData["user"]["status"]
		if status == "" {
			client.PutOut(common.NewIMResponseSimple(301, "状态不能为空!", common.SEND_STATUS_CHANGE))
			return
		}
		user, err := model.GetUserByToken(client.Login.Token)
		if err != nil {
			client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.SEND_STATUS_CHANGE))
			return
		}
		//判断用户的合法性
		if user.Id == "" {
			//判断要求改变的状态和当前该用户的状态是否一致
			if strings.EqualFold(user.Status, status) {
				//FIXME 此处不做如果状态是离线就删除用户连接的操作,状态改变认为是客户端手动操作或者网络异常
				_, err := model.UpdateUserStatus(user.Id, status)
				if err != nil {
					client.PutOut(common.NewIMResponseSimple(304, err.Error(), common.SEND_STATUS_CHANGE))
					return
				}
				keys, err := model.GetBuddiesKeyById(user.Id)
				if err != nil {
					client.PutOut(common.NewIMResponseSimple(300, err.Error(), common.SEND_STATUS_CHANGE))
					return
				}
				for i := 0; i < len(keys); i++ {
					//给对应的连接推送好友状态变化的通知
					data := make(map[string]string)
					data["id"] = user.Id
					data["state"] = reqData["user"]["status"]
					this.clients[keys[i]].PutOut(common.NewIMResponseData(util.SetData("user", data), common.PUSH_STATUS_CHANGE))
				}
			} else {
				client.PutOut(common.NewIMResponseSimple(303, "请退出重新登录!", common.SEND_STATUS_CHANGE))
				return
			}

		} else {
			client.PutOut(common.NewIMResponseSimple(302, "Token不合法!", common.SEND_STATUS_CHANGE))
			return
		}
	case common.LOGOUT_REQUEST:
		client.Quiting()
	case common.SEND_BUDDY_REQUEST:
		receiver := reqData["buddyRequest"]["receiver"]
		//判断接收者是不是在线
		user, _ := model.GetUserById(receiver)
		if user == nil || user.Status == "0" { //不在线 记录到好友请求表中
			id, _ := model.AddBuddyRequest(reqData["buddyRequest"]["sender"], reqData["buddyRequest"]["senderCateId"], receiver)
			if id != nil {

			}
		} else { //在线直接推送给接收者
			conn, _ := model.GetConnByUserId(receiver)
			data := make(map[string]string)
			data["sender"] = reqData["buddyRequest"]["sender"]
			data["senderCateId"] = reqData["buddyRequest"]["senderCateId"]
			data["receiver"] = reqData["buddyRequest"]["receiver"]
			this.clients[conn.Key].PutOut(common.NewIMResponseData(util.SetData("buddyRequest", data), common.PUSH_BUDDY_REQUEST))
		}

	}
}

func (this *Server) start() {
	// 设置监听地址及端口
	addr := fmt.Sprintf("0.0.0.0:%d", model.Config.IMPort)
	this.listener, _ = net.Listen("tcp", addr)
	log.Printf("开始监听服务器[%d]端口", model.Config.IMPort)
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
//		log.Printf("新连接地址为:[%s]", conn.RemoteAddr())
		this.joinsniffer <- conn
	}
}

// FIXME: need to figure out if this is the correct approach to gracefully
// terminate a server.
func (this *Server) Stop() {
	this.listener.Close()
}

// 服务端关闭时执行
func (this *Server) interruptHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("captured %v, stopping profiler and exiting..", sig)
		// 清除客户端连接
		for _, v := range this.clients {
			this.quitHandler(v)
		}
		// 退出
		os.Exit(1)
	}()
}
