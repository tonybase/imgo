package model

import (
	"database/sql"
	"log"
	"time"
)

/*
 连接对象
 */
type Conn struct {
	Key    string    `json:"key"`        // 连接key 唯一标识符
	UserId string    `json:"id"`        // 用户ID
	Token  string    `json:"token"`    // token
	CreateAt   time.Time `json:"create_at"`// 时间
	UpdateAt   time.Time `json:"update_at"`// 时间
}

/*
 根据Token获取连接的数量
 */
func CountConnByKey(key string) int64 {
	var num int64
	err := Database.QueryRow("select count(*) from im_conn where id=?", key).Scan(&num)
	if err != nil {
		num = -1
		log.Printf("根据Token获取连接数量错误: ", err)
	}
	return num
}

/*
 根据Token获取连接的数量
 */
func CountConnByToken(token string) int64 {
	var num int64
	err := Database.QueryRow("select count(*) from im_conn where token=?", token).Scan(&num)
	if err != nil {
		num = 0
		log.Printf("根据Token获取连接数量错误: ", err)
	}
	return num
}

/*
 根据用户ID修改连接
 */
func UpdateConnByToken(key string, userId string, token string) int64 {
	var num int64
	updateStmt, _ := Database.Prepare("UPDATE im_conn SET `id`=?, `user_id` = ?, update_at=? WHERE token =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(key, userId, time.Now().Format("2006-01-02 15:04:05"), token)
	if err != nil {
		log.Println("更新用户连接错误:", err)
		return 0
	}
	num, err = res.RowsAffected()
	if err != nil {
		log.Println("读取修改用户连接影响行数错误:", err)
		return 0
	}
	return num
}

/*
 根据token获取token
 */
func GetConnByToken(token string) Conn {
	var conn Conn
	row := Database.QueryRow("select * from im_conn where token=?", token)
	err := row.Scan(&conn.Key, &conn.UserId, &conn.Token, &conn.CreateAt, &conn.UpdateAt)
	if err != nil {
		log.Println("根据Token获取用户连接记录错误", err)
	}
	return conn
}

/*
 根据用户ID获取连接
 */
func GetConnByUserId(userId string) Conn {
	var conn Conn
	row := Database.QueryRow("select * from im_conn where user_id=?", userId)
	err := row.Scan(&conn.Key, &conn.UserId, &conn.Token, &conn.CreateAt, &conn.UpdateAt)
	if err != nil {
		log.Println("根据用户ID获取用户连接记录错误", err)
	}
	return conn
}

/*
 根据token删除连接
 */
func DeleteConnByKey(key string) int64 {
	//删除连接该token的连接
	delStmt, _ := Database.Prepare("delete from im_conn where id=?")
	defer delStmt.Close()
	res, err := delStmt.Exec(key)
	if err != nil {
		log.Println("删除用户连接错误:", err)
		return 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Println("读取删除用户连接影响行数错误:", err)
		return 0
	}
	return num
}

/*
 根据token删除连接
 */
func DeleteConnByToken(tx *sql.Tx, token string) int64 {
	//删除连接该token的连接
	delStmt, _ := tx.Prepare("delete from im_conn where token=?")
	defer delStmt.Close()
	res, err := delStmt.Exec(token)
	if err != nil {
		tx.Rollback()
		log.Println("删除用户连接错误:", err)
		return 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println("读取删除用户连接影响行数错误:", err)
		return 0
	}
	return num
}

/*
 添加连接
 */
func AddConn(key string, userId string, token string) int64 {
	insertStmt, _ := Database.Prepare("insert into im_conn VALUES (?, ?, ?, ?, ?)")
	defer insertStmt.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	res, err := insertStmt.Exec(key, userId, token, now, now)
	if err != nil {
		log.Println("保存用户连接错误:", err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		num = 0
		log.Println("读取保存用户连接影响行数错误:", err)

	}
	return num
}
