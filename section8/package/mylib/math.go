/*
mylib is yosemite's lib.
*/
package mylib

import "fmt"

// Person2 doc
type Person2 struct {
	Name string
	Age  int
}

func (p *Person2) Say() {
	fmt.Println("Person2")
}

// Average returns the average of a series fo numbers
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
