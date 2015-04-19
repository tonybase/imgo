package im

import (
    "encoding/json"
)

//返回消息
type IMResponse struct {
    Status int                          `json:"status"`
    Msg    string                       `json:"msg"`
    Data   interface{}                  `json:"data"`
    Refer  string                       `json:"refer"`
}

func NewIMResponseSimple(status int, msg string) *IMResponse {
    return &IMResponse{status, msg, nil, ""}
}

func NewIMResponseData(data  interface{}, refer string) *IMResponse {
    return &IMResponse{0, "", data, refer}
}

func (this *IMResponse) Encode() []byte {
    s, _ := json.Marshal(*this)
    return s
}

func (this *IMResponse) Decode(data []byte) error {
    err := json.Unmarshal(data, this)
    return err
}

func DecodeIMResponse(data []byte) (*IMResponse, error) {
    resp := new(IMResponse)
    err := resp.Decode(data)
    if err != nil {
        return nil, err
    }
    return resp, nil
}