package main

import "fmt"

func main()  {
	// バッファの数を 2 としておく
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	// チャネルを一旦クローズしてあげる必要がある
	close(ch)

	// 取り出す
	for c := range ch {
		fmt.Println(c)
	}
}