package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//申明一个map到时候存取配置文件
var per map[string]interface{}

func main() {
	//实例化这个map
	per = make(map[string]interface{})
	//打开这个ini文件
	f, _ := os.Open("test.ini")
	//读取文件到buffer里边
	buf := bufio.NewReader(f)
	for {
		//按照换行读取每一行
		l, err := buf.ReadString('\n')
		//相当于PHP的trim
		line := strings.TrimSpace(l)
		//判断退出循环
		if err != nil {
			if err != io.EOF {
				//return err
				panic(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		//匹配[db]然后存储
		case line[0] == '[' && line[len(line)-1] == ']':
			section := strings.TrimSpace(line[1 : len(line)-1])
			fmt.Println(section)
		default:
			//dnusername = xiaowei 这种的可以匹配存储
			i := strings.IndexAny(line, "=")
			per[strings.TrimSpace(line[0:i])] = strings.TrimSpace(line[i+1:])

		}
	}
	//循环输出结果
	for k, v := range per {
		fmt.Println(k, v)
	}
}
