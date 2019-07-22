package main

import "fmt"

// Q1. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示字されますか？
func main() {
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j) // A. 10

	hoge()
}

//Q2. 以下の実行結果をコードを書かずに答えてください。エラーがおきますか？正しく表示されるとすると何が表示されますか？
func hoge() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2
	// i が変わればポインタの p1 も変わる
	p2 = p1
	j = *p2 + i
	fmt.Println(j) // A. 600
}
