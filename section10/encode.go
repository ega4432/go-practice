package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	Name      string   `json:"name"` // `json:"-"` とすることで表示させなくすることもできる
	Age       int      `json:"age"`  // `json:"age,string"` とすることで型を string にできる
	Nicknames []string `json:"nicknames"`
	// `json:"name,omitempty"` のように後ろに omitempty をつけると空文字の際に表示しない
	// string だと空文字 int だと 0 slice 空slice

	// ポインタを渡さないと omitempty をつけていても表示する
	T *T `json:"T,omitempty"`
}

// Marshal が呼ばれた時にカスタマイズしたい時に追加で処理をしたい時に使う
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

// Unmarshal が呼ばれた時にカスタマイズしたい時に追加で処理をしたい時に使う
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func main() {
	// Unmarshal
	b := []byte(`{"name": "mike", "age":20, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// Marshal
	// どのような名前かを struct のところで指定できる
	v, _ := json.Marshal(p)
	fmt.Println((string(v)))

}
