package main

import (
	"flag"
	"im-go/im"
	"im-go/im/common"
	"log"
)

const (
	Name    string = "Go-IM"
	Version string = "1.0"
)

func main() {
	log.Println("*********************************************")
	log.Printf("           系统:[%s]版本:[%s]", Name, Version)
	log.Println("*********************************************")
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()
	// 读取配置信息
	config, err := common.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("读取配置文件错误: %s", err)
	}
	im.Start(config)
}
