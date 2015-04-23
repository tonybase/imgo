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
	key    string        //客户端连接的唯标志
	conn   net.Conn      //连接
	in     InMessage     //输入消息
	out    OutMessage    //输出消息
	reader *bufio.Reader //读取
	writer *bufio.Writer //输出
	quit   chan *Client  //退出
}

/*
获取客户端名称
*/
func (this *Client) GetKey() string {
	return this.key
}

/*
设置客户端名称
*/
func (this *Client) SetKey(key string) {
	this.key = key
}

/*
获取输入消息
*/
func (this *Client) GetIn() IMRequest {
	return <-this.in
}

/*
设置输出消息
*/
func (this *Client) PutOut(resp *IMResponse) {
	this.out <- *resp
}

/*
创建客户端
*/
func CreateClient(key string, conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	client := &Client{
		key:    key,
		conn:   conn,
		in:     make(InMessage),
		out:    make(OutMessage),
		quit:   make(chan *Client),
		reader: reader,
		writer: writer,
	}
	client.Listen()
	return client
}

/*
自动读入或者写出消息
*/
func (this *Client) Listen() {
	go this.read()
	go this.write()
}

/*
退出了一个连接
*/
func (this *Client) Quit() {
	this.quit <- this
}

/*
关闭连接通道
*/
func (this *Client) Close() {
	this.conn.Close()
}

/*
读取消息
*/
func (this *Client) read() {
	for {
		if line, _, err := this.reader.ReadLine(); err == nil {
			req, err := DecodeIMRequest(line)
			if err == nil {
				req.Client = this
				this.in <- *req
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

/*
输出消息
*/
func (this *Client) write() {
	for resp := range this.out {
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
