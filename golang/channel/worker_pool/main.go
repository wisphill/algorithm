package main

import (
	"fmt"
	"sync"
)

// spawn more than 10000 jobs
// spawn only 50 goroutines
func main() {
	jobs := make(chan int, 10000)

	var wg sync.WaitGroup
	// spawn 50 workers - 50 goroutines only
	for i := 0; i < 50; i++ {
		go func(wokerId int) {
			fmt.Println("Running worker ", wokerId)

			for job := range jobs {
				processJob(job)
				wg.Done()
			}
		}(i)
	}

	// spawn 1000 jobs
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		jobs <- i
	}

	wg.Wait()
}

func processJob(jobId int) {
	fmt.Println("Processing job ", jobId)
}
