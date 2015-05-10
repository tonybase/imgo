package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	in := bufio.NewReader(os.Stdin)

	go func() {
		for {
			if line, _, err := reader.ReadLine(); err == nil {
				log.Println(string(line))
			}
		}
	}()

	for {
		line, _, _ := in.ReadLine()
		// 模拟一个请求
		// {"command":"GET_CONN","data":null}
		// {"command":"GET_BUDDY_LIST","data":null}
		writer.WriteString(string(line) + "\n")
		writer.Flush()
	}

}
