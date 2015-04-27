package server

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"im-go/im/util"
	"log"
	"net"
	"im-go/im/common"
	"im-go/im/model"
)
/**
var(
InMessage=make(common.InMessage)//读取消息通道
OutMessage=make(common.OutMessage)//输出消息通道
ClientTable=make(common.ClientTable)//客户端列表
)
**/


/*
服务端结构体
*/
type Server struct {
	listener    net.Listener  //服务端监听器 监听xx端口
	clients     common.ClientTable   //客户端列表 抽象出来单独维护和入参 更方便管理连接
	joinsniffer chan net.Conn //访问连接嗅探器 触发创建客户端连接处理方法
	quitsniffer chan *common.Client  //连接退出嗅探器 触发连接退出处理方法
	insniffer   common.InMessage     //接收消息嗅探器 触发接收消息处理方法 对应客户端中in属性
}

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
	client := common.CreateClient(key, conn)
	//给客户端指定key
	this.clients[key] = client
	log.Printf("设置新请求客户端Key为:[%s]", client.GetKey())
	//开启协程不断地接收消息
	go func() {
		for {
			//客户端读取消息
			msg := <-client.In
			//消息交给嗅探器 触发对应的处理方法
			this.insniffer <- msg
		}
	}()
	//开启协程一直等待断开
	go func() {
		for {
			//客户端接收断开请求
			conn := <-client.Quit
			log.Printf("客户端:[%s]退出", client.GetKey())
			//请求交给嗅探器 触发对应的处理方法
			this.quitsniffer <- conn
		}
	}()
	//返回客户端的唯一标识
	data := make(map[string]interface{})
	data["key"] = key
	client.PutOut(common.NewIMResponseData(util.SetData("conn", data), common.GET_KEY_RETURN))
}

/*
客户端退出处理方法
*/
func (this *Server) quitHandler(client *common.Client) {
	if client != nil {
		//调用客户端关闭方法
		client.Close()
		delete(this.clients, client.GetKey())
	}
}

