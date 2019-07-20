package main

import "fmt"

var i int = 1
var f64 float64 = 1.2
var s string = "test"
var t, f bool = true, false

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
