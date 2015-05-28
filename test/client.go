package main

import (
	"fmt"
	"net"
	"encoding/json"

)

const (
	tcpaddr = "127.0.0.1:9090"
)
type IMResponse struct {
	Status int         `json:"status"` //状态 0成功，非0错误
	Msg    string      `json:"msg"`    //消息
	Data   map[string]map[string]interface{} `json:"data"`   //数据
	Refer  string      `json:"refer"`  //来源
}
func main() {

	Client()
	/**for i:=0;i<3 ;i++  {
		 Client()
	}**/

}

func Client() {
	conn, err := net.Dial("tcp", tcpaddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	sms := make([]byte, 1024)
	sms=[]byte("")
	for {

		fmt.Println(string(sms))
		conn.Write(sms)
		buf := make([]byte, 1024)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		msg:=string(buf[0:c])
		fmt.Println(msg)
		if msg!=""{
			var res IMResponse
			json.Unmarshal(buf[0:c],&res)
			fmt.Println(res.Refer)
			fmt.Println(res.Data["conn"]["key"])
			switch res.Refer{
				case "GET_KEY_RETURN":
				sms=[]byte("{\"command\":\"GET_CONN\",\"data\":{\"user\":{\"id\":\"11\",\"token\":\"3233\",\"key\":\""+res.Data["conn"]["key"].(string)+"\"}}}")
				conn.Write(sms)
				case "GET_CONN_RETURN":

				case "CREATE_SESSION_RETURN":

			}

		}

	}

}
