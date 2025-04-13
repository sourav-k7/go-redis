package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func ServerSetup() {
	listener, err := net.Listen("tcp", ":6379")
	fmt.Println("Listing on port 6379")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			go handleConnection(conn)
		}
	}()
	<-stop
	fmt.Println("\nShutting down server...")

	// Close the listener before exiting
	listener.Close()
	fmt.Println("Server closed.")
}
