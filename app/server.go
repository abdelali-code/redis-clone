package main

import (
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
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			fmt.Println("error while reading data")
		}
		buf = append(buf, tmp[:n]...)
		fmt.Println("the message is ", buf)
		cleanMessage := string(buf)
		fmt.Println(cleanMessage)
		leng := len(cleanMessage)
		start := 0
		strLen := len("*1\r\n$4\r\nPING\r\n")
		fmt.Println(leng)
		for leng >= strLen {
			if strings.Compare(cleanMessage[start:strLen], "*1\r\n$4\r\nPING\r\n") == 0 {
				_, err = conn.Write([]byte("+PONG\r\n"))
				if err != nil {
					os.Exit(1)
				}
			}
			start += strLen
			leng -= strLen
		}
	}

}
