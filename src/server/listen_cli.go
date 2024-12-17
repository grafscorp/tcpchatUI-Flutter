package main

import (
	"bufio"
	"net"
	"os"
	"time"
)

/*
 This is more like a prototipe. This function should wait for input to stdin and react to it.
*/
func listenCLI(listener net.Listener) {
	for {
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			logger(log { time.Now(), false, true, false, "Can't read CLI input", err })
			continue	
		}
		text := string(data)
		switch(text[:len(text) - 1]) {
		/* Actions performed when command is entered */
		case "quit":
			logger(log { time.Now(), false, false, true, "Shutting server down...", nil })
			finish(listener, 0)
		case "reset":
			logger(log { time.Now(), false, false, true, "Reset...", nil })
			cleanup(nil)
		default:
			logger(log { time.Now(), false, true, false, "Unknown command", nil })
		}
	}
}