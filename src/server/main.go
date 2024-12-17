package main

import (
	"net"
	"os"
	"time"
)


func main() {
	cleanup(nil)
	listener, err := net.Listen("tcp4", ":50051")
	if err != nil {
		logger(log { time.Now(), true, false, false, "could not open listener", err })
		os.Exit(1)
	}
	defer finish(listener, 0)
	go logger(log { time.Now(), false, false, true, "Server started", nil })
	go sender()
	go listenCLI(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			go logger(log { time.Now(), false, true, false, "could not accept connection (most likely server is shutting down)", err })
			time.Sleep(time.Second) // It's probably because cleanup() closed connection
			continue                // So... we'll just wait
		}
		go handleConnection(conn)
	}
}