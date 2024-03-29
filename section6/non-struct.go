package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	// i は type : main.myInt を返す
	fmt.Printf("%T %v\n", i, i) // main.MyInt 10
	fmt.Printf("%T %v\n", 1, 1) // int 1
	// int にキャストしている（しないとエラー）
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
