package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("beak")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	fmt.Println(sum)
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// infinite
	for {
		fmt.Println("infinite!")
	}
}
