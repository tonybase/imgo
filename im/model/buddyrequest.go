package model

import (
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"log"
	"time"
)

type BuddyRequest struct {
	Id                 string    `json:"id"`               // ID
	Sender             string    `json:"sender"`           // 请求者
	SenderCategoryId   string    `json:"senderCategoryId"` // 请求者分组ID
	Receiver           string    `json:"receiver"`         // 接收者
	ReceiverCategoryId string    `json:"receiverCategoryId"`
	SendAt             time.Time `json:"sendAt"` // 请求时间
	Status             string    `json:"status"` // 状态
}

/*
 添加好友请求数据(好友未在线所以存好友请求表)
*/
func AddBuddyRequest(sender string, sender_cate_id string, receiver string) (*string, error) {
	insStmt, err := Database.Prepare("insert into im_buddy_request (id,sender,sender_category_id,receiver,send_at,status) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return nil, &DatabaseError{"保存好友请求错误"}
	}
	defer insStmt.Close()
	id := uuid.New()
	_, err = insStmt.Exec(id, sender, sender_cate_id, receiver, time.Now().Format("2006-01-02 15:04:05"), "0")
	if err != nil {
		log.Println(err)
		return nil, &DatabaseError{"保存好友请求错误"}
	}
	return &id, nil
}

/*
 根据receiver获取未读的好友请求
*/
func GetBuddyRequestsByReceiver(receiver string) ([]BuddyRequest, error) {
	var buddyRequests []BuddyRequest
	rows, err := Database.Query("select * from im_buddy_request where status='0' and receiver=?", receiver)
	if err != nil {
		return nil, &DatabaseError{"根据receiver获取未读的好友请求错误"}
	}
	defer rows.Close()
	for rows.Next() {
		var buddyRequest BuddyRequest
		rows.Scan(&buddyRequest.Id, &buddyRequest.Sender, &buddyRequest.SenderCategoryId, &buddyRequest.Receiver, &buddyRequest.ReceiverCategoryId, &buddyRequest.SendAt, &buddyRequest.Status)
		buddyRequests = append(buddyRequests, buddyRequest)
	}
	return buddyRequests, nil
}

/*
 根据ID获取未读的好友请求
*/
func GetBuddyRequestById(id string) (*BuddyRequest, error) {
	var buddyRequest BuddyRequest
	row := Database.QueryRow("select id,sender,sender_category_id,receiver from im_buddy_request where status='0' and id=?", id)
	err := row.Scan(&buddyRequest.Id, &buddyRequest.Sender, &buddyRequest.SenderCategoryId, &buddyRequest.Receiver)
	if err != nil {
		return nil, &DatabaseError{"根据ID查询好友请求-将结果映射至对象错误"}
	}
	return &buddyRequest, nil
}

/*
根据ID修改好友请求状态
*/
func UpdateBuddyRequestStatus(tx *sql.Tx, id string, status string) (int64, error) {
	var num int64
	updateStmt, err := tx.Prepare("update im_buddy_request SET `status` = ? WHERE id =?")
	if err != nil {
		return -1, &DatabaseError{"修改好友请求数据库处理错误"}
	}
	defer updateStmt.Close()
	res, err := updateStmt.Exec(status, id)
	if err != nil {
		return -1, &DatabaseError{"更新好友请求错误"}
	}
	num, err = res.RowsAffected()
	if err != nil {
		return -1, &DatabaseError{"读取修改好友请求影响行数错误"}
	}
	return num, nil
}
func UpdateBuddyRequestReceiverCategoryId(tx *sql.Tx, id string, receiver_category_id string) (int64, error) {
	var num int64
	updateStmt, err := tx.Prepare("update im_buddy_request SET `receiver_category_id` = ? WHERE id =?")
	if err != nil {
		return -1, &DatabaseError{"修改好友请求数据库处理错误"}
	}
	defer updateStmt.Close()
	res, err := updateStmt.Exec(receiver_category_id, id)
	if err != nil {
		return -1, &DatabaseError{"更新好友请求错误"}
	}
	num, err = res.RowsAffected()
	if err != nil {
		return -1, &DatabaseError{"读取修改好友请求影响行数错误"}
	}
	return num, nil
}
