package mylib

import (
	"fmt"
	"testing"
)

var Dubug bool = true

func Example() {
	v := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func ExampleAverage() {
	v := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func ExamplePerson2_Say() {
	p := Person2{"Mike", 20}
	fmt.Println(p)
}

func TestAverage(t *testing.T) {
	if Dubug {
		t.Skip("this is skip reason...")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
	// テストの実行は "$ go test ./..." or Golan　からできる
	// testing フレームワーク Ginkgo を使っても良い( https://onsi.github.io/ginkgo/ )

	// gofmt <file name> で綺麗なフォーマットを表示してくれる
	// 勝手に整形までして欲しい時はオプションで -w をつける
}
