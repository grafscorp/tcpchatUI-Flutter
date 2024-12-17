package main

import (
	"fmt"
	"time"
	"net"
)

/*
 This function waits for any data to arrive on one connection (maximum is 1024 bytes at once), then
 this data is logged and sent to msgs slice. After that sender routine is notified via sender_chan
 channel
 */
func listenConnection(conn net.Conn, done chan bool) {
	for {
		var data []byte = make([]byte, 1024)
		/* Wait for incoming data */
		_, err := conn.Read(data)
		if err != nil {
			done<-true
			return
		}
		/* Write data to msgs */
		msg_mutex.Lock()
		msgs = append(msgs, data)
		msgid := uint32(len(msgs) - 1) // Message id
		fmt.Println(time.Now().Format(time.Kitchen) + " > " + string(data))
		msg_mutex.Unlock()
		sender_chan<-msgid // Notify sender
	}
}