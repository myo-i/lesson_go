package _if

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func If() {
	result := by2(10)
	if result == "ok" {
		fmt.Println("It's by 2!!")
	}

	// 上記を一行で書くと
	if result2 := by2(14); result2 == "ok" {
		fmt.Println("Great!!")
	}
	// 1行で書いた場合はif文の中でのみ変数を利用することができるため下記ではコンパイルエラーとなる
	//fmt.Println(result)
	//fmt.Println(result2)

	num := 6
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 10, 30
	if x == 10 && y == 20 {
		fmt.Println("AND")
	} else if x == 10 || y == 20 {
		fmt.Println("OR")
	}
}
