package _pointer

import (
	"fmt"
)

func NewAndMake() {
	// 値を何も入れないでメモリにポインタが入る領域を確保したい場合はnewを使う
	// 下記を出力するとアドレスが返ってくるが、値は何も入っていない
	var p *int = new(int)
	// pの中身を見てみると0が入ってる！！
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// こちらはポインタ型で宣言はしているものの、メモリは確保されていないのでnilが出力される
	//var p2 *int
	//fmt.Println(p2)

	// newで宣言されたものと、makeで宣言されたものの型を見てみる
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p2 *int = new(int)
	fmt.Printf("%T\n", p2)

	// newとmakeはポインタを返すか返さないかの違い（今のところは）
}
