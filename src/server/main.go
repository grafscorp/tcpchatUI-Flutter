package main

import (
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

func sender() {
	for {
		id := <-sender_chan
		msg_mutex.Lock()
		data := msgs[id]
		msg_mutex.Unlock()
		conns_mutex.Lock()
		for _, c := range conns {
			c.Write(data) // errors here shoud be ignored
		}
		conns_mutex.Unlock()
	}
}

func listenConnection(conn net.Conn, done chan bool) {
	for {
		var data []byte = make([]byte, 1024)
		_, err := conn.Read(data)
		if err != nil {
			done<-true
			return
		}
		msg_mutex.Lock()
		msgs = append(msgs, data)
		n := uint32(len(msgs) - 1)
		fmt.Println(time.Now().Format(time.Kitchen) + " Message received\n>" + string(data))
		msg_mutex.Unlock()
		sender_chan<-n
	}
}

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

func main() {
	listener, err := net.Listen("tcp4", ":50051")
	if err != nil {
		fmt.Println("Error: could not open listener\n", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println(time.Now().Format(time.Kitchen) + " Server started")
	go sender()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: could not accept connection")
			continue
		}
		go handleConnection(conn)
	}
}