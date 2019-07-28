package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// "lesson.go" の中身を読み込む
	content, err := ioutil.ReadFile("lesson.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// "ioutil_tmp.go" というファイルを作成して、書き込む
	if err := ioutil.WriteFile("ioutil_tmp.go", content, 0666); err != nil {
		log.Fatal(err)
	}

	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2))
}
