package main

import "fmt"

// Pipeline processing: generator → stage1 → stage2 → stage3 → output, multiple stages
// and communicate via channels
// Sample: generate numbers → square → print
func main() {
	nums := generateNumber(1, 2, 3, 4, 5, 6, 7, 8, 9)
	squares := calculateSquare(nums)
	for v := range squares {
		fmt.Printf("%v ", v)
	}
}

func generateNumber(numbers ...int) <-chan int {
	numberChan := make(chan int)
	// handle in the goroutines
	go func() {
		for number := range numbers {
			numberChan <- number
		}
		close(numberChan)
	}()
	return numberChan
}

// only pipeline the data, not fanout to workers here
func calculateSquare(numberChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for number := range numberChan {
			out <- number * number
		}

		close(out)
	}()

	return out
}
