package model

import (
	"database/sql"
	"time"
)

/*
 连接对象
*/
type Conn struct {
	Key      string    `json:"key"`       // 连接key 唯一标识符
	UserId   string    `json:"id"`        // 用户ID
	Token    string    `json:"token"`     // token
	CreateAt time.Time `json:"create_at"` // 时间
	UpdateAt time.Time `json:"update_at"` // 时间
}

/*
 根据Key获取连接的数量
*/
func CountConnByKey(key string) (int64, error) {
	var num int64
	err := Database.QueryRow("select count(*) from im_conn where id=?", key).Scan(&num)
	if err != nil {
		return -1, &DatabaseError{"根据Token获取连接数量错误"}
	}
	return num, nil
}

/*
 根据Token获取连接的数量
*/
func CountConnByToken(token string) (int64, error) {
	var num int64
	err := Database.QueryRow("select count(*) from im_conn where token=?", token).Scan(&num)
	if err != nil {
		return -1, &DatabaseError{"根据Token获取连接数量错误"}
	}
	return num, nil
}

/*
 根据Token获取连接的数量
*/
func CountConnByUserId(userId string) (int64, error) {
	var num int64
	err := Database.QueryRow("select count(*) from im_conn where user_id=?", userId).Scan(&num)
	if err != nil {
		return -1, &DatabaseError{"根据Token获取连接数量错误"}
	}
	return num, nil
}

/*
 根据用户ID修改连接
*/
func UpdateConnByToken(key string, userId string, token string) (int64, error) {
	updateStmt, err := Database.Prepare("UPDATE im_conn SET `id`=?, `user_id` = ?, update_at=? WHERE token =?")
	if err != nil {
		return -1, &DatabaseError{"读取修改用户连接影响行数错误"}
	}
	defer updateStmt.Close()
	res, err := updateStmt.Exec(key, userId, time.Now().Format("2006-01-02 15:04:05"), token)
	if err != nil {
		return -1, &DatabaseError{"更新用户连接错误"}
	}
	num, err := res.RowsAffected()
	if err != nil {
		return -1, &DatabaseError{"读取更新用户连接影响行数错误"}
	}
	return num, nil
}

/*
 根据token获取token
*/
func GetConnByToken(token string) (*Conn, error) {
	var conn Conn
	row := Database.QueryRow("select * from im_conn where token=?", token)
	err := row.Scan(&conn.Key, &conn.UserId, &conn.Token, &conn.CreateAt, &conn.UpdateAt)
	if err != nil {
//		log.Println("根据Token获取用户连接记录错误", err)
		return nil, &DatabaseError{"根据Token获取用户连接记录错误"}
	}
	return &conn, nil
}

/*
 根据用户ID获取连接
*/
func GetConnByUserId(userId string) (*Conn, error) {
	var conn Conn
	row := Database.QueryRow("select * from im_conn where user_id=?", userId)
	err := row.Scan(&conn.Key, &conn.UserId, &conn.Token, &conn.CreateAt, &conn.UpdateAt)
	if err != nil {
		return nil, &DatabaseError{"根据用户ID获取用户连接记录错误"}
	}
	return &conn, nil
}

/*
 根据token删除连接
*/
func DeleteConnByKey(key string) error {
	//删除连接该token的连接
	delStmt, err := Database.Prepare("delete from im_conn where id=?")
	if err != nil {
		return &DatabaseError{"删除用户连接错误"}
	}
	defer delStmt.Close()
	_, err = delStmt.Exec(key)
	if err != nil {
		return &DatabaseError{"删除用户连接错误"}
	}
	return nil
}

/*
 根据token删除连接
*/
func DeleteConnByToken(tx *sql.Tx, token string) error {
	//删除连接该token的连接
	delStmt, err := tx.Prepare("delete from im_conn where token=?")
	if err != nil {
		return &DatabaseError{"删除用户连接错误"}
	}
	defer delStmt.Close()
	_, err = delStmt.Exec(token)
	if err != nil {
		tx.Rollback()
		if err != nil {
			return &DatabaseError{"删除用户连接错误"}
		}
	}
	return nil
}

/*
 添加连接
*/
func AddConn(key string, userId string, token string) (*string, error) {
	insertStmt, err := Database.Prepare("insert into im_conn VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, &DatabaseError{"保存用户连接错误"}
	}
	defer insertStmt.Close()
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err = insertStmt.Exec(key, userId, token, now, now)
	if err != nil {
		return nil, &DatabaseError{"保存用户连接错误"}
	}
	return &key, nil
}
