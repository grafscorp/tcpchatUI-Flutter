package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp4", ":50051")
	if err != nil {
		fmt.Println("Error: could not open listener\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println(time.Now().Format(time.Kitchen) + " Server started")
	go sender()
	go listenCLI()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: could not accept connection")
			continue
		}
		go handleConnection(conn)
	}
}