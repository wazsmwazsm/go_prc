package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var conn net.Conn
	var err error
	var inputReader *bufio.Reader
	var input string
	var clientName string

	conn, err = net.Dial("tcp", ":4444")
	checkError(err)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, err = inputReader.ReadString('\n')
	checkError(err)
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to shutdown server.")
		input, _ = inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")

		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error()) // terminate program
	}
}
