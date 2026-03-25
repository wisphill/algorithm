package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// avoiding back pressure
type RateLimitter struct {
	tokens chan struct{}
}

func NewRateLimitter(ctx context.Context, ratePerSec int, burst int) *RateLimitter {
	rl := RateLimitter{
		tokens: make(chan struct{}, burst),
	}

	for _ = range burst {
		rl.tokens <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(ratePerSec))
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				rl.tokens <- struct{}{}
			case <-ctx.Done():
				fmt.Println("context has been exceeded")
				return
			default:
				// skip to add to the chan
			}
		}
	}()

	return &rl
}

func (rl *RateLimitter) Acquire(ctx context.Context) error {
	select {
	case <-rl.tokens:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("context is exceeded")
	}
}

func main() {
	var wg sync.WaitGroup
	workerRun := func(i int) {
		time.Sleep(time.Second * 5)
		wg.Done()
	}

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go workerRun(i)
	}

	wg.Wait()
	fmt.Println("All workers (without rate limiter) are completed")

	ctx := context.Background()
	rl := NewRateLimitter(ctx, 5, 20)
	var wg2 sync.WaitGroup
	workerRun2 := func(i int) {
		defer wg2.Done()
		err := rl.Acquire(ctx)
		if err != nil {
			fmt.Println("Error while running worker ", i)
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println("Proceeded worker successfully ", i)
	}

	for i := 0; i < 200; i++ {
		wg2.Add(1)
		go workerRun2(i)
	}

	wg2.Wait()
}
