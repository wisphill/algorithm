package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("pprof at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/handle", leakHandler)
	fmt.Println("serve requests at :8080")
	http.ListenAndServe(":8080", nil)
}

func leakHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		<-ch // waiting for sender forever
	}()

	fmt.Fprintf(w, "goroutines: %d\n", runtime.NumGoroutine())
}
