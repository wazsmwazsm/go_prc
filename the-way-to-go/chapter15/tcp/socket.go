package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	var (
		host   = "www.apache.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]byte, 4096)
		read   = true
		count  = 0
	)
	conn, err := net.Dial("tcp", remote)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	// 发送一个 http 请求
	io.WriteString(conn, msg)
	// 读取响应
	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Printf("%s\n", data[0:count])
	}
	conn.Close()
}
