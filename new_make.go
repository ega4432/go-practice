package main

import "fmt"

func main() {
	// new
	var p *int = new(int)
	fmt.Println(p)   // 0xc000090000
	fmt.Println(*p)  // 0
	fmt.Println(&p)  // 0xc00000e020

	*p++
	fmt.Println(*p)

	var q *int
	fmt.Println(q)  // nil
	//*q++
	// fmt.Println(*q)
	// panic: runtime error: invalid memory address or nil pointer dereference

	// make
	s := make([]int, 0)
	fmt.Printf("%T\n", s)  // []int

	m := make(map[string]int)
	fmt.Printf("%T\n", m)  // map[string]int

	ch := make(chan int)
	fmt.Printf("%T\n", ch)  // chan int

	var st = new(struct{})
	fmt.Printf("%T\n", st)  // *struct {}
}
