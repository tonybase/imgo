package im

import (
    "code.google.com/p/go-uuid/uuid"
    "log"
    "net"
    "fmt"
    "im-go/im/common"
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
    data["key"] = name  // 返回这个conn的唯一标识uuid
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

    switch req.Command {
        case GET_CONN:
        // 建立TCP长连接
        // "conn": { “status”:1 }
        data := make(map[string]interface{})
        data["status"] = 1 // key
        client.PutOutgoing(NewIMResponseData(common.GetJson("conn", data), GET_CONN_RETURN))
        break

        case GET_BUDDY_LIST:
        // 获取好友分组列表
        categories := GetCategories("account");
        client.PutOutgoing(NewIMResponseData(common.GetJson("categories", categories), GET_BUDDY_LIST_RETURN))
        break

        case CREATE_SESSION:
        // 创建会话
        client.PutOutgoing(NewIMResponseData(nil, CREATE_SESSION_RETURN))
        break

        case SEND_MSG:
        // 发送消息，转发消息

        break

        case SEND_STATUS_CHANGE:
        // 发送状态，转发状态
        break

        case LOGOUT_REQUEST:
        // 退出
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

