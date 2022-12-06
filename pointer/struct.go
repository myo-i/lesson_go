package _pointer

import "fmt"

// Vertex
// struct内で小文字で宣言するとprivateという意味になるため
// 内部からしか書き込んだり読み込んだりできない
type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

// Struct これがストラクト
func Struct() {
	/*
		v := Vertex{X: 1, Y: 2}
		fmt.Println(v)
		fmt.Println(v.X, v.Y)
		v.X = 100
		fmt.Println(v)

		v2 := Vertex{X: 1}
		fmt.Println(v2)

		v3 := Vertex{1, 2, "hello"}
		fmt.Println(v3)

		// 値を何も入れない場合
		v4 := Vertex{}
		var v5 Vertex
		fmt.Printf("%T %v\n", v4, v4)
		fmt.Printf("%T %v\n", v5, v5)

		// newで初期化した場合
		v6 := new(Vertex)
		v7 := &Vertex{} // 書き方としてはこちらが推奨
		fmt.Printf("%T %v\n", v6, v6)
		fmt.Printf("%T %v\n", v7, v7)
	*/

	v := Vertex{1, 2, "test"}
	changeVertex(v) // 値渡しなのでもう一個コピーされているような状態なので本体は変更されない
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2) // 値渡しなのでもう一個コピーされているような状態なので本体は変更されない
	fmt.Println(*v2)
}
