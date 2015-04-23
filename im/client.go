package im

import (
	"bufio"
	"log"
	"net"
)

/*
客户端结构体
*/
type Client struct {
	name     string //客户端连接的唯一性 FIXME 需要修改成KEY name太笼统不能体现唯一性的概念
	conn     net.Conn
	incoming InMessage
	outgoing OutMessage
	reader   *bufio.Reader
	writer   *bufio.Writer
	quiting  chan *Client
}

func (this *Client) GetName() string {
	return this.name
}

func (this *Client) SetName(name string) {
	this.name = name
}

func (this *Client) GetIncoming() IMRequest {
	return <-this.incoming
}

func (this *Client) PutOutgoing(resp *IMResponse) {
	this.outgoing <- *resp
}

func CreateClient(name string, conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	client := &Client{
		name:     name,
		conn:     conn,
		incoming: make(InMessage),
		outgoing: make(OutMessage),
		quiting:  make(chan *Client),
		reader:   reader,
		writer:   writer,
	}
	client.Listen()
	return client
}

// 自动读入或者写出消息
func (this *Client) Listen() {
	go this.read()
	go this.write()
}

// 退出了一个连接
func (this *Client) Quit() {
	this.quiting <- this
}

// 关闭连接通道
func (this *Client) Close() {
	this.conn.Close()
}

func (this *Client) read() {
	for {
		if line, _, err := this.reader.ReadLine(); err == nil {
			req, err := DecodeIMRequest(line)
			if err == nil {
				req.Client = this
				this.incoming <- *req
			} else {
				// 忽略消息，连命令都不知道，没办法处理
				log.Printf("解析JSON错误: %s", line)
			}
		} else {
			log.Printf("Read error: %s\n", err)
			this.Quit()
			return
		}
	}
}

func (this *Client) write() {
	for resp := range this.outgoing {
		if _, err := this.writer.WriteString(string(resp.Encode()) + "\n"); err != nil {
			this.Quit()
			return
		}
		if err := this.writer.Flush(); err != nil {
			log.Printf("Write error: %s\n", err)
			this.Quit()
			return
		}
	}
}
