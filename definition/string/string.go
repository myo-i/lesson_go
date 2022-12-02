package string

import (
	"fmt"
	"strings"
)

func String() {
	fmt.Println("Hello World"[0])         // ASCII 72
	fmt.Println(string("Hello World"[0])) // string H

	// 文字列の置換
	s := "Hello How Are You?"

	// strings.Replaceは本体の値を変えるわけではないので、保持するには変数に代入する必要がある
	s = strings.Replace(s, "H", "X", 2)
	fmt.Println(s)

	// テンプレートリテラル？
	fmt.Println(`Test
         Test
Test`)

	// 文字列の中にダブルクォーテーション
	fmt.Println("\"Hello\"")
	fmt.Println(`"Good"`)
}
