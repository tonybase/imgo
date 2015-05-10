package model

import (
	"code.google.com/p/go-uuid/uuid"
	"log"
	"time"
)

type Conversation struct {
	Id         string        `json:"id"`
	Creator    string        `json:"creator"`
	Receiver   string        `json:"receiver"`
	Type       string        `json:"type"`
	Create_at  time.Time    `json:"create_at"`
}

/*
 创建会话
 */
func AddSession(sender string, receiver string) string {
	insertStmt, _ := Database.Prepare("insert into `im_session` VALUES (?, ?, ?, ?, ?)")
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

func GetSession(sender string, receiver string) Conversation {
	var conv Conversation
	rows, err := Database.Query("select * from im_session where creator=? and receiver=? ", sender, receiver)
	if err != nil {
		log.Printf("根据账号及密码查询用户错误: ", err)
	}
	r, _ := rows.Columns();
	log.Println(r)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&conv.Id, &conv.Creator, &conv.Receiver, &conv.Type, &conv.Create_at)
		if err != nil {
			log.Printf("根据账号及密码查询结果映射至对象错误:", err)
		}
	}
	return conv
}

/*
 根据ID获取会话
 */
func GetSessionById(id string) (*Conversation, error) {
	var conv Conversation
	row := Database.QueryRow("select * from im_session where id=?", id)
	err := row.Scan(&conv.Id, &conv.Creator, &conv.Receiver, &conv.Type, &conv.Create_at)
	if err != nil {
		return nil, &DatabaseError{"根据ID获会话错误"}
	}
	return &conv, nil

}

/*
 根据ticket获取会话
 */
func GetReceiverKeyByTicket(ticket string) ([]string, error) {
	var keys []string
	rows, err := Database.Query("select c1.`id` from im_conn c1 left join im_session c2 on c1.user_id=c2.receiver where c2.id=?", ticket)
	if err != nil {
		return nil, &DatabaseError{"根据Ticket获取接收者Key和发送者ID错误"}
	}
	for rows.Next() {
		var key string
		rows.Scan(&key)

		keys = append(keys, key)
	}
	return keys, nil
}
