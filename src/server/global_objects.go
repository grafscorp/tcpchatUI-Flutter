package main

import (
	"sync"
	"net"
)

/* Messages and mutex for access to them */
var msgs [][]byte
var msg_mutex sync.Mutex

/* Connections and mutex for access */
var conns map[uint32]net.Conn = make(map[uint32]net.Conn)
var conns_mutex sync.Mutex

/* Channel that is used by Sender() (listening) and listenConnection() (sending) */
var sender_chan chan uint32 = make(chan uint32, 10)