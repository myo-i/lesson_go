package _func

import "fmt"

type Test func(string2 string, int2 int)

func ttt(s string, n int) {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	fmt.Println(result)

}

func test(fn func(string2 string, int2 int)) Test {
	return func(s string, i int) {
		if i < 3 {
			s += " world"
			i += 3
			fn(s, i)
		}
		fn("Hello World", 6)
	}
}

func AltimateTest(s string, n int, test func(string2 string, int2 int)) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	test("hello", 2)
	return result
}

// golangは返り値を２つ設定できる
func add(x, y int) (int, int) {
	return x + y, x - y
}

// 関数の返り値に変数の宣言と型を指定した場合
// 変数は既に宣言されているのでresult := ~ のように宣言できない
// さらに関数はint型のresultを返すとわかっているのでreturnの後に何も書かなくても返す
// 引数が多い場合などに使うと、何を返すのかわかりやすい（引数や返り値が明示的な場合は使わなくてよい）
func cal(price, item int) (result int) {
	result = price * item
	return
}

func Func() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	// 関数を変数に入れて実行することができる
	f := func(s string) {
		fmt.Println("inner func", s)
	}
	f("Hello")

	// また変数を宣言しなくても引数を渡すことができる
	func(s string) {
		fmt.Println("inner func", s)
	}("World")

	// type Hoge funcの実験
	AltimateTest("world", 1, test(ttt))
}
