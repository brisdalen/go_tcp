package main

import(
	"fmt"
	"bufio"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error occured")
	}
	fmt.Println("Server started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error occured")
		}
		go handleConnection(conn)
	}
}

func handleConnection(client net.Conn) {

	message, _ := bufio.NewReader(client).ReadString('\n')

	echo := strings.ToUpper(message)

	client.Write([]byte(echo + "\n"))
}

