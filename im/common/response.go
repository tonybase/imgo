package common

import (
	"encoding/json"
)

/*
返回消息结构体
*/
type IMResponse struct {
	Status int         `json:"status"` //状态 0成功，非0错误
	Msg    string      `json:"msg"`    //消息
	Data   interface{} `json:"data"`   //数据
	Refer  string      `json:"refer"`  //来源
}

/*
输出消息通道
*/
type OutMessage chan IMResponse

/*
错误消息构造方法
*/
func NewIMResponseSimple(status int, msg string, refer string) *IMResponse {
	return &IMResponse{status, msg, nil, refer}
}

/*
成功消息构造方法
*/
func NewIMResponseData(data interface{}, refer string) *IMResponse {
	return &IMResponse{0, "", data, refer}
}

/*
将返回消息转成JSON
*/
func (this *IMResponse) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
将JSON转成返回消息
*/
func (this *IMResponse) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

