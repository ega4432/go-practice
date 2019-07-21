package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(20, 10)
	fmt.Println(r1, r2)

	r3 := cal(1000, 3)
	fmt.Println(r3)

	r4 := convert(234)
	fmt.Println(r4)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(2)
}
