package main

import (
	"fmt"
	"time"
)

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
