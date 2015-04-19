package main

import (
	"flag"
	"im-go/im"
	"log"
	"im-go/im/common"
)

const (
	Name    string = "GO-IM"
	Version string = "1.0"
)

func main() {
	//打印系统信息
	log.Println(Name, "- Version", Version)

	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()

	// 读取配置信息
	config, err := common.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("读取配置文件错误: %s", err)
	} else {
		log.Println(config)
	}

	im.Start(config)

}
