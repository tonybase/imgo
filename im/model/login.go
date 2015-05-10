package model
import (
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
func GetLoginByToken(token string) (*Login, error) {
	var login Login
	row := Database.QueryRow("select id, user_id, token, login_at, login_ip from im_login where token=?", token)
	err := row.Scan(&login.Id, &login.UserId, &login.Token, &login.LoginAt, &login.LoginIp)
	if err != nil {
		return nil, &DatabaseError{"根据Token获取用户登录错误"}
	}
	return &login, nil
}

/*
 保存登录状态
 */
func SaveLogin(userId string, token string, ip string) (*string, error) {
	insStmt, errStmt := Database.Prepare("insert into im_login (id, user_id, token, login_at, login_ip) VALUES (?, ?, ?, ?,?)")
	if errStmt != nil {
		return nil, &DatabaseError{"保存用户登录记录错误，数据库语句错误"}
	}
	defer insStmt.Close()
	id := uuid.New();
	_, err := insStmt.Exec(id, userId, token, time.Now().Format("2006-01-02 15:04:05"), ip)
	if err != nil {
		return nil, &DatabaseError{"保存用户登录记录错误"}
	}
	return &id, nil
}