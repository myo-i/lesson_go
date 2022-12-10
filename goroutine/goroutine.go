package _goroutine

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Goroutine() {
	var wg sync.WaitGroup
	// wgに1つの並列処理があることを伝える
	wg.Add(1)
	// goroutineのスレッドが生成されて実行される前にプログラムが終了するとgoroutineも終了してしまう
	go goroutine("world", &wg)
	normal("hello")

	// wgがDoneするまで待つ
	wg.Wait()
}
