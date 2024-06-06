package main



import (
	"fmt" 
	"io"
	"net"
)



func main() {

	listenner, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		fmt.Println("error while creating server", err)
	}

	fmt.Println("Listening on port 8888")


	for {

		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err)
			continue
		}
		handleConnection(conn)
	}
}



func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New connection", conn.RemoteAddr())
	_, err := io.WriteString(conn, "+PONG\r\n")
	if err != nil {
		fmt.Println("Error creating data", err)
	}

}
