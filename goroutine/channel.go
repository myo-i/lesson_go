package _goroutine

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// channelに渡す
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// channelに渡す
	c <- sum
}

// Channel channelとは例えばmain関数とgoroutineのデータの受け渡しのこと
func Channel() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // どんどん実行して受け取っている
	go goroutine1(s, c)
	go goroutine2(s, c)
	// channelがsumを受け取るまで待っている
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
