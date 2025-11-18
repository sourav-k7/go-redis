package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}

		line = strings.TrimSpace(line)

		// If user types exit → send "exit" to server first
		if strings.EqualFold(line, "exit") {
			// Send exit command to server
			fmt.Fprintf(conn, "exit\n")

			// Optional: try reading server's exit response
			resp, _ := serverReader.ReadString('\n')
			fmt.Println(resp)

			fmt.Println("Closing client.")
			return
		}

		// Normal command
		_, err = fmt.Fprintf(conn, line+"\n")
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Server closed connection:", err)
			return
		}
		response = strings.TrimRight(response, "\r\n")
		fmt.Println(response)
	}
}
