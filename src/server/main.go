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

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var conn_id uint32 = 0

	conns_mutex.Lock()
	for ; ; conn_id++ {
		if _, ok := conns[conn_id]; !ok {
			conns[conn_id] = conn
			break
		}
	}
	conns_mutex.Unlock()

	var done chan bool = make(chan bool)
	go listenConnection(conn, done)
	res := <-done
	if res {
		fmt.Println("Connection closed:", conn)
	} else {
		fmt.Println("Error in connection:", conn)
	}

	conns_mutex.Lock()
	delete(conns, conn_id)
	conns_mutex.Unlock()
}

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