package im

import (
	"flag"
	"im-go"
	"testing"
)

func TestConfig(t *testing.T) {
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()

	config, err := im.ReadConfig(*configPath)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log("IMPort", config.IMPort)
	t.Log("HttpPort", config.HttpPort)
	t.Log("MaxClients", config.MaxClients)
	t.Log("db.Host", config.DBConfig.Host)
	t.Log("db.Username", config.DBConfig.Username)
	t.Log("db.Password", config.DBConfig.Password)
	t.Log("db.Name", config.DBConfig.Name)
	t.Log("db.MaxIdleConns", config.DBConfig.MaxIdleConns)
	t.Log("db.MaxOpenConns", config.DBConfig.MaxOpenConns)

	db, err := config.DBConfig.Connect()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("获取到得数据库连接:", db)
	defer db.Close()
}
