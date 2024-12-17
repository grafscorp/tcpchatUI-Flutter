package main

import (
	"net"
	"fmt"
)

/*
 This function manages connection life cycle. Here it is added to and removed from conns slice,
 here connaction will be closed. Here listenConnection() function is being called for conn connection.
 */
func handleConnection(conn net.Conn) {
	/* It is not really needed since this function will get to this call
	   obly when connection will be closed (or some another error will occurr) */
	defer conn.Close()
	var conn_id uint32 = 0 // key in cons map
	/* Find minimal available key, add connection to map */
	conns_mutex.Lock()
	for ; ; conn_id++ {
		if _, ok := conns[conn_id]; !ok {
			conns[conn_id] = conn
			break
		}
	}
	conns_mutex.Unlock()
	/* Wait for listenConnection() to end */
	var done chan bool = make(chan bool)
	go listenConnection(conn, done)
	res := <-done
	if res {
		fmt.Println("Connection closed:", conn)
	} else {
		fmt.Println("Error in connection:", conn)
	}
	/* Delete connection from list */
	conns_mutex.Lock()
	delete(conns, conn_id)
	conns_mutex.Unlock()
}