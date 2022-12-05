package _defer

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func Defer() {
	foo()

	// deferはdeferが定義されている関数の実行が終了したら実行するもの
	defer fmt.Println("world")
	fmt.Println("hello")

	// deferは後ろから実行されていく
	// run
	// 3
	// 2
	// 1
	// success
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// deferの使いどころ
	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
