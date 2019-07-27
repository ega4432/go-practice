package main

import (
	"fmt"
	"sync"
	"time"
)

/*
** イメージ：
** goroutine で並列で走らせた処理を
** Producer で取得した値を channel を使って Consumer に渡す
 */

func producer(ch chan int, i int) {
	// 何かやる
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {

		func() {
			defer wg.Done()
			fmt.Println("proccess", i*1000)
		}()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
