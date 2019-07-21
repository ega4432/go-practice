package main

import (
	"fmt"
	"os"
)

func foo()  {
	defer fmt.Println("world foo")
	fmt.Println("Hello foo")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")
	foo()

	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 1000)
	file.Read(data)
	fmt.Println(string(data))
}
