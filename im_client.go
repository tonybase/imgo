package main

import (
    "im-go/im"
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
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)

    client := im.CreateClient(conn)

    go func() {
        for {
            msg := client.GetIncoming()
            out.Write(msg.Encode())
            out.WriteString("\n")
            out.Flush()
        }
    }()

    for {
        line, _, _ := in.ReadLine()
        msg := new(im.IMResponse)
        msg.Data = string(line)
        client.PutOutgoing(msg)
    }

}