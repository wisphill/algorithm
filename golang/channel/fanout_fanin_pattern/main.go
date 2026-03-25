package main

import (
	"context"
	"fmt"
	"sync"
)

// readonly jobs <-chan int
// writeonly results chan<- int
func executeWorker(ctx context.Context, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// fan-in to one channel
		results <- job * job
	}
}

// numbers = [1,2,3,4,5,6]
// workers = 3
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	numberOfWorkers := 3

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// fanout to multiple workers
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		// worker
		go executeWorker(context.TODO(), jobs, results, &wg)
	}

	// distribute the jobs to workers to process
	go func() {
		for number := range numbers {
			jobs <- number
		}

		close(jobs) // there is no data comming to the jobs channel, so the for range can be closed.
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	// for-range only be done when the channel is closed
	for result := range results {
		fmt.Printf("Fanning out to print the value of result: %v \n", result)
	}
}
