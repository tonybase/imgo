package common

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

//获取IP
func GetIp(r *http.Request) string {
	ip := net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]).String()
	if ip == "<nil>" {
		ip = "127.0.0.1"
	}
	return ip
}

func GetJson(key string, data interface{}) interface{} {
	datamap := make(map[string]interface{})
	datamap[key] = data
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Println("转JSON错误:", err)
	}
	return dataJson
}
