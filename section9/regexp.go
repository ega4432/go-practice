package main

// https://golang.org/pkg/regexp/
import (
	"fmt"
	"regexp"
)

func main() {
	// 最初が"a"、最後が"e"に挟まれた a から z の文字列にヒットするかどうか
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match) // return true

	r := regexp.MustCompile("a([a-z])+e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.MatchString("/view/test")
	fmt.Println(fs)

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}
