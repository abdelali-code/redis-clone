package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("server listenning on port 6379")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("the the addrssing of the connection is ", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error while reading data")
		}
		cleanMessage := strings.TrimSpace(message)
		fmt.Println(cleanMessage)
		if strings.Compare(cleanMessage, "PING") == 0 {
			_, err = conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

}
