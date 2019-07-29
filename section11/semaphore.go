package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// DOC( https://godoc.org/golang.org/x/sync/semaphore )

// 同時に走る goroutine の数を制限する
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// 他の goroutine をキャン節にしたい場合
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}

	// 他の goroutine を待たせたい場合
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	// Release されたあとだと実行される
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
