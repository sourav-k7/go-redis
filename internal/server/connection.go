package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/sourav-k7/go-redis/internal/cmd"
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
		} else {
			cmdRes, err := cmd.Execute(message)

			if err != nil {
				// Send error message + CRLF
				resp := err.Error() + "\r\n"
				if _, e := conn.Write([]byte(resp)); e != nil {
					fmt.Println("Error sending error response:", e)
				}
				continue
			}

			// Handle successful response
			switch v := cmdRes.(type) {

			case string:
				// If it's a string, just send it + CRLF
				resp := v + "\r\n"
				if _, e := conn.Write([]byte(resp)); e != nil {
					fmt.Println("Error sending response to", conn.RemoteAddr(), ":", e)
				}

			default:
				// Any other type → fallback string
				resp := "Not handled response type\r\n"
				if _, e := conn.Write([]byte(resp)); e != nil {
					fmt.Println("Error sending fallback response:", e)
				}
			}

		}
	}
}
