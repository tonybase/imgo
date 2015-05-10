package model

import (
	"encoding/json"
	"code.google.com/p/go-uuid/uuid"
	"time"
)

/*
分组对象
*/
type Category struct {
	Id       string    `json:"id"`        // 分组ID
	Name     string    `json:"name"`      // 分组名称
	Creator  string    `json:"creator"`   // 分组名称
	CreateAt time.Time `json:"create_at"` // 创建时间
	Buddies  []User    `json:"buddies"`   // 好友列表
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
func GetCategoriesByToken(token string) ([]Category, error) {
	var categories []Category
	rows, err := Database.Query("select g.id, g.name from im_category g left join im_login l on l.user_id=g.creator where token=?", token)
	if (err != nil) {
		return nil, &DatabaseError{"根据Token获取好友分类错误"}
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	return categories, nil
}

/*
 根据UserId获取分组数据
 */
func GetCategoriesByUserId(token string) ([]Category, error) {
	var categories []Category
	rows, err := Database.Query("select * from im_category where creator=?", token)
	if (err != nil) {
		return nil, &DatabaseError{"根据用户ID获取好友分类错误"}
	}
	defer rows.Close()
	for rows.Next() {
		var category Category
		rows.Scan(&category.Id, &category.Name, &category.Creator, &category.CreateAt)
		categories = append(categories, category)
	}
	return categories, nil
}

/*
 添加好友分类
 */
func AddCategory(userId string, name string) (*string, error) {
	insStmt, err := Database.Prepare("insert into im_category (id, name, creator, create_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, &DatabaseError{"保存好友分类记录错误"}
	}
	defer insStmt.Close()
	id := uuid.New()
	_, err = insStmt.Exec(id, name, userId, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, &DatabaseError{"保存好友分类记录错误"}
	}
	return &id, nil;
}