/*
接收消息处理方法
*/
func (this *Server) receivedHandler(request common.IMRequest) {
	log.Println("开始读取数据")
	log.Println("读取的数据为", request)
	//获取请求的客户端
	client := request.Client
	log.Printf("客户端:[%s]发送命令:[%s]消息内容:[%s]", client.GetKey(), request.Command, request.Data)
	//获取请求数据
	reqData := request.Data
	switch request.Command {
	case common.GET_CONN:
		if reqData["user"]["id"] == "" {
			client.PutOut(common.NewIMResponseSimple(302, "用户ID不能为空!", common.GET_CONN_RETURN))
			return
		}
		if reqData["user"]["token"] == "" {

			client.PutOut(common.NewIMResponseSimple(303, "用户令牌不能为空!", common.GET_CONN_RETURN))
			return
		}
		//FIXME 暂时不加，方便测试 !strings.EqualFold(reqData["user"]["key"], name)
		if reqData["user"]["key"] == "" {
			client.PutOut(common.NewIMResponseSimple(304, "连接的KEY错误!", common.GET_CONN_RETURN))
			return
		}
		var num int64
		//校验用户 得到完整的User对象
		user := model.GetUserById(reqData["user"]["id"])
		if user.Id != "" {
			//校验用户是否登录 得到map
			data := model.GetLoginByToken(reqData["user"]["token"])

			if data["id"] != "" && user.Status == "1" {
				conn := model.GetConnByUserId(user.Id)
				if conn["id"] != "" {
					num = model.UpdateConnByUserId(reqData["user"]["id"], reqData["user"]["token"], reqData["user"]["key"])

				} else {
					num = model.AddConn(reqData["user"]["id"], reqData["user"]["token"], reqData["user"]["key"])
				}
				data := make(map[string]interface{})
				data["status"] = num
				client.PutOut(common.NewIMResponseData(util.SetData("conn", data), common.GET_CONN_RETURN))
			} else {
				client.PutOut(common.NewIMResponseSimple(304, "用户未登录!", common.GET_CONN_RETURN))
				return
			}
		} else {
			client.PutOut(common.NewIMResponseSimple(304, "用户不存在!", common.GET_CONN_RETURN))
			return
		}

	case common.GET_BUDDY_LIST:
		// 获取好友分组列表
		if reqData["user"]["token"] == "" {
			client.PutOut(common.NewIMResponseSimple(402, "用户令牌不能为空!", common.GET_CONN_RETURN))
			return
		}
		//校验用户是否登录 得到map
		user := model.GetUserByToken(reqData["user"]["token"])
		data := model.GetLoginByToken(reqData["user"]["token"])
		log.Println("获取数据如下")
		log.Println("token:", reqData["user"]["token"])
		log.Println("id：", data["id"])
		log.Println("获取到用户", user)
		//return
		id := data["id"]
		if user.Id != "" {

			if id != "" && user.Status == "1" {
				groups := model.GetGroupsByToken(reqData["user"]["token"])
				groups = model.GetBuddiesByGroups(groups)
				client.PutOut(common.NewIMResponseData(util.SetData("categories", groups), common.GET_BUDDY_LIST_RETURN))
			} else {
				client.PutOut(common.NewIMResponseSimple(304, "用户未登录!", common.GET_CONN_RETURN))
				return
			}
		} else {
			client.PutOut(common.NewIMResponseSimple(304, "用户不存在!", common.GET_CONN_RETURN))
			return
		}
	case common.CREATE_SESSION:
		// 创建会话  //{"command":"CREATE_SESSION","data":{"session":{"sender":"xxx","receiver":"xxx","token":"xxxx"}}}
		if reqData["session"]["sender"] == "" {
			client.PutOut(common.NewIMResponseSimple(401, "发送者不能为空!", common.CREATE_SESSION_RETURN))
			return
		}
		if reqData["session"]["receiver"] == "" {
			client.PutOut(common.NewIMResponseSimple(303, "接收者不能为空!", common.CREATE_SESSION_RETURN))
			return
		}
		if reqData["session"]["token"] == "" {
			client.PutOut(common.NewIMResponseSimple(303, "用户令牌不能为空!", common.CREATE_SESSION_RETURN))
			return
		}
		sender := model.GetUserById(reqData["session"]["sender"])
		receiver := model.GetUserById(reqData["session"]["receiver"])
		data := model.GetLoginByToken(reqData["session"]["token"])
		var id string
		if sender.Id != "" {
			if receiver.Id != "" {
				if sender.Status == "1" && data["id"] != "" {
					//FIXME 还应该判断接收者是不是在线，如果不在线，消息缓存到留言表中 留着触发登录事件推送消息
					id = model.AddConversation(reqData["session"]["sender"], reqData["session"]["receiver"], reqData["session"]["token"])
					if id == "" {
						client.PutOut(common.NewIMResponseSimple(304, "创建会话失败", common.GET_CONN_RETURN))
						return

					} else {
						data := make(map[string]string)
						data["ticket"] = id
						data["receiver"] = reqData["session"]["receiver"]
						client.PutOut(common.NewIMResponseData(util.SetData("session", data), common.CREATE_SESSION_RETURN))
					}
				} else {
					client.PutOut(common.NewIMResponseSimple(304, "您还未登录!", common.GET_CONN_RETURN))
					return
				}

			} else {
				client.PutOut(common.NewIMResponseSimple(304, "接收者不存在!", common.GET_CONN_RETURN))
				return
			}

		} else {
			client.PutOut(common.NewIMResponseSimple(304, "账号不存在!", common.GET_CONN_RETURN))
			return
		}

	case common.SEND_MSG:
		if reqData["message"]["content"] == "" {
			client.PutOut(common.NewIMResponseSimple(401, "消息内容不能为空!", common.SEND_MSG_RETURN))
			return
		}
		if reqData["message"]["ticket"] == "" {
			client.PutOut(common.NewIMResponseSimple(402, "Ticket不能为空!", common.SEND_MSG_RETURN))
			return
		}
		if reqData["message"]["token"] == "" {
			client.PutOut(common.NewIMResponseSimple(402, "Token不能为空!", common.SEND_MSG_RETURN))
			return
		}
		conversion := model.GetConversationById(reqData["message"]["ticket"])
		if conversion["id"] != "" {
			sender := model.GetUserById(conversion["creater"])
			recevier := model.GetUserById(conversion["receiver"])
			if sender.Id != "" {
				if recevier.Id != "" {
					data := model.GetLoginByToken(reqData["message"]["token"])
					if sender.Status == "1" && data["id"] != "" {
						key := model.GetReceiverKeyByTicket(reqData["message"]["ticket"])
						/*
							if key == "" {
								client.PutOut(common.NewIMResponseSimple(402, "对方还未登录!", SEND_MSG_RETURN))
							}
						*/
						//把消息转发给接收者
						data := make(map[string]string)
						data["sender"] = sender.Id
						data["content"] = reqData["message"]["content"]
						data["ticket"] = reqData["message"]["ticket"]
						log.Println("开始转发给:", key)
						log.Println("当前的所有连接：",this.clients)
						this.clients[key].PutOut(common.NewIMResponseData(util.SetData("message", data), common.PUSH_MSG))
					} else {
						client.PutOut(common.NewIMResponseSimple(402, "您还未登录!", common.SEND_MSG_RETURN))
						return
					}

				} else {
					client.PutOut(common.NewIMResponseSimple(402, "接收者账户不存在!", common.SEND_MSG_RETURN))
					return
				}

			} else {
				client.PutOut(common.NewIMResponseSimple(402, "您的账户不存在不能发送消息!", common.SEND_MSG_RETURN))
				return
			}
		} else {
			client.PutOut(common.NewIMResponseSimple(402, "会话已关闭!", common.SEND_MSG_RETURN))
			return
		}

	case common.SEND_STATUS_CHANGE:
		//{"command":"SEND_STATUS_CHANGE","data":{"user":{"token":"xxxx","status":"1"}}}
		if reqData["user"]["token"] == "" {
			client.PutOut(common.NewIMResponseSimple(501, "TOKEN不能为空!", common.SEND_STATUS_CHANGE))
			return
		}
		if reqData["user"]["status"] == "" {
			client.PutOut(common.NewIMResponseSimple(501, "状态不能为空!", common.SEND_STATUS_CHANGE))
			return
		}
		conn := model.GetConnByToken(reqData["user"]["token"])
		login := model.GetLoginByToken(reqData["user"]["token"])
		user := model.GetUserByToken(reqData["user"]["token"])
		//判断用户的合法性
		if user.Id == "" {
			//判断该用户是不是合法已经登录的用户 不校验当前用户的状态 只校验当前用户是不是和系统还存在一致的连接
			if login["user_id"] == conn["user_id"] {
				//判断要求改变的状态和当前该用户的状态是否一致
				if user.Status != reqData["user"]["status"] {
					//FIXME 此处不做如果状态是离线就删除用户连接的操作,状态改变认为是客户端手动操作或者网络异常
					num := model.UpdateUserStatus(reqData["user"]["status"], user.Id)
					if num != 0 {
						keys := model.GetBuddiesKeyById(user.Id)
						for i := 0; i < len(keys); i++ {
							//给对应的连接推送好友状态变化的通知
							data := make(map[string]string)
							data["id"] = user.Id
							data["state"] = reqData["user"]["status"]
							this.clients[keys[i]].PutOut(common.NewIMResponseData(util.SetData("user", data), common.PUSH_STATUS_CHANGE))
						}
					} else {
						client.PutOut(common.NewIMResponseSimple(501, "修改状态失败,请重新尝试!", common.SEND_STATUS_CHANGE))
						return
					}
				} else {
					//将要修改的状态和当前状态一样 不做处理
				}

			} else {
				client.PutOut(common.NewIMResponseSimple(501, "请退出重新登录!", common.SEND_STATUS_CHANGE))
				return
			}

		} else {
			client.PutOut(common.NewIMResponseSimple(501, "Token不合法!", common.SEND_STATUS_CHANGE))
			return
		}
	case common.LOGOUT_REQUEST:
		// 退出//{"command":"SEND_STATUS_CHANGE","data":{"user":{"token":"xxxx"}}}
		if reqData["user"]["token"] == "" {
			client.PutOut(common.NewIMResponseSimple(501, "TOKEN不能为空!", common.SEND_STATUS_CHANGE))
			return
		}
		conn := model.GetConnByToken(reqData["user"]["token"])
		login := model.GetLoginByToken(reqData["user"]["token"])
		user := model.GetUserByToken(reqData["user"]["token"])
		//判断用户的合法性
		if user.Id == "" {
			//判断该用户是不是合法已经登录的用户 不校验当前用户的状态 只校验当前用户是不是和系统还存在一致的连接
			if login["user_id"] == conn["user_id"] {
				//判断要求改变的状态和当前该用户的状态是否一致
				if user.Status != "0" {
					tx, _ := model.Database.Begin()
					res1 := model.UpdateUserStatusTx(tx, "0", user.Id)
					res2 := model.DeleteConnByToken(tx, reqData["user"]["token"])
					if res1 == 1 && res2 == 1 {
						tx.Commit()
						client.PutOut(common.NewIMResponseSimple(501, "退出成功!", common.SEND_STATUS_CHANGE))
						client.Quiting()
						return
					} else {
						//全部回滚
						tx.Rollback()
						client.PutOut(common.NewIMResponseSimple(501, "退出失败，请稍后再试!", common.SEND_STATUS_CHANGE))
						return
					}
				} else {
					//将要修改的状态和当前状态一样 不做处理
				}

			} else {
				client.PutOut(common.NewIMResponseSimple(501, "请退出重新登录!", common.SEND_STATUS_CHANGE))
				return
			}

		} else {
			client.PutOut(common.NewIMResponseSimple(501, "Token不合法!", common.SEND_STATUS_CHANGE))
			return
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
		log.Printf("新连接地址为:[%s]", conn.RemoteAddr())
		this.joinsniffer <- conn
	}
}

// FIXME: need to figure out if this is the correct approach to gracefully
// terminate a server.
func (this *Server) Stop() {
	this.listener.Close()
}
