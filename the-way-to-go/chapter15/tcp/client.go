package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:50000") // 拨号服务器
	if err != nil {
		log.Fatal(err)
	}

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if input == "Q" {
			return
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Println(err)
		}
	}
}
