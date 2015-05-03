package model

import (
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"encoding/json"
	"log"
	"time"
)

/*
用户对象
*/
type User struct {
	Id     string `json:"id"`     //id
	Nick   string `json:"nick"`   //昵称
	Status string `json:"status"` //状态 0离线,1在线
	Sign   string `json:"sign"`   //个性签名
	Avatar string `json:"avatar"` //头像
	CreateAt time.Time `json:"create_at"` //注册日期
	UpdateAt time.Time `json:"update_at"` //更新日期
}

/*
 转JSON数据
 */
func (this *User) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
 解析JSON
 */
func (this *User) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

/*
 检查账号是否存在
 */
func CheckAccount(account string) int {
	var num int
	rows, err := Database.Query("select count(*) from im_user where account=? ", account)

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
func GetUserById(id string) User {
	var user User
	row := Database.QueryRow("select id, nick, status, sign, avatar, create_at, update_at from im_user where id=?", id)
	err := row.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Printf("根据ID查询用户-将结果映射至对象错误: %s", err.Error())
	}
	return user
}

/*
 根据token获取用户
 */
func GetUserByToken(token string) User {
	var user User
	row := Database.QueryRow("select u.id, u.nick, u.status, u.sign, u.avatar, u.create_at, u.update_at from  im_user u left join im_login l on u.id=l.user_id where l.token=?", token)
	err := row.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Printf("根据Token查询用户-将结果映射至对象错误: %s", err.Error())
		return user
	}
	return user
}

/*
 根据分组获取好友列表
 */
func GetBuddiesByCategories(categories []Category) []Category {
	for k, v := range categories {
		rows, _ := Database.Query("select u.id, u.nick, u.status, u.sign, u.avatar, u.create_at, u.update_at from im_user u left join im_relation_user_category ug on u.id=ug.user_id where ug.category_id=?", v.Id)
		for rows.Next() {
			var user User
			rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
			categories[k].AddUser(user)
		}
	}
	return categories
}

/*
 登录账号
 */
func LoginUser(account string, password string) User {
	var user User
	rows, err := Database.Query("select id, nick, status, sign, avatar, create_at, update_at from im_user where account=? and password=? ", account, password)
	if err != nil {
		log.Printf("根据账号及密码查询用户错误: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			log.Printf("根据账号及密码查询结果映射至对象错误:", err)
		}
	}
	return user
}

/*
 保存用户
 */
func SaveUser(account string, password string, nick string, avatar string) int64 {
	insStmt, _ := Database.Prepare("insert into im_user (id, account, password, nick, avatar, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer insStmt.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	uid := uuid.New()
	res, err := insStmt.Exec(uid, account, password, nick, avatar, now, now)
	if err != nil {
		log.Printf("保存用户记录错误: ", err)
		return 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Printf("读取保存用户记录影响行数错误:", err)
		return 0
	}
	// 添加默认分类
	AddCategory(uid, "我的好友");
	return num
}

/*
 修改用户状态
 */
func UpdateUserStatus(userId string, status string) bool {
	updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, userId)
	if err != nil {
		log.Println("更新用户状态错误:", err)
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		log.Println("读取修改用户状态影响行数错误:", err)
		return false
	}
	return true
}

/*
 修改用户状态(事务)
 */
func UpdateUserStatusTx(tx *sql.Tx, userId string, status string) int64 {
	var num int64
	updateStmt, _ := tx.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, userId)
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
	rows, _ := Database.Query("select co.`id` from im_conn co where co.user_id in (select ug.user_id from im_relation_user_category ug where ug.category_id in (select g.id from im_category g where g.creator=?))", id)
	for rows.Next() {
		var key string
		rows.Scan(&key)
		keys = append(keys, key)
	}
	return keys
}