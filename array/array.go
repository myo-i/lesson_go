package array

import "fmt"

func Array() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	// 配列は固定長なのでb.append(b, 300)と記述してもコンパイルエラーが発生する
	fmt.Println(b)

	// スライスは可変長
	var c []int = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)
}
