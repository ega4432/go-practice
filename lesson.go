package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])

	m["banana"] = 300
	fmt.Println(m)

	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, hoge := m["apple"]
	fmt.Println(v, hoge)

	v2, hoge2 := m["nothing"]
	fmt.Println(v2, hoge2)

	m2 := make(map[string]int)
	m2["pc"] = 500000
	fmt.Println(m2)

	//var m3 map[string]int
	//m3["pc"] = 500
	//fmt.Println(m3)

	var s []int
	if s == nil {
		fmt.Println("This is nil!!!!")
	}
}
