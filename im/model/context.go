package model

import (
	"database/sql"
	"im-go/im/util"
)

var (
	Database *sql.DB = nil //数据库操作对象
	Config *util.IMConfig
)
