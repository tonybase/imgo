package model

import (
	"encoding/json"
	"time"
)

/*
 消息对象
 */
type Message struct {
	Id        string    //id
	Sender    string    //发送人
	To        string    //接收人
	Ticket    string    //ticket
	Token     string    //发送人登录token
	Content   string    //内容
	Create_at time.Time //时间
}

/*
 转JSON数据
 */
func (this *Message) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
 解析JSON数据
 */
func (this *Message) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}
