package main

import "fmt"

// channel を使うと goroutine でデータのやりとりができる
func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine(s, c)

	// goroutine で受け取ったチャネルc次々に取り出していく
	for i := range c {
		fmt.Println(i)
	}
}
