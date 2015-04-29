package model

import (
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"encoding/json"
	"im-go/im/util"
	"log"
	"time"
)

/*
用户对象
*/
type IMUser struct {
	Id     string `json:"id"`     //id
	Nick   string `json:"nick"`   //昵称
	Status string `json:"status"` //状态 0离线,1在线
	Sign   string `json:"sign"`   //个性签名
	Avatar string `json:"avatar"` //头像
}

/*
转JSON数据
*/
func (this *IMUser) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
解析JSON
*/
func (this *IMUser) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

/*
检查账号是否存在
*/
func CheckAccount(account string) int {
	var num int
	rows, err := Database.Query("select count(*)  from im_user where account=? ", account)

	if err != nil {
		log.Printf("根据账号查询用户错误: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&num)
	}
	return num

}

/*
根据ID获取用户
*/
func GetUserById(id string) IMUser {
	var user IMUser
	row := Database.QueryRow("select id,nick,status,sign,avatar from im_user where id=?", id)
	err := row.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
	if err != nil {
		log.Printf("根据ID查询用户-将结果映射至对象错误:", err)
	}
	return user
}

/*
根据token获取用户
*/
func GetUserByToken(token string) IMUser {
	var user IMUser
	row := Database.QueryRow("select u.id,u.nick,u.status,u.sign,u.avatar from  im_user u left join im_login l on u.id=l.user_id where l.token=?", token)
	err := row.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
	if err != nil {
		log.Printf("根据Token查询用户-将结果映射至对象错误:", err)
		return user
	}
	return user
}

/*
根据token获取用户登录
*/
func GetLoginByToken(token string) map[string]string {
	var data map[string]string
	res, err := Database.Query("select id,user_id,token from im_login where token=?", token)
	if err != nil {
		log.Println("根据Token获取用户登录错误", err)
	} else {
		data = util.ResToMap(res)
	}
	return data

}

/*
根据分组获取好友列表
*/
func GetBuddiesByGroups(groups []Group) []Group {
	for k, v := range groups {
		rows, _ := Database.Query("select u.id,u.nick,u.status,u.sign,u.avatar from im_user u left join im_relation_user_group ug on u.id=ug.user_id where ug.group_id=?", v.Id)
		for rows.Next() {
			var user IMUser
			rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
			groups[k].Buddies = append(v.Buddies, user)
		}
	}
	return groups
}

/*
登录账号
*/
func LoginUser(account string, password string) IMUser {
	var user IMUser
	rows, err := Database.Query("select id,nick,status,sign,avatar from im_user where account=? and password=? ", account, password)
	if err != nil {
		log.Printf("根据账号及密码查询用户错误: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
		if err != nil {
			log.Printf("根据账号及密码查询结果映射至对象错误:", err)
		}
	}
	return user
}

/*
保存登录状态
*/
func SaveLogin(userId string, token string, ip string) int64 {
	insStmt, _ := Database.Prepare("insert into im_login (id,user_id,token,login_date,login_ip) VALUES (?, ?, ?, ?,?)")
	defer insStmt.Close()
	res, err := insStmt.Exec(uuid.New(), userId, token, time.Now().Format("2006-01-02 15:04:05"), ip)
	if err != nil {
		log.Printf("保存用户登录记录错误: ", err)
		return 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Printf("读取保存用户登录记录影响行数错误:", err)
		return 0
	}
	return num
}

/*
修改用户状态
*/
func UpdateUserStatus(status string, id string) int64 {
	var num int64
	updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, id)
	if err != nil {
		log.Println("更新用户状态错误:", err)
		return 0
	}
	num, err = res.RowsAffected()
	if err != nil {
		log.Println("读取修改用户状态影响行数错误:", err)
		return 0
	}
	return num
}

/*
修改用户状态(事务)
*/
func UpdateUserStatusTx(tx *sql.Tx, status string, id string) int64 {
	var num int64
	updateStmt, _ := tx.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, id)
	if err != nil {
		tx.Rollback()
		log.Println("更新用户状态错误:", err)
		return 0
	}
	num, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println("读取修改用户状态影响行数错误:", err)
		return 0
	}
	return num
}

/*
根据用户ID获取在线好友的连接KEY列表
*/
func GetBuddiesKeyById(id string) []string {
	var keys []string
	rows, _ := Database.Query("select co.`key` from im_conn co where co.user_id in (select ug.user_id from im_relation_user_group ug where ug.group_id in (select g.id from  im_group g where g.creater=?))", id)
	for rows.Next() {
		var key string
		rows.Scan(&key)
		keys = append(keys, key)
	}
	return keys
}
