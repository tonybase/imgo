package main

import (
	"bufio"
	"im-go/im"
	"im-go/im/common"
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
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	client := common.CreateClient("test", conn)

	go func() {
		for {
			msg := client.GetIn()
			out.Write(msg.Encode())
			out.WriteString("\n")
			out.Flush()
		}
	}()

	for {
		line, _, _ := in.ReadLine()
		msg := new(common.IMResponse)
		msg.Data = string(line)
		client.PutOut(msg)
	}

}
