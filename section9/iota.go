package main

import "fmt"

// const の連番を振ることができる
const (
	c1 = iota
	c2
	c3
)

const (
	_ = iota
	// "<<" は、ビット演算子。2進数で表した時に数値をずらしたりする
	KB int = 1 << (10 * iota) // 1 に 0 を 10 個つけると、"10000000000"。これを10進数で表すと 1024
	MB
	GB
)

func main() {
	fmt.Println(c1, c2, c3) // 0 1 2
	fmt.Println(KB, MB, GB) // 1024 1048576 1073741824
}
