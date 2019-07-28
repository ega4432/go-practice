package main

import (
	"fmt"
	"github.com/ciruclation-dev/go-practice/section8/package/mylib"
	"github.com/ciruclation-dev/go-practice/section8/package/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
}
