package model
import (
	"log"
	"code.google.com/p/go-uuid/uuid"
	"time"
)

type Login struct {
	Id        string    `json:"id"`         // id
	UserId    string    `json:"user_id"`    // 用户ID
	Token     string    `json:"token"`      // 用户TOKEN
	LoginAt   time.Time `json:"login_at"`   // 登录日期
	LoginIp   string    `json:"login_ip"`   // 登录IP
}

/*
 根据token获取用户登录
 */
func GetLoginByToken(token string) Login {
	var login Login
	row := Database.QueryRow("select id, user_id, token, login_at, login_ip from im_login where token=?", token)
	err := row.Scan(&login.Id, &login.UserId, &login.Token, &login.LoginAt, &login.LoginIp)
	if err != nil {
		log.Println("根据Token获取用户登录错误", err)
	}
	return login

}

/*
 保存登录状态
 */
func SaveLogin(userId string, token string, ip string) int64 {
	insStmt, _ := Database.Prepare("insert into im_login (id, user_id, token, login_at, login_ip) VALUES (?, ?, ?, ?,?)")
	defer insStmt.Close()
	res, err := insStmt.Exec(uuid.New(), userId, token, time.Now().Format("2006-01-02 15:04:05"), ip)
	if err != nil {
		log.Printf("保存用户登录记录错误: ", err)
		return 0
	}
	num, err := res.RowsAffected()
	if err != nil {
		log.Printf("读取保存用户登录记录影响行数错误:", err)
		return 0
	}
	return num
}