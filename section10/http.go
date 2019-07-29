package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// http
	resp, _ := http.Get("https://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// url
	base, _ := url.Parse("https://example.com")
	fmt.Println(base)

	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	// GET なので nil で良いが、 POST だと byte を渡す => ex: req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	// "req" には、リクエストを送る構造体が入っている
	// &{GET https://example.com/test?a=1&b=2 HTTP/1.1 1 1 map[] <nil> <nil> 0 [] false example.com map[] map[] <nil> map[]   <nil> <nil> <nil> <nil>}

	// リクエストに適当な文字列をヘッダに追加する
	req.Header.Add("If-None-Match", `W/wyxxy`)
	q := req.URL.Query()
	fmt.Println(q)

	// リクエストに適当な文字列を追加してみてどうエンコードされるか見てみる
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	// 実際に作ったリクエストで client にアクセスする
	var client *http.Client = &http.Client{}
	resp2, _ := client.Do(req)
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(string(body2))
}
