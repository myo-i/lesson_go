package var_parameter

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	// iterateしてる。 item in lists みたいな感じ
	for _, param := range params {
		fmt.Println(param)
	}
}

func Parameter() {
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)
	// スライスを展開
	foo(s...)
}
