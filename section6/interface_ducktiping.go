package main

import "fmt"

// インターフェース
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	// 下記のように p を書き換えたい場合は、レシーバでポインタを渡す
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get Out!")
	}
}

func main() {
	// Person の中で使いたいときは、アドレス & を渡す
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	//var dog Dog = Dog{"dog"}
	DriveCar(mike)
	DriveCar(x)
	// Dog は　Say()が実装されていないのでえらーになる
	// DriveCar(dog)
}
