package main

import (
	"net"
	"os"
	"time"
)


func main() {
	/* Parse CLI arguments */
	args := map[string]string { "port" : "", "pwd" : "" }
	parseArgs(args)
	/* Check if port is correct (it is most basic check that do not cover all possible cases) 
	   also check for another options*/
	if len(args["port"]) < 2 || args["port"][0] != ':' {
		logger(log { time.Now(), true, false, false, "Usage: " + os.Args[0] + " port :<your port num> pwd <password>", nil })
		os.Exit(1)
	}
	cleanup(nil)
	/* If port value is incorrect, user will see an error */
	listener, err := net.Listen("tcp4", args["port"])
	if err != nil {
		logger(log { time.Now(), true, false, false, "could not open listener", err })
		os.Exit(1)
	}
	defer finish(listener, 0)
	go logger(log { time.Now(), false, false, true, "Server started on address " + listener.Addr().String(), nil })
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