package main

import "fmt"

type Person struct {
	Name string
	Age  int
}


// String メソッドでは、 p にある一部しか表示させたくない時などに利用できる
func (p Person) String() string {
	return fmt.Sprintf("My name is '%v", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
