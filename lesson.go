package main

import "fmt"

func thirdPartyConnnectDB()  {
	panic("Unable to connect database!")
}

func save()  {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnnectDB()
}
func main() {
	save()
	fmt.Println("OK?")
}
