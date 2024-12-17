package main

import (
	"os"
)

/*
 This function takes map, where keys are argumnent names and values are... their values.
 Value should be right after its key. There are no type checks etc. You can initialise 
 values with any string you want ("" is recommended though). 
 */
func parseArgs(args map[string]string) {
	n := len(os.Args)
	for name := range args {
		for i := 1; i < n - 1; i++ {
			if name == os.Args[i] {
				args[name] = os.Args[i + 1]
			}
		}
	}
}