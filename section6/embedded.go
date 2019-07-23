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

//// ただの関数
//func Area(v Vertex) int {
//	return v.x * v.y
//}

type Vertex3D struct {
	Vertex
	z int
}

// こっちがメソッド（ドットで呼び出せる）
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

// "*" で書き換えができる
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// コンストラクタ
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	//v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}
