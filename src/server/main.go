package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func finish(listener net.Listener, code int) {
	cleanup(listener)
	os.Exit(code)
}

func cleanup(listener net.Listener) {
	msgs = make([][]byte, 0)
	for _, val := range conns {
		val.Close()
	}
	conns = make(map[uint32]net.Conn, 0)
	sender_chan = make(chan uint32, 10)
	if listener != nil {
		listener.Close()
	}
}

func main() {
	cleanup(nil)
	listener, err := net.Listen("tcp4", ":50051")
	if err != nil {
		fmt.Println("Error: could not open listener\n", err)
		os.Exit(1)
	}
	defer finish(listener, 0)
	fmt.Println(time.Now().Format(time.Kitchen) + " Server started")
	go sender()
	go listenCLI(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: could not accept connection (most likely server is shutting down)")
			time.Sleep(time.Second) // It's probably because cleanup() closed connection
			continue                // So... we'll just wait
		}
		go handleConnection(conn)
	}
}