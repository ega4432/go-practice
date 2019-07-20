package main

import "fmt"

// comment

/*
this is comment for long area
*/

func init() {
	fmt.Println("this it init!")
}

func bazz() {
	fmt.Println("Bazz")
}
func main() {
	bazz()
	fmt.Println("Hello world!")
}
