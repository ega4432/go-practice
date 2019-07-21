package main

import (
	"fmt"
	"time"
)

func getOSName() string {
	return "hogefuga"
}

func main() {
	switch os := getOSName(); os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Win!")
	default:
		fmt.Println("other...", os)
	}

	t := time.Now().Hour()
	fmt.Println(t)

	switch {
	case t > 3 && t < 12:
		fmt.Println("Morning")
	case t > 12 && t < 18:
		fmt.Println("Afternoon")
	case t > 18 && t < 21:
		fmt.Println("NIGHT...")
	case t > 21:
		fmt.Println("Midnight")
	default:
		fmt.Println("Time is Unknown")
	}
}
