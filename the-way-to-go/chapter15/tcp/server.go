package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatal(err)
	}
	// serve
	for {
		conn, err := lis.Accept() // 接收一个连接
		if err != nil {
			log.Fatal(err)
		}
		go doServerStuff(conn) // 处理连接
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf) // 读数据到缓冲区
		if err == io.EOF {
			fmt.Println("conn " + conn.RemoteAddr().String() + " disconnect")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(conn.RemoteAddr().String()+" Received data: %s\n", buf[:len])
	}
}
