package model

import (
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

/*
用户对象
*/
type User struct {
	Id       string    `json:"id"`        //id
	Nick     string    `json:"nick"`      //昵称
	Status   string    `json:"status"`    //状态 0离线,1在线
	Sign     string    `json:"sign"`      //个性签名
	Avatar   string    `json:"avatar"`    //头像
	CreateAt time.Time `json:"create_at"` //注册日期
	UpdateAt time.Time `json:"update_at"` //更新日期
	Token    string    `json:"token"`
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
func CheckAccount(account string) (int, error) {
	var num int
	rows, err := Database.Query("select count(*) from im_user where account=? ", account)
	if err != nil {
		return -1, &DatabaseError{"根据账号查询用户错误"}
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&num)
	}
	return num, nil

}

/*
 根据ID获取用户
*/
func GetUserById(id string) (*User, error) {
	var user User
	row := Database.QueryRow("select id, nick, status, sign, avatar, create_at, update_at from im_user where id=?", id)
	err := row.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		return nil, &DatabaseError{"根据ID查询用户-将结果映射至对象错误"}
	}
	return &user, err
}

/*
 根据token获取用户
*/
func GetUserByToken(token string) (*User, error) {
	var user User
	row := Database.QueryRow("select u.id, u.nick, u.status, u.sign, u.avatar, u.create_at, u.update_at from  im_user u left join im_login l on u.id=l.user_id where l.token=?", token)
	err := row.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		return nil, &DatabaseError{"根据Token查询用户-将结果映射至对象错误"}
	}
	return &user, nil
}

/*
 根据分组获取好友列表
*/
func GetBuddiesByCategories(categories []Category) ([]Category, error) {
	for k, v := range categories {
		rows, err := Database.Query("select u.id, u.nick, u.status, u.sign, u.avatar, u.create_at, u.update_at from im_user u left join im_relation_user_category ug on u.id=ug.user_id where ug.category_id=?", v.Id)
		if err != nil {
			return categories, &DatabaseError{"根据分类查询好友列表错误"}
		}
		for rows.Next() {
			var user User
			rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
			categories[k].AddUser(user)
		}
	}
	return categories, nil
}

/*
 登录账号
*/
func LoginUser(account string, password string) (*User, error) {
	var user User
	rows, err := Database.Query("select id, nick, status, sign, avatar, create_at, update_at from im_user where account=? and password=? ", account, password)
	if err != nil {
		return nil, &DatabaseError{"根据账号及密码查询用户错误"}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			return nil, &DatabaseError{"根据账号及密码查询结果映射至对象错误"}
		}
	}
	return &user, nil
}

/*
 保存用户
*/
func SaveUser(account string, password string, nick string, avatar string) (*string, error) {
	insStmt, err := Database.Prepare("insert into im_user (id, account, password, nick, avatar, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, &DatabaseError{"保存用户数据库处理错误"}
	}
	defer insStmt.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	uid := uuid.New()
	_, err = insStmt.Exec(uid, account, password, nick, avatar, now, now)
	if err != nil {
		return nil, &DatabaseError{"保存用户记录错误"}
	}
	// 添加默认分类
	AddCategory(uid, "我的好友")
	return &uid, nil
}

/*
 修改用户状态
*/
func UpdateUserStatus(userId string, status string) (int64, error) {
	updateStmt, _ := Database.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, userId)
	if err != nil {
		return -1, &DatabaseError{"更新用户状态错误"}
	}
	num, err := res.RowsAffected()
	if err != nil {
		return -1, &DatabaseError{"读取修改用户状态影响行数错误"}
	}
	return num, nil
}

/*
 修改用户状态(事务)
*/
func UpdateUserStatusTx(tx *sql.Tx, userId string, status string) (int64, error) {
	var num int64
	updateStmt, err := tx.Prepare("UPDATE im_user SET `status` = ? WHERE id =?")
	if err != nil {
		return -1, &DatabaseError{"修改用户状态数据库处理错误"}
	}
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, userId)
	if err != nil {
		tx.Rollback()
		return -1, &DatabaseError{"更新用户状态错误"}
	}
	num, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return -1, &DatabaseError{"读取修改用户状态影响行数错误"}
	}
	return num, nil
}

/*
 根据用户ID获取在线好友的连接KEY列表
*/
func GetBuddiesKeyById(id string) ([]string, error) {
	var keys []string
	rows, err := Database.Query("select co.`id` from im_conn co where co.user_id in (select ug.user_id from im_relation_user_category ug where ug.category_id in (select g.id from im_category g where g.creator=?))", id)
	if err != nil {
		return keys, &DatabaseError{"根据用户ID获取在线好友的连接KEY列表错误"}
	}
	for rows.Next() {
		var key string
		rows.Scan(&key)
		keys = append(keys, key)
	}
	return keys, nil
}

/*
 根据条件查询获取好友列表
*/
func QueryUser(clumn string, reg string, data string) ([]User, error) {
	var users []User
	sql := "select u.id, u.nick, u.status, u.sign, u.avatar, u.create_at, u.update_at from im_user u where %s %s %s "

	if reg == "like" {
		data = "'%" + data + "%'"
	}
	sql = fmt.Sprintf(sql, clumn, reg, data)
	rows, err := Database.Query(sql)
	if err != nil {
		return users, &DatabaseError{"根据" + clumn + "查询用户错误"}
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar, &user.CreateAt, &user.UpdateAt)
		users = append(users, user)
	}
	return users, nil
}
