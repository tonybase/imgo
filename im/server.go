package im

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"im-go/im/common"
	"im-go/im/model"
	"log"
	"net"
	"time"
)

type InMessage chan IMRequest       //读取消息通道
type OutMessage chan IMResponse     //输出消息通道
type ClientTable map[string]*Client //客户端列表
/*
服务端结构体
*/
type Server struct {
	listener    net.Listener  //服务端监听器 监听xx端口
	clients     ClientTable   //客户端列表 抽象出来单独维护和入参 更方便管理连接
	joinsniffer chan net.Conn //访问连接嗅探器 触发创建客户端连接处理方法
	quitsniffer chan *Client  //连接退出嗅探器 触发连接退出处理方法
	insniffer   InMessage     //接收消息嗅探器 触发接收消息处理方法 对应客户端中in属性
}

/*
IM服务启动方法
*/
func StartIMServer(config common.IMConfig) {
	log.Println("服务端启动中...")
	//初始化服务端
	server := &Server{
		clients:     make(ClientTable, Config.MaxClients),
		joinsniffer: make(chan net.Conn),
		quitsniffer: make(chan *Client),
		insniffer:   make(InMessage),
	}
	//启动监听方法(包含各类嗅探器)
	server.listen()
	//启动服务端端口监听(等待连接)
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
	//创建一个客户端
	client := CreateClient(key, conn)
	//给客户端指定key
	this.clients[key] = client
	log.Printf("设置新请求客户端Key为:[%s]", client.GetKey())
	//开启协程不断地接收消息
	go func() {
		for {
			//客户端读取消息
			msg := <-client.in
			//消息交给嗅探器 触发对应的处理方法
			this.insniffer <- msg
		}
	}()
	//开启协程一直等待断开
	go func() {
		for {
			//客户端接收断开请求
			conn := <-client.quit
			log.Printf("客户端:[%s]退出", client.GetKey())
			//请求交给嗅探器 触发对应的处理方法
			this.quitsniffer <- conn
		}
	}()
	//返回客户端的唯一标识
	data := make(map[string]interface{})
	data["key"] = key
	client.PutOut(NewIMResponseData(common.SetData("conn", data), GET_KEY_RETURN))
}

/*
客户端退出处理方法
*/
func (this *Server) quitHandler(client *Client) {
	if client != nil {
		//调用客户端关闭方法
		client.Close()
		delete(this.clients, client.GetKey())
	}
}

