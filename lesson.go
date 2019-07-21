package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "Odd"
	} else {
		return "Even"
	}
}
func main() {
	num := 7
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else!")
	}

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("$$")
	} else if x == 10 || y == 10 {
		fmt.Println("||")
	} else {
		fmt.Println("false")
	}

	result := by2(5)
	if result == "Odd" {
		fmt.Println("This num is odd!")
	} else if result == "Even" {
		fmt.Println("This num is Even!")
	} else {
		fmt.Println("ERROR!")
	}
}
