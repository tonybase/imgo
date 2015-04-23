package model

import (
	"encoding/json"
)

type Group struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Buddies []IMUser `json:"buddies"`
}

func (this *Group) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

func (this *Group) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

func (this *Group) AddUser(u IMUser) {
	this.Buddies = append(this.Buddies, u)
}
