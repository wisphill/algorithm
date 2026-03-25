package main

import (
	"fmt"
	"time"
)

// main thread and goroutine
func happen_before_1() {
	ch := make(chan struct{})
	var x int

	go func() {
		x = 42
		ch <- struct{}{} // sender
	}()

	<-ch // receiver
	// all assignments happen before the corresponded receiving
	// fmt.Println(x) is in the same channel with receiver > print x = 42
	fmt.Println(x)
}

// different goroutines
func happen_before_2() {
	ch := make(chan struct{})
	var x int

	go func() {
		x = 45
		ch <- struct{}{} // sender
	}()

	go func() {
		<-ch // receiver
		// all assignments happen before the corresponded receiving
		// fmt.Println(x) is in the same channel with receiver > print x = 42
		fmt.Println(x)
	}()
}

func main() {
	happen_before_1()
	happen_before_2()
	time.Sleep(time.Second)
}
