package model

import (
	"encoding/json"
)

/*
分组对象
*/
type Group struct {
	Id      string   `json:"id"`      //分组ID
	Name    string   `json:"name"`    //分组名称
	Buddies []IMUser `json:"buddies"` //好友列表
}

/*
转JSON数据
*/
func (this *Group) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
 解析JSON数据
 */
func (this *Group) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

/*
 分组PUSH用户(仅传递数据 非操作数据库方法)
 */
func (this *Group) AddUser(u IMUser) {
	this.Buddies = append(this.Buddies, u)
}

/*
 根据token获取分组数据
 */
func GetGroupsByToken(token string) []Group {
	var groups []Group
	rows, _ := Database.Query("select g.id, g.name from im_group g left join im_login l on l.user_id=g.creator where token=?", token)
	defer rows.Close()
	for rows.Next() {
		var group Group
		rows.Scan(&group.Id, &group.Name)
		groups = append(groups, group)
	}
	return groups
}
