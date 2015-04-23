package common

import (
	"net"
	"net/http"
	"strings"
)

/*
获取IP
*/
func GetIp(r *http.Request) string {
	ip := net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]).String()
	if ip == "<nil>" {
		ip = "127.0.0.1"
	}
	return ip
}

/*
组合数据 原转JSON(已修正不需要转JSON)
*/
//FIXME 此处方法名需要重新命名 否则会产生干扰
func SetData(key string, data interface{}) interface{} {
	datamap := make(map[string]interface{})
	datamap[key] = data
	return datamap
}
