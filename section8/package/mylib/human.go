package mylib

import "fmt"

var Pubilc string = "Pubilc"  // 外部から呼び出し可能
var public string = "private" // 外部から呼び出し不可

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
