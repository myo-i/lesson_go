package slice_make_cap

import "fmt"

func MakeCap() {
	// length:3, capacity:5
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// capacityを超えた場合
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// 引数が1つの場合
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	// 0を入れる
	b := make([]int, 0)
	// nil(null)を入れる
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	// 演習[実行結果はどうなるのか？]
	//d := make([]int, 5)
	d := make([]int, 0, 5)
	for i := 1; i <= 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)

}
