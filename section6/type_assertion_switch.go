package main

import "fmt"

func do(i interface{}) {
	// int が渡ってきた時
	//ii := i.(int)
	//ii *= 2
	//fmt.Println(ii)

	// string が渡ってきた時
	//s := i.(string)
	//fmt.Println(s + "!")

	// どの型が入ったかによってスイッチする
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + " hoge")
	default:
		fmt.Printf("I do not know %T\n", v)
	}
}

func main() {
	// この時点ではただの interface で int ではない
	//var i interface{} = 10
	//do(i)

	// 次に文字列を私見る
	//var i interface{} = "Mike"
	//do(i)

	// なんでも入れてみる
	do(10)
	do("Mike")
	do(true)

	// type を明示的に int に指定している
	var i int = 10
	// キャスト、タイプコンバージョンしている
	ii := float64(10)
	fmt.Println(i, ii)
}
