package _map

import "fmt"

func Map() {
	// Dictionaryと同じ
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	// キーから値を取り出す
	fmt.Println(m["apple"])

	// キーから値を代入
	m["banana"] = 300
	fmt.Println(m)

	// キーの取り出しと取り出し可否の表示
	v, ok := m["apple"]
	fmt.Println(v, ok)      // 100 true
	v2, ok2 := m["nothing"] // 0 false
	fmt.Println(v2, ok2)

	// makeはslice, map, chanなどのオブジェクトを初期化する際に使用する（メソッドの定義より）
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// 下記を実行すると[panic: assignment to entry in nil map]とエラーが表示される
	//var m3 map[string]int
	//m3["pc"] = 5000
	//fmt.Println(m3)

	// varで宣言すると初期値はnil
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
