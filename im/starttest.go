package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"encoding/json"
)

type IMResponse struct {
	Status int                               `json:"status"` //状态 0成功，非0错误
	Msg    string                            `json:"msg"`    //消息
	Data   map[string]map[string]interface{} `json:"data"`   //数据
	Refer  string                            `json:"refer"`  //来源
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	var recv chan string = make(chan string)

	go func() {
		for {
			if line, _, err := reader.ReadLine(); err == nil {
				var str string
				str = string(line)
				log.Println(str)
				recv <- str
			} else {
				os.Exit(0)
			}
		}
	}()

	//	token := "11"
	sender := "11"
	receiver := "22"
	token := sender
	var res IMResponse
	var ticket string
	for {
		// 收到消息，然后再回复
		line := <-recv
		if line != "" {
			json.Unmarshal([]byte(line), &res)
			switch res.Refer {
			case "GET_KEY_RETURN":
				// 建立连接
				line = "{\"command\":\"GET_CONN\",\"data\":{\"user\":{\"id\":\"" + sender + "\",\"token\":\"" + token + "\",\"key\":\"" + res.Data["conn"]["key"].(string) + "\"}}}"
			case "GET_CONN_RETURN":
				// 创建会话
				line = "{\"command\":\"CREATE_SESSION\",\"data\":{\"session\":{\"sender\":\"" + sender + "\",\"receiver\":\"" + receiver + "\",\"token\":\"" + token + "\"}}}"

			case "CREATE_SESSION_RETURN":
				// 发送消息
				ticket = res.Data["session"]["ticket"].(string)
				line = "{\"command\":\"SEND_MSG\",\"data\":{\"message\":{\"content\":\"Hello  World\",\"ticket\":\"" + ticket + "\",\"token\":\"" + token + "\"}}}"

			case "PUSH_MSG":
				// 发送消息
				line = "{\"command\":\"SEND_MSG\",\"data\":{\"message\":{\"content\":\"Hello  World\",\"ticket\":\"" + res.Data["session"]["ticket"].(string) + "\",\"token\":\"" + token + "\"}}}"

			}
		}
		writer.WriteString(string(line) + "\n")
		err := writer.Flush()
		if (err != nil) {
			os.Exit(0)
		}
	}

}
