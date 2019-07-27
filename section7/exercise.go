package main

import "fmt"

//  Q1. 以下のように文字列を連結させて出力したいコードがありますが、
//  test1!
//	test1!test2!
//	test1!test2!test3!
//	test1!test2!test3!test4!
//	以下のコードには間違いがあります。上記の出力になるように修正してください。

// false : func goroutine(s []string, c chan int) {
func goroutine(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	// add this
	close(c)
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	// false : c := make(chan int)
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
