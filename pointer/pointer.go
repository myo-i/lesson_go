package _pointer

import "fmt"

func one(x *int) {
	*x = 1
}

func Pointer() {
	var n int = 100
	fmt.Println(n)  // 100
	fmt.Println(&n) // 0xc00001c098
	// 上記解説
	// n int = 100とは、変数nのメモリにint型の100を格納するということ
	// &nでメモリに格納されている値のアドレスを見に行っている

	// integerのポインタ型
	var p *int = &n
	fmt.Println(p)  // 0xc00001c098
	fmt.Println(*p) // 100
	// 上記解説
	// p *int = &nとは、変数pのメモリにnのアドレスを格納するということ
	// *pでメモリに格納されているアドレスの中身を見に行っている

	// ただのint型の変数に数値を代入しようとしても変化はないが、ポインタ型の中身を変更すると値が変更される
	var i int = 100
	one(&i)
	fmt.Println(i) // 100
}
