package server

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"im-go/im/common"
	"im-go/im/model"
	"im-go/im/util"
	"log"
	"net/http"
	"strings"
)

// 启动HTTP服务
func StartHttpServer(config util.IMConfig) error {
	log.Printf("Http服务器启动中...")

	// 设置请求映射地址及对应处理方法
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/query", handleQuery)
	http.HandleFunc("/users/relation/add", handleUserRelationAdd)
	http.HandleFunc("/users/relation/del", handleUserRelationDel)
	http.HandleFunc("/users/category/add", handleUserCategoryAdd)
	http.HandleFunc("/users/category/del", handleUserCategoryDel)
	http.HandleFunc("/users/category/edit", handleUserCategoryEdit)
	//打印监听端口
	log.Printf("Http服务器开始监听[%d]端口", config.HttpPort)
	log.Println("*********************************************")
	// 设置监听地址及端口
	addr := fmt.Sprintf("0.0.0.0:%d", config.HttpPort)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return fmt.Errorf("监听Http失败: %s", err)
	}
	return nil
}

// 注册请求
func handleRegister(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		account := req.FormValue("account")
		password := req.FormValue("password")
		nick := req.FormValue("nick")
		avatar := req.FormValue("avatar")

		register(resp, account, password, nick, avatar)
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

/**
登录请求处理方法
*/
func handleLogin(resp http.ResponseWriter, req *http.Request) {
	// POST登录请求
	if req.Method == "POST" {
		ip := util.GetIp(req)
		account := req.FormValue("account")
		password := req.FormValue("password")
		login(resp, account, password, ip)
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

/**
查询请求处理方法
*/
func handleQuery(resp http.ResponseWriter, req *http.Request) {
	// GET登录请求
	if req.Method == "GET" {
		nick := req.FormValue("nick")
		users, err := model.QueryUser("nick", "like", nick)
		if err != nil {
			resp.Write(common.NewIMResponseData(util.SetData("users", users), "").Encode())
		}
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

// 添加好友分类
func handleUserCategoryAdd(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		//获取好友列表
		userId := req.FormValue("user_id")

		categories, err := model.GetCategoriesByUserId(userId)
		if err != nil {
			resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
			return
		}
		categories, err = model.GetBuddiesByCategories(categories)
		if err != nil {
			resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
			return
		}
		resp.Write(common.NewIMResponseData(util.SetData("categories", categories), "").Encode())
	case "POST":
		// 添加好友列表
		userId := req.FormValue("user_id")
		name := req.FormValue("name")

		if userId == "" {
			resp.Write(common.NewIMResponseSimple(101, "用户ID不能为空", "").Encode())
		} else if name == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别名称不能为空", "").Encode())
		} else {
			_, err := model.AddCategory(userId, name)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(103, err.Error(), "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(0, "添加分类成功", "").Encode())
			}
		}
	default:
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())

	}
}

// 删除好友分类
func handleUserCategoryDel(resp http.ResponseWriter, req *http.Request) {
	categoryId := req.FormValue("category_id")
	switch req.Method {
	case "GET":
		if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别ID不能为空", "").Encode())
		} else {
			num, err := model.DelCategoryById(categoryId)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "已删除好友分类", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "删除好友分类失败", "").Encode())
			}
		}
	case "POST":
		if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别ID不能为空", "").Encode())
		} else {
			num, err := model.DelCategoryById(categoryId)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "已删除好友分类", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "删除好友关系分类", "").Encode())
			}
		}
	default:
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

