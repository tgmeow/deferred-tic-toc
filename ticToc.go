// Package deferredTicToc provides a function to run a deferred timer
package deferredTicToc

import (
	"log"
	"time"
)

// A timer that starts on execution and returns a function that prints the elapsed time.
// Example use: defer ticToc("function1")()
func TicToc(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		log.Printf("%v took %s\n", name, elapsed)
	}
}
