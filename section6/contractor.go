package main

import "fmt"

type Vertex struct {
	x, y int
}

// こっちがメソッド（ドットで呼び出せる）
func (v Vertex) Area() int {
	return v.x * v.y
}

// "*" で書き換えができる
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// ただの関数
func Area(v Vertex) int {
	return v.x * v.y
}

// コンストラクタ
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	//v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
