package common

import (
	"encoding/json"
)

type IMRequest struct {
	Client  *Client                      `json:"-"`
	Command string                       `json:"command"`
	Data    map[string]map[string]string `json:"data"`
}
/*
输入消息通道
*/
type InMessage chan IMRequest
/*
转成JSON数据
*/
func (this *IMRequest) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}
/*
解析JSON数据
*/
func (this *IMRequest) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}
/*
解析JSON数据
*/
func DecodeIMRequest(data []byte) (*IMRequest, error) {
	req := new(IMRequest)
	err := req.Decode(data)
	if err != nil {
		return nil, err
	}
	return req, nil
}
