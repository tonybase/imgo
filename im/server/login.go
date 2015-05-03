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
	http.HandleFunc("/user_relation", handleUserRelation)
	http.HandleFunc("/user_category", handleUserCategory)
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

// 登录请求处理方法
func handleLogin(resp http.ResponseWriter, req *http.Request) {
	// POST登录请求
	if req.Method == "POST" {
		ip := util.GetIp(req)
		account := req.FormValue("account")
		password := req.FormValue("password")

		log.Printf("ip %s", ip)
		log.Printf("account %s", account)
		log.Printf("password %s", password)

		login(resp, account, password, ip)
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

// 添加好友关系
func handleUserCategory(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		userId := req.FormValue("user_id")
		name := req.FormValue("name")

		if userId == "" {
			resp.Write(common.NewIMResponseSimple(101, "用户ID不能为空", "").Encode())
		} else if name == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别名称不能为空", "").Encode())
		} else {
			if model.AddCategory(userId, name) > 0 {
				resp.Write(common.NewIMResponseSimple(0, "添加分类成功", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "添加分类失败", "").Encode())
			}
		}
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

// 添加好友关系
func handleUserRelation(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		userId := req.FormValue("user_id")
		categoryId := req.FormValue("category_id")
		if userId == "" {
			resp.Write(common.NewIMResponseSimple(101, "用户ID不能为空", "").Encode())
		} else if categoryId == "" {
			resp.Write(common.NewIMResponseSimple(102, "类别ID不能为空", "").Encode())
		} else {
			if model.AddFriendRelation(userId, categoryId) > 0 {
				resp.Write(common.NewIMResponseSimple(0, "已建立好友关系", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(103, "建立好友关系失败", "").Encode())
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
		var user model.User
		num := model.CheckAccount(account)
		if num > 0 {
			user = model.LoginUser(account, password)
			if !strings.EqualFold(user.Id, "") {
				token := uuid.New()
				if model.SaveLogin(user.Id, token, ip) > 0 {
					returnData := make(map[string]string)
					returnData["id"] = user.Id
					returnData["nick"] = user.Nick
					returnData["avatar"] = user.Avatar
					returnData["status"] = user.Status
					returnData["token"] = token //token uuid 带 横杠
					resp.Write(common.NewIMResponseData(util.SetData("user", returnData), "LOGIN_RETURN").Encode())
				} else {
					resp.Write(common.NewIMResponseSimple(105, "保存登录记录错误,请稍后再试", "").Encode())
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
		num := model.CheckAccount(account)
		if num > 0 {
			resp.Write(common.NewIMResponseSimple(104, "用户名已存在", "").Encode())
		} else {
			num := model.SaveUser(account, password, nick, avatar)
			if (num > 0) {
				resp.Write(common.NewIMResponseSimple(0, "注册成功", "").Encode())
			} else {
				resp.Write(common.NewIMResponseSimple(105, "注册失败", "").Encode())
			}
		}
	}
}
