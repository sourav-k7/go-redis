package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Connected new client:", conn.RemoteAddr())
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by client:", conn.RemoteAddr())
			break
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Received from client: %s\n", message)

		if message == "exit" {
			fmt.Println("Closing connection for client:", conn.RemoteAddr())
			break
		}

		response := "Server received: " + message + "\n"
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error sending response:", err)
			break
		}
	}

}
