package main

import (
	"fmt"
	"time"
)

type log struct {
	time  time.Time
	is_error   bool
	is_warning bool
	is_info    bool
	text     string
	err       error
}

func logger(l log) {
	text := "[" + l.time.Format(time.UnixDate) + "]"
	if l.is_error {
		text += "[ERROR]"
	}
	if l.is_warning {
		text += "[WARNING]"
	}
	if l.is_info {
		text += "[INFO]"
	}
	if l.err != nil {
		text += " (" + l.err.Error() + ") "
	}
	text += ": " + l.text
	/* TODO: logging into file */
	console_mutex.Lock()
	fmt.Println(text)
	console_mutex.Unlock()
}