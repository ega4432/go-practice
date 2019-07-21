package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("ERROR1")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("ERROR2")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("ERROR3")
	}
}
