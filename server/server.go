package main

import(
	"fmt"
	"bufio"
	"net"
	"strings"
	"log"
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
		fmt.Println("new client connection")
		go handleConnection(conn)
	}
}

func handleConnection(client net.Conn) {

	for {

		message, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("From client: ", message)
		echo := strings.ToUpper(message)

		client.Write([]byte(echo + "\n"))
	}
}

