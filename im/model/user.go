package model

import (
    "encoding/json"
)
//用户模型
type IMUser struct {
    Id     string `json:"id"`
    Nick   string `json:"nick"`
    Status string `json:"status"`
    Sign   string `json:"sign"`
    Avatar string `json:"avatar"`
}

func (this *IMUser) Encode() []byte {
    s, _ := json.Marshal(*this)
    return s
}

func (this *IMUser) Decode(data []byte) error {
    err := json.Unmarshal(data, this)
    return err
}
