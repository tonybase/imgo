package common

/*
 全局常量
*/
const (
	GET_KEY_RETURN        = "GET_KEY_RETURN"        // 请求TCP获取连接KEY
	GET_CONN              = "GET_CONN"              // 建立TCP长连接
	GET_CONN_RETURN       = "GET_CONN_RETURN"       // 获取连接返回
	GET_BUDDY_LIST        = "GET_BUDDY_LIST"        // 获取好友列表
	GET_BUDDY_LIST_RETURN = "GET_BUDDY_LIST_RETURN" // 获取好友列表返回
	CREATE_SESSION        = "CREATE_SESSION"        // 创建会话
	CREATE_SESSION_RETURN = "CREATE_SESSION_RETURN" // 创建会话返回
	SEND_MSG              = "SEND_MSG"              // 发送消息
	SEND_MSG_RETURN       = "SEND_MSG_RETURN"       // 发送消息返回
	PUSH_MSG              = "PUSH_MSG"              // 接收消息
	SEND_STATUS_CHANGE    = "SEND_STATUS_CHANGE"    // 发送状态
	PUSH_STATUS_CHANGE    = "PUSH_STATUS_CHANGE"    // 接收状态
	LOGOUT_REQUEST        = "LOGOUT_REQUEST"        // 退出
	UNAUTHORIZED          = "UNAUTHORIZED"          // 未授权
	SEND_BUDDY_REQUEST    = "SEND_BUDDY_REQUEST"    // 发送好友请求
	PUSH_BUDDY_REQUEST    = "PUSH_BUDDY_REQUEST"    // 接收好友请求
	ADD_BUDDY             = "ADD_BUDDY"
	FORMAT_DATE           = "2006-01-02"
	FORMAT_DATETIME       = "2006-01-02 15:04:05"
	FORMAT_DATETIME_ZONE  = "2006-01-02T15:04:05+08:00"
)
