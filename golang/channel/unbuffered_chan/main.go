package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	signal := make(chan int)

	go func() {
		for {
			select {
			case <-signal:
				fmt.Println("catch a signal")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// fatal error, cannot catched by recover
			signal <- 1
		}()
	}

	wg.Wait()
	fmt.Println("Exit main")
}
