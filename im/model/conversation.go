package model

import (
	"code.google.com/p/go-uuid/uuid"
	"im-go/im/util"
	"log"
	"time"
)

/*
创建会话
*/
func AddConversation(sender string, receiver string, token string) string {
	insertStmt, _ := Database.Prepare("insert into `im_conversation` VALUES (?, ?, ?, ?,?,?)")
	defer insertStmt.Close()
	id := uuid.New()
	res, err := insertStmt.Exec(id, sender, time.Now().Format("2006-01-02 15:04:05"), receiver, "0", token)
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
func GetConversationById(id string) map[string]string {
	var data map[string]string
	res, err := Database.Query("select * from im_conversation where id=?", id)
	if err != nil {
		log.Println("根据ID获会话错误", err)
	} else {
		data = util.ResToMap(res)
	}
	return data

}
func GetReceiverKeyByTicket(ticket string) string {
	var key string
	err := Database.QueryRow("select c1.`key` from im_conn c1 left join im_conversation c2 on c1.user_id=c2.receiver where c2.id=?", ticket).Scan(&key)
	if err != nil {
		log.Println("根据Ticket获取接收者Key和发送者ID错误:", err)
	}
	return key
}
