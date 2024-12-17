package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
 This is more like a prototipe. This function should wait for input to stdin and react to it.
*/
func listenCLI(listener net.Listener) {
	for {
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("Error:", err)
			continue	
		}
		text := string(data)
		switch(text[:len(text) - 1]) {
		/* Actions performed when command is entered */
		case "quit":
			fmt.Println("Shutting server down...")
			finish(listener, 0)
		default:
			fmt.Println("Unknown command")
		}
	}
}