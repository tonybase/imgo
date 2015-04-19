package im
import (
    "code.google.com/p/go-uuid/uuid"
    "log"
    "im-go/im/model"
    "time"
)

// 检查账号是否存在
func CheckAccount(account string) int {
    var num int
    rows, err := Database.Query("select count(*)  from im_user where account=? ", account)

    if err != nil {
        log.Printf("根据账号查询用户错误: ", err)
    }
    defer rows.Close()
    for rows.Next() {
        rows.Scan(&num)
    }
    return num

}

// 登录账号
func LoginUser(account string, password string) model.IMUser {
    var user model.IMUser
    rows, err := Database.Query("select id,nick,status,sign,avatar from im_user where account=? and password=? ", account, password)
    if err != nil {
        log.Printf("根据账号及密码查询用户错误: ", err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&user.Id, &user.Nick, &user.Status, &user.Sign, &user.Avatar)
        if err != nil {
            log.Printf("根据账号及密码查询结果映射至对象错误:", err)
        }
    }
    return user
}

// 保存登录状态
func SaveLogin(userId string, token string, ip string) int64 {
    insStmt, _ := Database.Prepare("insert into im_login (id,user_id,token,login_date,login_ip) VALUES (?, ?, ?, ?,?)")
    defer insStmt.Close()
    res, err := insStmt.Exec(uuid.New(), userId, token, time.Now().Format("2006-01-02 15:04:05"), ip)
    if err != nil {
        log.Printf("保存用户登录记录错误:", err)
        return 0
    }
    num, err := res.RowsAffected()
    if err != nil {
        log.Printf("读取保存用户登录记录影响行数错误:", err)
        return 0
    }
    return num
}

// 获取好友分类列表
func GetCategories(account string) []model.IMCategory {

    return nil
}