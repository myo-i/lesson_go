package _structs

import "fmt"

type Vertex struct {
	x, y int
}

// / Area こいつはメソッド（v.Area()で呼び出せる）
func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// / Area こいつは関数
//func Area(v Vertex) int {
//	return v.x * v.y
//}

/// New コンストラクタ
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func PointerValue() {
	//v := Vertex{3, 4}
	//fmt.Println(Area(v))

	// コンストラクタパターンは一種のデザインパターン
	v := New(3, 4)

	// vertex.Areaとしているのでコードしてわかりやすく、インテリセンスもしてくれて利用できる物がわかりやすい
	v.Scale(10)
	fmt.Println(v.Area())
}
