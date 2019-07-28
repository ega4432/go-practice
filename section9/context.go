package main

import (
	"context"
	"fmt"
	"time"
)

// https://golang.org/pkg/context/
func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	// ctx := context.TODO()
	// エラーとかにしたくない時はとりあえず TODO() にしておくこともある
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	go longProcess(ctx, ch)
	// goroutine がよばれたらすぐにキャンセルになる
	// cancel()

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("#########################")
}
