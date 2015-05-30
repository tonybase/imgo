package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Clienter struct {
	client  net.Conn
	isAlive bool
	SendStr chan string
	RecvStr chan string
}
type IMResponse struct {
	Status int                               `json:"status"` //状态 0成功，非0错误
	Msg    string                            `json:"msg"`    //消息
	Data   map[string]map[string]interface{} `json:"data"`   //数据
	Refer  string                            `json:"refer"`  //来源
}

func (c *Clienter) Connect() bool {
	if c.isAlive {
		return true
	} else {
		var err error
		c.client, err = net.Dial("tcp", "127.0.0.1:9090")
		if err != nil {
			fmt.Printf("Failure to connet:%s\n", err.Error())
			return false
		}
		c.isAlive = true
	}
	return true
}

func (c *Clienter) Echo() {
	line := <-c.SendStr
	c.client.Write([]byte(line))
	buf := make([]byte, 1024)
	n, err := c.client.Read(buf)
	if err != nil {
		c.RecvStr <- string("Server close...")
		c.client.Close()
		c.isAlive = false
		return
	}
	c.RecvStr <- string(buf[0:n])
}

func Work(tc *Clienter) {
	if !tc.isAlive {
		if tc.Connect() {
			tc.Echo()
		} else {
			<-tc.SendStr
			tc.RecvStr <- string("Server close...")
		}
	} else {
		tc.Echo()
	}
}
func main() {
	var tc Clienter
	tc.SendStr = make(chan string)
	tc.RecvStr = make(chan string)
	if !tc.Connect() {
		return
	}
	var line string
	var res IMResponse
	for {

		go Work(&tc)
		fmt.Printf("发送:%s\n", line)
		tc.SendStr <- line
		s := <-tc.RecvStr
		fmt.Printf("返回:%s\n", s)
		if tc.Connect() && s != "" {
			json.Unmarshal([]byte(s), &res)
			switch res.Refer {
			case "GET_KEY_RETURN":
				//建立连接
				line = "{\"command\":\"GET_CONN\",\"data\":{\"user\":{\"id\":\"11\",\"token\":\"11\",\"key\":\"" + res.Data["conn"]["key"].(string) + "\"}}}"
			case "GET_CONN_RETURN":
				//创建会话
				line = "{\"command\":\"CREATE_SESSION\",\"data\":{\"session\":{\"sender\":\"11\",\"receiver\":\"22\",\"token\":\"11\"}}}"

			case "CREATE_SESSION_RETURN":
				//发送消息session
				line = "{\"command\":\"SEND_MSG\",\"data\":{\"message\":{\"content\":\"Hello  World\",\"ticket\":\"" + res.Data["session"]["ticket"].(string) + "\",\"token\":\"11\"}}}"

			}

		}

		line += "\n"
	}
}
