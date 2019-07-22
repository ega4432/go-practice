package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n) // 1

	fmt.Println(&n) // 0xc00008a000
	fmt.Println(*&n)

	var p *int = &n
	fmt.Println(p) // 0xc00008a000

	fmt.Println(*p) // 1
}
