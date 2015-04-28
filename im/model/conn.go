package model

import (
	"im-go/im/util"
	"log"
	"time"
	"database/sql"
)

//连接模型
type IMConn struct {
	UserId string
	Token  string
	Key    string
	Date   time.Time
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

func UpdateConnByUserId(userId string, token string, key string) int64 {
	var num int64
	updateStmt, _ := Database.Prepare("UPDATE im_conn SET `token` = ? ,`key`=? ,date=? WHERE user_id =?")
	defer updateStmt.Close()
	res, err := updateStmt.Exec(token, key, time.Now().Format("2006-01-02 15:04:05"), userId)
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
func GetConnByToken(token string) map[string]string {
	var data map[string]string
	res, err := Database.Query("select * from im_conn where token=?", token)
	if err != nil {
		log.Println("根据Token获取用户连接记录错误", err)
	} else {
		data = util.ResToMap(res)
	}
	return data
}
func GetConnByUserId(id string) map[string]string {
	var data map[string]string
	res, err := Database.Query("select * from im_conn where user_id=?", id)
	if err != nil {
		log.Println("根据用户ID获取用户连接记录错误", err)
	} else {
		data = util.ResToMap(res)
	}
	return data
}
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
func AddConn(userId string, token string, key string) int64 {
	insertStmt, _ := Database.Prepare("insert into im_conn VALUES (?, ?, ?, ?)")
	defer insertStmt.Close()

	res, err := insertStmt.Exec(userId, token, key, time.Now().Format("2006-01-02 15:04:05"))
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
