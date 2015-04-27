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
func GetGroupsByToken(token string) []Group {
	var groups []Group
	rows, _ := Database.Query("select g.id,g.name from im_group g left join im_login l on l.user_id=g.creater where token=?", token)
	defer rows.Close()
	for rows.Next() {
		var group Group
		rows.Scan(&group.Id, &group.Name)
		groups = append(groups, group)
	}
	return groups
}
