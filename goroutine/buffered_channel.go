package _goroutine

import "fmt"

func Buffer() {
	// チャネルを２つまで受け取る
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	//ch <- 300 ←　チャネルを使用してない状態で追加するとエラー
	x := <-ch
	// 最初に入ったチャネルが消費される
	fmt.Println(x)
	fmt.Println(len(ch))

	ch <- 300
	fmt.Println(len(ch))

	// チャネルのクローズをしないとrangeで範囲外でも読み込もうとしてしまう
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
