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

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// / Area こいつは関数
//func Area(v Vertex) int {
//	return v.x * v.y
//}

// / New コンストラクタ
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func PointerValue() {
	//v := Vertex{3, 4}
	//fmt.Println(Area(v))

	// コンストラクタパターンは一種のデザインパターン
	v := New(3, 4, 5)

	// vertex.Areaとしているのでコードしてわかりやすく、インテリセンスもしてくれて利用できる物がわかりやすい
	//v.Scale(10)
	v.Scale3D(10)
	//fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}
