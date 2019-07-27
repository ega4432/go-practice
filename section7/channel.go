package main

import "fmt"

// channel を使うと goroutine でデータのやりとりができる
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(ss []int, c chan int) {
	sum := 0
	for _, v := range ss {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	ss := []int{2, 4, 6, 8, 10}
	c := make(chan int) // 15, 15
	go goroutine1(s, c)
	go goroutine2(ss, c)
	// goroutine1 で受け取ったチャネルを x に入れる
	x := <-c
	fmt.Println(x)
	// goroutine2 で受け取ったチャネルを y に入れる
	y := <-c
	fmt.Println(y)
}
