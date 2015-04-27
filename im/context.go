package im

import (
	"im-go/im/model"
	"im-go/im/server"
	"im-go/im/util"
	"log"
)


/*
启动服务方法
*/
func Start(config *util.IMConfig) {

	var err error
	//初始化model包下全局变量值
	model.Config=config
	model.Database, err = config.DBConfig.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer model.Database.Close()
	go func() {
		err := server.StartHttpServer(*config)
		log.Fatalf("HttpServer", err)
	}()
	// 启动IM服务
	server.StartIMServer(*config)
}
