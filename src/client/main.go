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
		fmt.Println(time.Now().Format(time.Kitchen) + ": " + string(buf))
	}
}

func main() {
	args := map[string]string{ "name" : "", "pwd" : "", "address" : "" }
	parseArgs(args)
	if args["address"] == "" {
		fmt.Println("Usage: " + os.Args[0] + " address <server address> pwd <password>")
		os.Exit(1)
	}
	conn, err := net.Dial("tcp4", args["address"])
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
		line := string(data)
		/* Empty string */
		if (len(line) < 2) {
			continue
		}
		line = line[:len(line) - 1]
		if (line == "!stop") {
			fmt.Println("Closing client...")
			break
		}
		/* We don't need newline here */
		data = ([]byte(args["name"] + ": " + line))
		if len(data) >= 1024 {
			data = data[:1024]
		} // Yeah, text could be truncated
		
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error sending data:", err)
			break	
		}
	}
}