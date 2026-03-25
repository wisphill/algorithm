package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	signal := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// runtime fatal error, cannot catched by recover, because of no receivers on the channel signal
			signal <- 1
		}()
	}

	wg.Wait()
	fmt.Println("Cannot reach here, because of the runtime fatal error. Exit main!!!")
}
