package main

import "fmt"

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	// & を付けてポインタを渡してあげる ( Error ではこれが良しとされる書き方)
	return &UserNotFound{Username: "mike"}
}

func main() {
	e1 := UserNotFound{"hoge"}
	e2 := UserNotFound{"fuga"}
	fmt.Println(e1 == e2)

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
