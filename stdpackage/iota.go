package _stdpackage

import "fmt"

const (
	// 連番を自動で生成
	c1 = iota
	c2
	c3
)

const (
	_ = iota
	// 10回シフト
	KB int = 1 << (10 * iota)
	// 20回シフト
	MB
	// 30回シフト
	GB
)

func Iota() {
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}
