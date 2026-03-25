package main

import (
	"fmt"
	"time"
)

// all kinds of data like array, slices, struct of the Go are not thread-safe because they does not have the lock inside
// the channel type is thread safe, because
func main() {

	// runtime fatal error because of concurrent read-write at the same time
	m := map[int]int{}

	// writer
	go func() {
		for {
			m[1] = 1
		}
	}()

	// reader
	go func() {
		for {
			fmt.Println(m[1])
		}
	}()

	time.Sleep(5 * time.Second)
}
