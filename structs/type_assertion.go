package _structs

import "fmt"

// 空のインターフェースを引数とする場合は、なんの型でも引数として受け取れる
func do(i interface{}) {
	// これはキャストでなく、タイプアサーション
	/*
			ii := i.(int)
			ii *= 2
			fmt.Println(ii)
		ss := i.(string)
		fmt.Println(ss)
	*/

	// ここで使用しているtypeは特別なものでswitch文とセット、あらゆるタイプを表す
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know this %t\n", v)
	}
}

// 引数をinterface{}として受け取って、受け取った先でタイプアサーションする
func TypeAssertion() {
	do(10)
	do("Bob")
	do(true)

	// インターフェースではないものを型変換する際は、キャストやタイプコンバージョンというが、
	//インターフェースを型変換する際は、タイプアサーションという
}
