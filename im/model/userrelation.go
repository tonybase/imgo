package model

import (
	"database/sql"
	"time"
)

type UserRelation struct {
	UserId     string    `json:"user_id"`
	CategoryId string    `json:"category_id"`
	CreateAt   time.Time `json:"create_at"`
}

/**
添加好友关系数据库
*/
func AddFriendRelation(tx *sql.Tx, userId string, categoryId string) (int64, error) {
	insStmt, err := tx.Prepare("insert into im_relation_user_category (user_id, category_id, create_at) VALUES (?, ?, ?)")
	if err != nil {
		return -1, &DatabaseError{"添加好友关系数据库处理错误"}
	}
	defer insStmt.Close()
	res, err := insStmt.Exec(userId, categoryId, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return -1, &DatabaseError{"保存好友分类记录错误"}
	}
	num, err := res.RowsAffected()
	if err != nil {
		return -1, &DatabaseError{"读取保存好友分类记录影响行数错误"}
	}
	return num, nil
}

/**
删除好友关系数据库
*/
func DelFriendRelation(userId string, categoryId string) (int64, error) {
	delStmt, err := Database.Prepare("delete from im_relation_user_category where user_id=? and category_id=? ")
	if err != nil {
		return -1, &DatabaseError{"删除好友关系数据库处理错误"}
	}
	defer delStmt.Close()
	res, err := delStmt.Exec(userId, categoryId)
	if err != nil {
		return -1, &DatabaseError{"删除好友关系记录错误"}
	}
	num, err := res.RowsAffected()
	if err != nil {
		return -1, &DatabaseError{"读取删除好友关系记录影响行数错误"}
	}
	return num, nil
}
