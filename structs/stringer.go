package _structs

import "fmt"

type AMan struct {
	Name string
	Age  int
}

// String()は特別なメソッド（定義はfmt/print.goにされている）
// ドキュメント訳（fmt/print.go）
// 【Stringer は、String メソッドを持つすべての値で実装されています。その値の「ネイティブ」フォーマットを定義します。】
// 大事なことだから二度・・・Stringerはネイティブフォーマットの定義
func (m AMan) String() string {
	return fmt.Sprintf("My name is %v\n", m.Name)
}

// 普通に出力すると年齢まで出力されてしまうが、Stringerを用いれば出力の結果を変えることができる
func Stringer() {
	bob := AMan{"Bob", 23}
	fmt.Println(bob)
}
