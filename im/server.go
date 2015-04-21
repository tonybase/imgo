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

type InMessage chan IMRequest
type OutMessage chan IMResponse
type ClientTable map[string]*Client

type Server struct {
	listener net.Listener
	clients  ClientTable
	pending  chan net.Conn
	quiting  chan *Client
	incoming InMessage
	outgoing OutMessage
}

// 启动IM服务
func StartIMServer(config common.IMConfig) error {
	log.Printf("IMServer starting...")

	server := &Server{
		clients:  make(ClientTable, Config.MaxClients),
		pending:  make(chan net.Conn),
		quiting:  make(chan *Client),
		incoming: make(InMessage),
		outgoing: make(OutMessage),
	}
	server.listen()
	server.start()
	return nil
}

func (this *Server) listen() {
	go func() {
		for {
			select {
			// 接收到了消息
			case message := <-this.incoming:
				this.received(message)

			// 新来了一个连接
			case conn := <-this.pending:
				this.join(conn)

			// 退出了一个连接
			case client := <-this.quiting:
				this.leave(client)
			}
		}
	}()
}

func (this *Server) join(conn net.Conn) {
	name := uuid.New()
	client := CreateClient(name, conn)
	this.clients[name] = client

	log.Printf("Auto assigned name %s\n", client.GetName())

	// 接收消息
	go func() {
		for {
			msg := <-client.incoming
			log.Printf("Got message: %s from client %s\n", msg, client.GetName())

			// fallthrough to normal message if it is not parsable or executable
			this.incoming <- msg
		}
	}()

	// 等待断开
	go func() {
		for {
			conn := <-client.quiting
			log.Printf("Client %s is quiting\n", client.GetName())
			this.quiting <- conn
		}
	}()

	// 返回连接的唯一标识
	// { "data": { "conn": { “key”:”xxxxx” } }, "msg": "", “status”:0, “refer”:”GET_KEY_RETURN” }
	data := make(map[string]interface{})
	data["key"] = name // 返回这个conn的唯一标识uuid
	resp := NewIMResponseData(common.GetJson("conn", data), GET_KEY_RETURN)
	client.PutOutgoing(resp)
}

func (this *Server) leave(client *Client) {
	if client != nil {
		client.Close()
		delete(this.clients, client.GetName())
	}
}

