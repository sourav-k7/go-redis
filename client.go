package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send command (with newline so the server knows the command ended)
	fmt.Fprintf(conn, "SET TEST VALUE\n")

	// Read server response (optional, but helps debugging)
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Server response: ", response)

	// Keep console open if running from double-click
	fmt.Println("Press Enter to exit")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
