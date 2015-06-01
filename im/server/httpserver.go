package server

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"fmt"
	"imgo/im/common"
	"imgo/im/model"
	"imgo/im/util"
	"log"
	"net/http"
	"strings"
	"time"
	"strconv"
)

// 启动HTTP服务
func StartHttpServer(config util.IMConfig) error {
	log.Printf("Http服务器启动中...")

	// 设置请求映射地址及对应处理方法
	http.HandleFunc("/system", handleSystem)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/query", handleQuery)
	http.HandleFunc("/users/relation/add", handleUserRelationAdd)
	http.HandleFunc("/users/relation/del", handleUserRelationDel)
	http.HandleFunc("/users/relation/push", handleUserRelationPush)
	http.HandleFunc("/users/relation/refuse", handleUserRelationRefuse)
	http.HandleFunc("/users/category/add", handleUserCategoryAdd)
	http.HandleFunc("/users/category/del", handleUserCategoryDel)
	http.HandleFunc("/users/category/edit", handleUserCategoryEdit)
	http.HandleFunc("/users/category/query", handleUserCategoryQuery)
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

// 系统状态信息
func handleSystem(resp http.ResponseWriter, req *http.Request) {
	mem, _ := mem.VirtualMemory()
	cpuNum, _ := cpu.CPUCounts(true);
	cpuInfo, _ := cpu.CPUPercent(10 * time.Microsecond, true);

	data := make(map[string]interface{})
	data["im.conn"] = len(ClientMaps)
	data["mem.total"] = fmt.Sprintf("%vMB", mem.Total/1024/1024)
	data["mem.free"] = fmt.Sprintf("%vMB", mem.Free/1024/1024)
	data["mem.used_percent"] = fmt.Sprintf("%s%%", strconv.FormatFloat(mem.UsedPercent, 'f', 2, 64))
	data["cpu.num"] = cpuNum
	data["cpu.info"] = cpuInfo

	resp.Write(common.NewIMResponseData(data, "").Encode())
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
	nick := req.FormValue("nick")
	users, err := model.QueryUser("nick", "like", nick)
	if err == nil {
		resp.Write(common.NewIMResponseData(util.SetData("users", users), "").Encode())
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
func handleUserCategoryQuery(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	categories, err := model.GetCategoriesByUserId(id)
	if err != nil {
		resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
	} else {
		resp.Write(common.NewIMResponseData(util.SetData("categories", categories), "").Encode())
	}
}

// 添加好友关系
func handleUserRelationAdd(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		receiver_category_id := req.FormValue("receiver_category_id")
		buddy_request_id := req.FormValue("buddy_request_id")
		buddyrequest, _ := model.GetBuddyRequestById(buddy_request_id)
		if buddyrequest != nil {
			receiver := buddyrequest.Receiver
			sender := buddyrequest.Sender
			sender_category_id := buddyrequest.SenderCategoryId
			//开启事务
			tx, _ := model.Database.Begin()
			//修改好友请求记录中接受人的好友分组ID
			_, err := model.UpdateBuddyRequestReceiverCategoryId(tx, buddy_request_id, receiver_category_id)
			//添加请求人好友关系数据
			_, err = model.AddFriendRelation(tx, receiver, sender_category_id)
			//添加接收人好友关系数据
			_, err = model.AddFriendRelation(tx, sender, receiver_category_id)
			//修改好友请求记录中状态
			_, err = model.UpdateBuddyRequestStatus(tx, buddy_request_id, "1")

			if err != nil {
				tx.Rollback()
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			} else {
				tx.Commit()
				//判断请求者是不是在线 在线就把接受者推送给请求者
				conn, _ := model.GetConnByUserId(sender)
				if conn != nil { //在线
					user, _ := model.GetUserById(receiver)
					data := make(map[string]interface{})
					data["category_id"] = sender_category_id
					data["user"] = user
					ClientMaps[conn.Key].PutOut(common.NewIMResponseData(util.SetData("user", data), common.ADD_BUDDY))
				}
				conn, _ = model.GetConnByUserId(receiver)
				if conn != nil {
					user, _ := model.GetUserById(sender)
					data := make(map[string]interface{})
					data["category_id"] = receiver_category_id
					data["user"] = user
					ClientMaps[conn.Key].PutOut(common.NewIMResponseData(util.SetData("user", data), common.ADD_BUDDY))
				}
				resp.Write(common.NewIMResponseSimple(0, "好友关系建立成功", "").Encode())
				return
			}

		} else {
			resp.Write(common.NewIMResponseSimple(104, "该好友请求不存在", "").Encode())
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
func handleUserRelationPush(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sender_category_id := req.FormValue("sender_category_id")
		sender := req.FormValue("sender")
		receiver := req.FormValue("receiver")
		if sender_category_id == "" {
			resp.Write(common.NewIMResponseSimple(101, "请选择分组", "").Encode())
		} else if sender == "" {
			resp.Write(common.NewIMResponseSimple(102, "请重新登录", "").Encode())
		} else {
			//判断接收人是不是在线 在线直接推送，不在线记录至请求表中
			conn, _ := model.GetConnByUserId(receiver)
			user, _ := model.GetUserById(sender)
			buddyRequestId, err := model.AddBuddyRequest(sender, sender_category_id, receiver)
			if err != nil {
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
			} else {
				if conn != nil { //在线 直接推送 不在线 客户登录时候会激活请求通知
					data := make(map[string]interface{})
					data["id"] = user.Id
					data["nick"] = user.Nick
					data["status"] = user.Status
					data["sign"] = user.Sign
					data["avatar"] = user.Avatar
					data["buddyRequestId"] = buddyRequestId
					ClientMaps[conn.Key].PutOut(common.NewIMResponseData(util.SetData("user", data), common.PUSH_BUDDY_REQUEST))
					resp.Write(common.NewIMResponseSimple(0, "发送好友请求成功", "").Encode())
				}
				resp.Write(common.NewIMResponseSimple(1, "发送好友请求成功", "").Encode())
				return
			}
		}
	} else {
		resp.Write(common.NewIMResponseSimple(404, "Not Found: "+req.Method, "").Encode())
	}
}

func handleUserRelationRefuse(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		buddy_request_id := req.FormValue("buddy_request_id")
		if buddy_request_id != "" {
			tx, _ := model.Database.Begin()
			//修改好友请求记录中状态
			_, err := model.UpdateBuddyRequestStatus(tx, buddy_request_id, "2")
			if err != nil {
				tx.Rollback()
				resp.Write(common.NewIMResponseSimple(100, err.Error(), "").Encode())
				return
			} else {
				tx.Commit()
				resp.Write(common.NewIMResponseSimple(0, "已经拒绝该好友请求成功", "").Encode())
				return
			}
		} else {
			resp.Write(common.NewIMResponseSimple(109, "该好友请求不合法", "").Encode())
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
