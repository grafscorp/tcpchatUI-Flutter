package main

import (
	"sync"
	"net"
	"os"
)

/* Messages and mutex for access to them */
var msgs [][]byte
var msg_mutex sync.Mutex

/* Connections and mutex for access */
var conns map[uint32]net.Conn = make(map[uint32]net.Conn)
var conns_mutex sync.Mutex

/* Channel that is used by Sender() (listening) and listenConnection() (sending) */
var sender_chan chan uint32 = make(chan uint32, 10)


/* Finish program with some exit code */ 
func finish(listener net.Listener, code int) {
	cleanup(listener)
	os.Exit(code)
}

/* Reset all global objects (and listener from main()) to default values */
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