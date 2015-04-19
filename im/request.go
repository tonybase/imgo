package im

import (
    "encoding/json"
)

type IMRequest struct {
    Client *Client  `json:"-"`
    Command string  `json:"command"`

    Data map[string]map[string]string     `json:"data"`
}

func (this *IMRequest) Encode() []byte {
    s, _ := json.Marshal(*this)
    return s
}

func (this *IMRequest) Decode(data []byte) error {
    err := json.Unmarshal(data, this)
    return err
}

func DecodeIMRequest(data []byte) (*IMRequest, error) {
    req := new(IMRequest)
    err := req.Decode(data)
    if err != nil {
        return nil, err
    }
    return req, nil
}