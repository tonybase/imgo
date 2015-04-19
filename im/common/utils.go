package common

import (
    "net"
    "net/http"
    "strings"
    "encoding/json"
)

//获取IP
func GetIp(r *http.Request) string {
    ip := net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]).String()
    if ip == "<nil>" {
        ip = "127.0.0.1"
    }
    return ip
}

func GetJson(key string, data interface{}) string {
    connData := make(map[string]interface{})
    dataJson, _ := json.Marshal(data)
    connData[key] = string(dataJson)
    dataJson2, _ := json.Marshal(connData)
    return string(dataJson2);
}