// 编辑好友分类
func handleUserCategoryEdit(resp http.ResponseWriter, req *http.Request) {
	categoryId := req.FormValue("category_id")
	categoryName := req.FormValue("category_name")
	switch req.Method {
	case "GET":
		if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(101, "类别ID不能为空", "").Encode())
		} else if categoryName == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别名称不能为空", "").Encode())
		} else {
			num, err := model.EditCategoryById(categoryId, categoryName)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "修改用户好友类别成功", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "修改用户好友类别失败", "").Encode())
			}
		}
	case "POST":
		if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(101, "类别ID不能为空", "").Encode())
		} else if categoryName == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别名称不能为空", "").Encode())
		} else {
			num, err := model.EditCategoryById(categoryId, categoryName)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "修改用户好友类别成功", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "修改用户好友类别失败", "").Encode())
			}
		}
	default:
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())

	}
}

// 添加好友关系
func handleUserRelationAdd(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		userId := req.FormValue("user_id")
		categoryId := req.FormValue("category_id")
		if userId == "" {
			resp.Write(common.NewIMResponseSimple(101, "用户ID不能为空", "").Encode())
		} else if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别ID不能为空", "").Encode())
		} else {
			num, err := model.AddFriendRelation(userId, categoryId)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "已建立好友关系", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "建立好友关系失败", "").Encode())
			}
		}
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

// 删除好友关系
func handleUserRelationDel(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		userId := req.FormValue("user_id")
		categoryId := req.FormValue("category_id")
		if userId == "" {
			resp.Write(common.NewIMResponseSimple(101, "用户ID不能为空", "").Encode())
		} else if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别ID不能为空", "").Encode())
		} else {
			num, err := model.DelFriendRelation(userId, categoryId)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "已删除好友关系", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "删除好友关系失败", "").Encode())
			}
		}
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

// 登录主方法
func login(resp http.ResponseWriter, account string, password string, ip string) {
	if account == "" {
		resp.Write(common.NewIMResponseSimple(101, "账号不能为空", "").Encode())
	} else if password == "" {
		resp.Write(common.NewIMResponseSimple(102, "密码不能为空", "").Encode())
	} else {
		num, err := model.CheckAccount(account)
		if err != nil {
			resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
			return
		}
		if num > 0 {
			user, err := model.LoginUser(account, password)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			}
			if !strings.EqualFold(user.Id, "") {
				token := uuid.New()
				if _, err := model.SaveLogin(user.Id, token, ip); err != nil {
					resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				} else {
					// returnData := make(map[string]string)
					// returnData["id"] = user.Id
					// returnData["nick"] = user.Nick
					// returnData["avatar"] = user.Avatar
					// returnData["status"] = user.Status
					// returnData["token"] = token //token uuid 带 横杠
					// returnData["sign"]=user.Sign
					user.Token = token
					resp.Write(common.NewIMResponseData(util.SetData("user", user), "LOGIN_RETURN").Encode())
				}
			} else {
				resp.Write(common.NewIMResponseSimple(104, "密码错误", "").Encode())
			}
		} else {
			resp.Write(common.NewIMResponseSimple(103, "账号不存在", "").Encode())
		}
	}
}

/*
 用户注册
 101	账号不能为空
 102	密码不能为空
 103	用户名已存在
 104	昵称不能为空
 105	注册失败
*/
func register(resp http.ResponseWriter, account string, password string, nick string, avatar string) {
	if account == "" {
		resp.Write(common.NewIMResponseSimple(101, "账号不能为空", "").Encode())
	} else if password == "" {
		resp.Write(common.NewIMResponseSimple(102, "密码不能为空", "").Encode())
	} else if nick == "" {
		resp.Write(common.NewIMResponseSimple(103, "昵称不能为空", "").Encode())
	} else {
		num, err := model.CheckAccount(account)
		if err != nil {
			resp.Write(common.NewIMResponseSimple(103, err.Error(), "").Encode())
			return
		}
		if num > 0 {
			resp.Write(common.NewIMResponseSimple(104, "用户名已存在", "").Encode())
		} else {
			_, err := model.SaveUser(account, password, nick, avatar)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(104, err.Error(), "").Encode())
				return
			}
			if num > 0 {
				resp.Write(common.NewIMResponseSimple(0, "注册成功", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(105, "注册失败", "").Encode())
			}
		}
	}
}
