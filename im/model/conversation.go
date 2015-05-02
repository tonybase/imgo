package model

import (
	"code.google.com/p/go-uuid/uuid"
	"log"
	"time"
)

type Conversation struct {
	Id         string        `json:"id"`
	Creator    string        `json:"creator"`
	Create_at  time.Time    `json:"create_at"`
	Receiver   string        `json:"receiver"`
	Type       string        `json:"type"`
}

/*
 创建会话
 */
func AddConversation(sender string, receiver string) string {
	insertStmt, _ := Database.Prepare("insert into `im_conversation` VALUES (?, ?, ?, ?, ?)")
	defer insertStmt.Close()
	id := uuid.New()
	res, err := insertStmt.Exec(id, sender, receiver, "0", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("创建会话错误:", err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		num = 0
		log.Println("读取保存用户连接影响行数错误:", err)
	}
	if num == 0 {
		return ""
	}
	return id

}

/*
 根据ID获取会话
 */
func GetConversationById(id string) Conversation {
	var conv Conversation
	row := Database.QueryRow("select * from im_conversation where id=?", id)
	err := row.Scan(&conv.Id, &conv.Creator, &conv.Receiver, &conv.Type, &conv.Create_at)
	if err != nil {
		log.Println("根据ID获会话错误", err)
	}
	return conv

}

/*
 根据ticket获取会话
 */
func GetReceiverKeyByTicket(ticket string) string {
	var key string
	err := Database.QueryRow("select c1.`key` from im_conn c1 left join im_conversation c2 on c1.user_id=c2.receiver where c2.id=?", ticket).Scan(&key)
	if err != nil {
		log.Println("根据Ticket获取接收者Key和发送者ID错误:", err)
	}
	return key
}
