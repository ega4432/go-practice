package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

// クライアントから送られてきたハッシュが正しいかを判定する
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC) // true
}

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	// hmac でよく使われる書き方 -> ( https://golang.org/pkg/crypto/hmac/ )
	h := hmac.New(sha256.New, []byte(apiSecret)) // byte の配列が生成される
	h.Write(data)                                // サーバに送りたいものを書く
	sign := hex.EncodeToString(h.Sum(nil))       // エンコードする

	fmt.Println(sign)

	Server(apiKey, sign, data)
}
