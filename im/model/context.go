package model

import (
	"database/sql"
	"im-go/im/util"
)

/*
包内上下文变量
*/
var (
	Database *sql.DB        = nil //数据库操作对象
	Config   *util.IMConfig       //配置对象
)
