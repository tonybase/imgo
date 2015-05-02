package model
import (
	"time"
	"log"
)

type UserRelation struct {
	UserId string        `json:"user_id"`
	CategoryId string    `json:"category_id"`
	CreateAt time.Time   `json:"create_at"`
}

func AddFriendRelation(userId string, categoryId string) int64 {
	insStmt, _ := Database.Prepare("insert into im_relation_user_category (user_id, category_id, create_at) VALUES (?, ?, ?)")
	defer insStmt.Close()
	res, err := insStmt.Exec(userId, categoryId, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Printf("保存好友分类记录错误: ", err)
		return 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Printf("读取保存好友分类记录影响行数错误:", err)
		return 0
	}
	return num
}