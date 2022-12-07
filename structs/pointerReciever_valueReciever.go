package _structs

import "fmt"

type Vertex struct {
	X, Y int
}

// / Area こいつはメソッド（v.Area()で呼び出せる）
func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// / Area こいつは関数
func Area(v Vertex) int {
	return v.X * v.Y
}

func PointerValue() {
	v := Vertex{3, 4}
	//fmt.Println(Area(v))

	// vertex.Areaとしているのでコードしてわかりやすく、インテリセンスもしてくれて利用できる物がわかりやすい
	v.Scale(10)
	fmt.Println(v.Area())
}
