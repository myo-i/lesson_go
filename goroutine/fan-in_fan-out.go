package _goroutine

import "fmt"

func producer1(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// 下記のように記述することでfirstは受信、secondは送信という役割を明示的に示すことができる
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for ch := range first {
		second <- ch * 2
	}
}

func multi4(second, third chan int) {
	defer close(third)
	for ch := range second {
		third <- ch * 4
	}
}

func FanInFanOut() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer1(first)
	go multi2(first, second)
	go multi4(second, third)

	for result := range third {
		fmt.Println(result)
	}
}
