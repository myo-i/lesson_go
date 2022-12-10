package _goroutine

import "fmt"

func routine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

// ChRangeClose channelとは例えばmain関数とgoroutineのデータの受け渡しのこと
func ChRangeClose() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // どんどん実行して受け取っている
	go routine1(s, c)

	// チャネルに値が入る度、随時実行される
	for ch := range c {
		fmt.Println(ch)
	}

}