/*
接收消息处理方法
*/
func (this *Server) receivedHandler(request IMRequest) {
	//获取请求的客户端
	client := request.Client
	log.Printf("客户端:[%s]发送命令:[%s]消息内容:[%s]", client.GetKey(), request.Command, request.Data)
	//获取请求数据
	reqData := request.Data
	switch request.Command {
	case GET_CONN:
		if reqData["user"]["id"] == "" {
			client.PutOut(NewIMResponseSimple(302, "用户ID不能为空!", GET_CONN_RETURN))
			return
		}
		if reqData["user"]["token"] == "" {

			client.PutOut(NewIMResponseSimple(303, "用户令牌不能为空!", GET_CONN_RETURN))
			return
		}
		//FIXME 暂时不加，方便测试 !strings.EqualFold(reqData["user"]["key"], name)
		if reqData["user"]["key"] == "" {
			client.PutOut(NewIMResponseSimple(304, "连接的KEY错误!", GET_CONN_RETURN))
			return
		}
		log.Println("ID:", reqData["user"]["id"], "TOKEN:", reqData["user"]["token"], "KEY:", reqData["user"]["key"], "TIME:", time.Now().Format("2006-01-02 15:04:05"))
		log.Println("获取到得数据库:", Database)
		log.Println("准备开始执行")
		insertStmt, _ := Database.Prepare("insert into im_conn VALUES (?, ?, ?, ?)")
		log.Println("获取的STMT:", insertStmt)
		defer insertStmt.Close()

		res, err := insertStmt.Exec(reqData["user"]["id"], reqData["user"]["token"], reqData["user"]["key"], time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Println("保存用户连接错误:", err)
			return
		}
		num, err := res.RowsAffected()

		if err != nil {
			log.Println("读取保存用户连接影响行数错误:", err)
			num = 0
			return
		}
		log.Println("获取到得影响行数:", num)
		data := make(map[string]interface{})
		data["status"] = num
		client.PutOut(NewIMResponseData(common.SetData("conn", data), GET_CONN_RETURN))
	// break go 中可以不加break 默认自动添加break

	case GET_BUDDY_LIST:
		// 获取好友分组列表
		if reqData["user"]["token"] == "" {
			client.PutOut(NewIMResponseSimple(402, "用户令牌不能为空!", GET_CONN_RETURN))
			return
		}
		var groups []model.Group
		rows, _ := Database.Query("select g.id,g.name from im_group g left join im_login l on l.user_id=g.creater where token=?", reqData["user"]["token"])
		defer rows.Close()
		for rows.Next() {
			var group model.Group
			rows.Scan(&group.Id, &group.Name)
			groups = append(groups, group)
		}
		//根据分组获取好友
		for k, v := range groups {
			rows, _ := Database.Query("select u.id,u.nick,u.status,u.sign,u.avatar from im_user u left join im_relation_user_group ug on u.id=ug.user_id where ug.group_id=?", v.Id)
			for rows.Next() {
				var user model.IMUser
				rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
				groups[k].Buddies = append(v.Buddies, user)
			}
		}
		client.PutOut(NewIMResponseData(common.SetData("categories", groups), GET_BUDDY_LIST_RETURN))
		break

	case CREATE_SESSION:
		// 创建会话  //{"command":"CREATE_SESSION","data":{"session":{"sender":"xxx","receiver":"xxx","token":"xxxx"}}}
		if reqData["session"]["sender"] == "" {
			client.PutOut(NewIMResponseSimple(401, "发送者不能为空!", CREATE_SESSION_RETURN))
			return
		}
		if reqData["session"]["receiver"] == "" {
			client.PutOut(NewIMResponseSimple(303, "接收者不能为空!", CREATE_SESSION_RETURN))
			return
		}
		if reqData["session"]["token"] == "" {
			client.PutOut(NewIMResponseSimple(303, "用户令牌不能为空!", CREATE_SESSION_RETURN))
			return
		}
		insertStmt, _ := Database.Prepare("insert into im_conversation VALUES (?, ?, ?, ?,?)")
		defer insertStmt.Close()
		var id = uuid.New()
		res, err := insertStmt.Exec(id, reqData["session"]["sender"], time.Now().Format("2006-01-02 15:04:05"), reqData["session"]["receiver"], "0")
		if err != nil {
			log.Println("创建会话错误:", err)
			id = ""
			return
		}
		num, err := res.RowsAffected()
		log.Println("影响行数：", num)
		log.Println("ID：", id)
		if err != nil {
			log.Println("读取保存用户连接影响行数错误:", err)
			num = 0
			return
		}
		data := make(map[string]string)
		data["ticket"] = id
		data["receiver"] = reqData["session"]["receiver"]
		client.PutOut(NewIMResponseData(common.SetData("session", data), CREATE_SESSION_RETURN))
		break

	case SEND_MSG:
		//{"command":"SEND_MSG","data":{"message":{"content":"xxxx","ticket":"xxx","token":"xxx"}}}
		if reqData["message"]["content"] == "" {
			client.PutOut(NewIMResponseSimple(401, "消息内容不能为空!", SEND_MSG_RETURN))
			return
		}
		if reqData["message"]["ticket"] == "" {
			client.PutOut(NewIMResponseSimple(402, "Ticket不能为空!", SEND_MSG_RETURN))
			return
		}
		//FIXME 此处应该先校验会话的有效性
		//获取接收人的连接的 name
		var key string
		var sender string
		err := Database.QueryRow("select c1.`key`,c2.`creater` from im_conn c1 left join im_conversation c2 on c1.user_id=c2.receiver where c2.id=?", reqData["message"]["ticket"]).Scan(&key, &sender)

		if err != nil || this.clients[key] == nil {
			client.PutOut(NewIMResponseSimple(403, "对方未登录!", SEND_MSG_RETURN))
			return
		} else {
			//把消息转发给接收者
			data := make(map[string]string)
			data["sender"] = sender
			data["content"] = reqData["message"]["content"]
			data["ticket"] = reqData["message"]["ticket"]

			this.clients[key].PutOut(NewIMResponseData(common.SetData("message", data), PUSH_MSG)) //reqData["message"]["content"]
		}
		// 发送消息，转发消息

		break

	case SEND_STATUS_CHANGE:
		//{"command":"SEND_STATUS_CHANGE","data":{"user":{"token":"xxxx","status":"1"}}}
		if reqData["user"]["token"] == "" {
			client.PutOut(NewIMResponseSimple(501, "TOKEN不能为空!", SEND_STATUS_CHANGE))
			return
		}
		if reqData["user"]["status"] == "" {
			client.PutOut(NewIMResponseSimple(501, "TOKEN不能为空!", SEND_STATUS_CHANGE))
			return
		}
		var id string
		//先获取当前用户的ID
		err := Database.QueryRow("SELECT user_id FROM im_conn WHERE token = ?", reqData["user"]["token"]).Scan(&id)
		if err != nil {
			client.PutOut(NewIMResponseSimple(403, "您已经掉线，请重新登录!", SEND_MSG_RETURN))
			return
		} else {
			//FIXME 还应该先校验token的有效性，当前用户的状态是不是和传得状态相同等
			updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
			defer updateStmt.Close()

			res, err := updateStmt.Exec(reqData["user"]["status"], id)
			if err != nil {
				log.Println("更新用户状态错误:", err)
				return
			}
			_, err = res.RowsAffected()

			if err != nil {
				log.Println("读取修改用户状态影响行数错误:", err)
				return
			}

			//FIXME 应该先校验用户的状态 再考虑广播好友
			//获取当前用户已连接的所有好友
			rows, _ := Database.Query("select co.`key` from im_conn co where co.user_id in (select ug.user_id from im_relation_user_group ug where ug.group_id in (select g.id from  im_group g where g.creater=?))", id)
			for rows.Next() {
				var key string
				rows.Scan(&key)
				//给对应的连接推送好友状态变化的通知
				data := make(map[string]string)
				data["id"] = id
				data["state"] = reqData["user"]["status"]
				this.clients[key].PutOut(NewIMResponseData(common.SetData("user", data), PUSH_STATUS_CHANGE))
			}
		}

		// 发送状态，转发状态
		break

	case LOGOUT_REQUEST:
		// 退出//{"command":"SEND_STATUS_CHANGE","data":{"user":{"token":"xxxx"}}}
		if reqData["user"]["token"] == "" {
			client.PutOut(NewIMResponseSimple(501, "TOKEN不能为空!", SEND_STATUS_CHANGE))
			return
		}
		var id string
		//先获取当前用户的ID
		err := Database.QueryRow("SELECT user_id FROM im_conn WHERE token = ?", reqData["user"]["token"]).Scan(&id)
		if err != nil {
			client.PutOut(NewIMResponseSimple(403, "您已经掉线，请重新登录!", SEND_MSG_RETURN))
			return
		} else {
			tx, _ := Database.Begin()
			//FIXME 还应该先校验token的有效性，当前用户的状态是不是和传得状态相同等
			updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = '0' WHERE id =?")
			defer updateStmt.Close()
			res, err := updateStmt.Exec(id)
			if err != nil {
				log.Println("更新用户状态错误:", err)
				return
			}
			_, err = res.RowsAffected()

			if err != nil {
				log.Println("读取修改用户状态影响行数错误:", err)
				return
			}
			//删除连接该token的连接
			delStmt, _ := Database.Prepare("delete from im_conn where token=?")
			defer delStmt.Close()
			_, err = delStmt.Exec(reqData["user"]["token"])
			if err != nil {
				log.Println("删除用户连接错误:", err)
				tx.Rollback()
				return
			}
			tx.Commit()
			//FIXME 应该先校验用户的状态 再考虑广播好友
			//获取当前用户已连接的所有好友
			rows, _ := Database.Query("select co.`key` from im_conn co where co.user_id in (select ug.user_id from im_relation_user_group ug where ug.group_id in (select g.id from  im_group g where g.creater=?))", id)
			for rows.Next() {
				var key string
				rows.Scan(&key)
				//给对应的连接推送好友状态变化的通知
				data := make(map[string]string)
				data["id"] = id
				data["state"] = reqData["user"]["status"]
				this.clients[key].PutOut(NewIMResponseData(common.SetData("user", data), PUSH_STATUS_CHANGE))
			}
		}
		client.Quit()
		break
	}
}

func (this *Server) start() {
	// 设置监听地址及端口
	addr := fmt.Sprintf("0.0.0.0:%d", Config.IMPort)
	this.listener, _ = net.Listen("tcp", addr)
	log.Printf("开始监听服务器[%d]端口", Config.IMPort)
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
		log.Printf("新连接地址为:[%s]", conn.RemoteAddr())
		this.joinsniffer <- conn
	}
}

// FIXME: need to figure out if this is the correct approach to gracefully
// terminate a server.
func (this *Server) Stop() {
	this.listener.Close()
}