func (this *Server) received(req IMRequest) {
	client := req.Client
	log.Printf("Received message: %s %s %s\n", client.GetName(), req.Command, req.Data)
	log.Println("获取到的命令：", req.Command)
	// defer client.Close()
	reqData := req.Data
	switch req.Command {
	case GET_CONN:
		log.Println("获取到的数据", req.Data)

		log.Println("获取到的数据:", reqData["user"]["key"])

		if reqData["user"]["id"] == "" {
			client.PutOutgoing(NewIMResponseSimple(302, "用户ID不能为空!", GET_CONN_RETURN))
			return
		}
		if reqData["user"]["token"] == "" {

			client.PutOutgoing(NewIMResponseSimple(303, "用户令牌不能为空!", GET_CONN_RETURN))
			return
		}
		//FIXME 暂时不加，方便测试 !strings.EqualFold(reqData["user"]["key"], name)
		if reqData["user"]["key"] == "" {
			client.PutOutgoing(NewIMResponseSimple(304, "连接的KEY错误!", GET_CONN_RETURN))
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
		client.PutOutgoing(NewIMResponseData(common.GetJson("conn", data), GET_CONN_RETURN))
	// break go 中可以不加break 默认自动添加break

	case GET_BUDDY_LIST:
		// 获取好友分组列表

		log.Println("获取好友列表")
		var groups []model.Group
		rows, _ := Database.Query("select g.id,g.name from im_group g left join im_login l on l.user_id=g.creater where token=?", reqData["user"]["token"])
		defer rows.Close()
		for rows.Next() {
			var group model.Group
			rows.Scan(&group.Id, &group.Name)
			groups = append(groups, group)
		}
		//根据分组获取好友
		log.Println("根据分组获取对应的好友列表")
		for _, v := range groups {
			var users []model.IMUser //每个分组最多拥有100个好友
			rows, _ := Database.Query("select u.id,u.nick,u.status,u.sign,u.avatar from im_user u left join im_relation_user_group ug on u.id=ug.user_id where ug.group_id=?", v.Id)
			for rows.Next() {
				var user model.IMUser
				rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
				users = append(users, user)

			}
			v.Buddies = users

		}
		client.PutOutgoing(NewIMResponseData(common.GetJson("categories", groups), GET_BUDDY_LIST_RETURN))
		break

	case CREATE_SESSION:
		// 创建会话  //{"command":"CREATE_SESSION","data":{"session":{"sender":"xxx","receiver":"xxx","token":"xxxx"}}}
		if reqData["session"]["sender"] == "" {
			client.PutOutgoing(NewIMResponseSimple(401, "发送者不能为空!", CREATE_SESSION_RETURN))
			return
		}
		if reqData["session"]["receiver"] == "" {
			client.PutOutgoing(NewIMResponseSimple(303, "接收者不能为空!", CREATE_SESSION_RETURN))
			return
		}
		if reqData["session"]["token"] == "" {
			client.PutOutgoing(NewIMResponseSimple(303, "用户令牌不能为空!", CREATE_SESSION_RETURN))
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
		client.PutOutgoing(NewIMResponseData(common.GetJson("session ", data), CREATE_SESSION_RETURN))
		break

	case SEND_MSG:
		//{"command":"SEND_MSG","data":{"message":{"content":"xxxx","ticket":"xxx","token":"xxx"}}}
		if reqData["message"]["content"] == "" {
			client.PutOutgoing(NewIMResponseSimple(401, "消息内容不能为空!", SEND_MSG_RETURN))
			return
		}
		if reqData["message"]["ticket"] == "" {
			client.PutOutgoing(NewIMResponseSimple(402, "Ticket不能为空!", SEND_MSG_RETURN))
			return
		}
		//FIXME 此处应该先校验会话的有效性
		//获取接收人的连接的 name
		var key string
		var sender string
		err := Database.QueryRow("select c1.`key`,c2.`creater` from im_conn c1 left join im_conversation c2 on c1.user_id=c2.receiver where c2.id=?", reqData["message"]["ticket"]).Scan(&key, &sender)

		if err != nil {
			client.PutOutgoing(NewIMResponseSimple(403, "对方未登录!", SEND_MSG_RETURN))
			return
		} else {
			//把消息转发给接收者
			data := make(map[string]string)
			data["sender"] = sender
			data["content"] = reqData["message"]["content"]
			data["ticket"] = reqData["message"]["ticket"]

			this.clients[key].PutOutgoing(NewIMResponseData(common.GetJson("message ", data), PUSH_MSG)) //reqData["message"]["content"]
		}
		// 发送消息，转发消息

		break

	case SEND_STATUS_CHANGE:
		//{"command":"SEND_STATUS_CHANGE","data":{"user":{"token":"xxxx","status":"1"}}}
		if reqData["user"]["token"] == "" {
			client.PutOutgoing(NewIMResponseSimple(501, "TOKEN不能为空!", SEND_STATUS_CHANGE))
			return
		}
		if reqData["user"]["status"] == "" {
			client.PutOutgoing(NewIMResponseSimple(501, "TOKEN不能为空!", SEND_STATUS_CHANGE))
			return
		}
		var id string
		//先获取当前用户的ID
		err := Database.QueryRow("SELECT user_id FROM im_conn WHERE token = ?", reqData["user"]["token"]).Scan(&id)
		if err != nil {
			client.PutOutgoing(NewIMResponseSimple(403, "您已经掉线，请重新登录!", SEND_MSG_RETURN))
			return
		} else {
			//FIXME 还应该先校验token的有效性，当前用户的状态是不是和传得状态相同等
			updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = ? WHERE id =?)")
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
				this.clients[key].PutOutgoing(NewIMResponseData(common.GetJson("user ", data), PUSH_STATUS_CHANGE))
			}
		}

		// 发送状态，转发状态
		break

	case LOGOUT_REQUEST:
		// 退出//{"command":"SEND_STATUS_CHANGE","data":{"user":{"token":"xxxx"}}}
		if reqData["user"]["token"] == "" {
			client.PutOutgoing(NewIMResponseSimple(501, "TOKEN不能为空!", SEND_STATUS_CHANGE))
			return
		}
		var id string
		//先获取当前用户的ID
		err := Database.QueryRow("SELECT user_id FROM im_conn WHERE token = ?", reqData["user"]["token"]).Scan(&id)
		if err != nil {
			client.PutOutgoing(NewIMResponseSimple(403, "您已经掉线，请重新登录!", SEND_MSG_RETURN))
			return
		} else {
			tx, _ := Database.Begin()
			//FIXME 还应该先校验token的有效性，当前用户的状态是不是和传得状态相同等
			updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = '0' WHERE id =?)")
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
				this.clients[key].PutOutgoing(NewIMResponseData(common.GetJson("user ", data), PUSH_STATUS_CHANGE))
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

	log.Printf("IMServer 开始监听端口: %d", Config.IMPort)

	for {
		conn, err := this.listener.Accept()

		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Printf("A new connection %s \n", conn.RemoteAddr())

		this.pending <- conn
	}
}

// FIXME: need to figure out if this is the correct approach to gracefully
// terminate a server.
func (this *Server) Stop() {
	this.listener.Close()
}
