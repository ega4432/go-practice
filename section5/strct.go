package main

import "fmt"

type Vertex struct {
	// 大文字にしておく
	X, Y int
	S    string
}

func hoge() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v) // {1 2 }

	fmt.Println(v.X, v.Y) // 1 2

	v.X = 1000
	v2 := Vertex{X: 1}
	fmt.Println(v2) //{1 0 }

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3) // {1 2 test}

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4) // main.Vertex {0 0  }

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5) // main.Vertex {0 0  }

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6) // main.Vertex &{0 0  }

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7) // *main.Vertex &{0 0 }
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	hoge()

	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v) // {1 2 test}

	vv := &Vertex{1, 2, "testtest"}
	changeVertex2(vv)
	fmt.Println(vv)  // &{1000 2 test}
	fmt.Println(*vv) // {1000 2 test}
}
