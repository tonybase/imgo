package model

import (
	"encoding/json"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"log"
)

/*
分组对象
*/
type Category struct {
	Id      string   `json:"id"`      //分组ID
	Name    string   `json:"name"`    //分组名称
	Buddies []User   `json:"buddies"` //好友列表
}

/*
转JSON数据
*/
func (this *Category) Encode() []byte {
	s, _ := json.Marshal(*this)
	return s
}

/*
 解析JSON数据
 */
func (this *Category) Decode(data []byte) error {
	err := json.Unmarshal(data, this)
	return err
}

/*
 分组PUSH用户(仅传递数据 非操作数据库方法)
 */
func (this *Category) AddUser(u User) {
	this.Buddies = append(this.Buddies, u)
}

/*
 根据token获取分组数据
 */
func GetCategoriesByToken(token string) []Category {
	var categories []Category
	rows, _ := Database.Query("select g.id, g.name from im_category g left join im_login l on l.user_id=g.creator where token=?", token)
	defer rows.Close()
	for rows.Next() {
		var category Category
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	return categories
}

func AddCategory(userId string, name string) int64 {
	insStmt, _ := Database.Prepare("insert into im_category (id, name, creator, create_at) VALUES (?, ?, ?, ?)")
	defer insStmt.Close()
	res, err := insStmt.Exec(uuid.New(), name, userId, time.Now().Format("2006-01-02 15:04:05"))
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
