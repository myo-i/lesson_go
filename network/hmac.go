package _network

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

// Server クライアントから送られてきたキーとDBのキーが一致するか判定する
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

func Hmac() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	// ここから
	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	// ハッシュを生成してそれをサーバーサイドに投げることでユーザーが正しいものかを判定したりする
	sign := hex.EncodeToString(h.Sum(nil))
	// ここまでのステップはほぼ同じなので頭に入れておく

	fmt.Println(sign)

	Server(apiKey, sign, data)
}
