package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func listen(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error Receiving data:", err)
			os.Exit(1)	
		}
		fmt.Println(time.Now().Format(time.Kitchen) + ":" + string(buf))
	}
}

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:50051")
	if err != nil {
		fmt.Println("Error: could not connect to server")
		os.Exit(1)
	}	
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	go listen(conn)
	for {
		data, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("Error:", err)
			break	
		}
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error sending data:", err)
			break	
		}
	}
}