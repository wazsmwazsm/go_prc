package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var userMap map[string]int

func main() {
	var lis net.Listener
	var err error
	var conn net.Conn
	userMap = make(map[string]int)
	fmt.Println("start server...")

	lis, err = net.Listen("tcp", ":4444")
	checkError(err)

	for {
		conn, err = lis.Accept()
		checkError(err)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var err error

	for {
		buf = make([]byte, 512)
		_, err = conn.Read(buf)
		checkError(err)
		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}

		if strings.Contains(input, ": WHO") {
			DisplayList()
		}

		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		userMap[string(clName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error()) // terminate program
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, user := range userMap {
		fmt.Printf("User %s is %d\n", key, user)
	}
	fmt.Println("--------------------------------------------")
}
