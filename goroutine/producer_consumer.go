package _goroutine

import (
	"fmt"
	"sync"
)

func producer(i int, ch chan int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	// main関数ではなくgoroutine上でrangeを行っており、chに値が入ってくるのを待つことになるのでcloseを実行する
	// closeされた段階でchに値が入らないことを認識し、for文が終了する
	for i := range ch {
		fmt.Println(i * 100)
		wg.Done()
	}
	// 10回目のDoneが実行された時点でWait()が実行されるため
	fmt.Println("####################")
}

func ProduceConsume() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		// 上記条件のループでwg.Add(1)を実行するということはwg.Done()を10回するということ
		wg.Add(1)
		go producer(i, ch)
	}

	go consumer(ch, &wg)
	// 10回目のDoneが実行された時点でWait()が実行される
	wg.Wait()
	close(ch)
	// ①closeされてからfunc consumer内のchに値が入らないことが認識されてfunc consumer内のfor文が終了する。
	// ②Sleepを設けることでclose(ch)からプログラムが終了するまでに余裕が生まれるので、func consumer内のfor文以下の処理が実行できる。
	// ③Sleepを設けない場合にfor文以下の処理が実行されないのは、closeされてからfunc consumer内のchに値が入らないことが認識されてfor文が終了する前に
	// プログラムが終了してしまうから
	//time.Sleep(2 * time.Second)
	//fmt.Println("Done")
}
