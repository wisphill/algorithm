package main

import (
	"context"
	"fmt"
	"time"
)

// 🟡 Bài 2 — Fan-out / Fan-in + Early Cancel (mid → senior)
// Đề
// Search từ 5 data source khác nhau (API, DB, cache...).
// Yêu cầu:
// Chạy song song
// Trả về kết quả đầu tiên tìm thấy
// Cancel các goroutine còn lại
// Không leak
type Result struct {
	data   string
	source string
}

func simulateApiCall(ctx context.Context, source string, delayInSecs int) *Result {
	select {
	case <-ctx.Done():
		fmt.Println("context has been cancelled for ", source)
		return nil
	case <-time.After(time.Second * time.Duration(delayInSecs)):
		return &Result{
			data:   "response data",
			source: source,
		}
	}
}

// 👉 Đây là pattern “First response wins”.
// Early cancel using one DONE signal + context cancel
func main() {
	type Source struct {
		source string
		delay  int
	}

	apiSourceCalls := []*Source{
		{
			source: "api",
			delay:  1,
		},
		{
			source: "external",
			delay:  2,
		},
		{
			source: "db",
			delay:  3,
		},
	}
	doneSignal := make(chan *Result)
	ctx, cancel := context.WithCancel(context.Background())
	for _, source := range apiSourceCalls {
		go func(source *Source) {
			res := simulateApiCall(ctx, source.source, source.delay)
			doneSignal <- res
		}(source)
	}

	anyRes := <-doneSignal
	cancel()
	fmt.Println("Get the quickest response: ", anyRes)

	time.Sleep(5 * time.Second)
}
