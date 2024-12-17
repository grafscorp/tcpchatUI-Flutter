package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var msgs [][]byte
var msg_mutex sync.Mutex

var conns map[uint32]net.Conn = make(map[uint32]net.Conn)
var conns_mutex sync.Mutex

var sender_chan chan uint32 = make(chan uint32, 10)

func listenCLI() {
	for {
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("Error:", err)
			continue	
		}
		text := string(data)
		switch(text[:len(text) - 1]) {
		case "quit":
			fmt.Println("Shutting server down...")
			os.Exit(1)
		}
	}
}

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