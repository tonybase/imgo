package main

import (
	"flag"
	"imgo/im/util"
	"log"
	"imgo/im/model"
	"imgo/im/server"
)

const (
	Name    string = "Go-IM"
	Version string = "1.0"
)

/*
启动服务方法
*/
func Start(config *util.IMConfig) {
	var err error
	//初始化model包下全局变量值
	model.Config = config
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

func main() {
	log.Println("*********************************************")
	log.Printf("           系统:[%s]版本:[%s]", Name, Version)
	log.Println("*********************************************")
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()
	// 读取配置信息
	config, err := util.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("读取配置文件错误: %s", err)
	}

	Start(config)
}