package main

import "fmt"

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
