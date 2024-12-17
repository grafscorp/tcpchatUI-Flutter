package main

/*
 This function uses channel named sender chan. It receives message id from it
 (this id is just position in msgs slice) and writes this message to every 
 connection in conns slice (including author, yes, this decision was made 
 intentionally - sender should know his message is delivered correctly).
 */
func sender() {
	for {
		id := <-sender_chan // Waiting
		/* Getting data */
		msg_mutex.Lock()
		data := msgs[id]
		msg_mutex.Unlock()
		/* Sending data */
		conns_mutex.Lock()
		for _, c := range conns {
			c.Write(data) // errors here shoud be ignored
		}
		conns_mutex.Unlock()
	}
}