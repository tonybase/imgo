package main

import (
	//"database/sql"
	//"flag"
	"fmt"
	"im-go/im"
)

func main() {
	data := im.GetLoginByToken("f727ae50-80fa-4739-9c0c-82b17f9c")
	fmt.Println("获取到得Login:", data)
}
