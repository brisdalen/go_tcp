package main

import(
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Client made. \n Attempting to connect..")
	conn, err := net.Dial("tcp", "142.93.135.21:1234")
	if err != nil {
		fmt.Println("Error occured while connecting")
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, text + "\n")

	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
