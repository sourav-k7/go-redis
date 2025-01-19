package server

import (
	"fmt"
	"net"
)

func ServerSetup() {
	listener, err := net.Listen("tcp", ":6379")
	fmt.Println("Listing on port 6379")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}