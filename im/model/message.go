package model

import (
    "time"
    "encoding/json"
)

type IMMessage struct {
    Id string
    Sender string
    To string
    Ticket string
    Token string
    Content string
    Create_at time.Time
}

func (this *IMMessage) Encode() []byte {
    s, _ := json.Marshal(*this)
    return s
}

func (this *IMMessage) Decode(data []byte) error {
    err := json.Unmarshal(data, this)
    return err
}