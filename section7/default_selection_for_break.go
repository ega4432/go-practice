package main

import (
	"fmt"
	"time"
)

// A Tour of Go を参考にしている（ https://tour.golang.org/concurrency/6 ）

func main() {
	tick := time.Tick(100 * time.Millisecond) // Tick, After は channel を返す
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			// case t := <-tick: とすることで channel で何が渡ってきているのか中身を知ることもできる
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			// 完全に main 関数から抜ける
			return
			// break はループから抜けるだけ
			// Golang には、OuterLoop というものがあり、ループの前に書いてあげることでループから抜けることもできる
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
