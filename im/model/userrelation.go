package model
import (
	"time"
)

type UserRelation struct {
	UserId string        `json:"user_id"`
	CategoryId string    `json:"category_id"`
	CreateAt time.Time   `json:"create_at"`
}

/**
 添加好友关系数据库
 */
func AddFriendRelation(userId string, categoryId string) (int64, error) {
	insStmt, err := Database.Prepare("insert into im_relation_user_category (user_id, category_id, create_at) VALUES (?, ?, ?)")
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