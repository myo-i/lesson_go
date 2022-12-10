package _goroutine

import (
	"fmt"
	"time"
)

// ソースコードはA Tour of GO「Default Section」
func DefaultBreak() {
	tick := time.Tick(100 * time.Millisecond) // Tickは返り値がチャネル。書きも同じ
	boom := time.After(500 * time.Millisecond)
Hogehoge: // 名前は何でもよい
	for {
		select {
		// チャネルは必要なければ書かなくてもよい
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break Hogehoge // どこのループを抜けるのか指定できる
			//break // breakだと無限ループしてしまう
			//return // returnだとループ後の処理が実行されない
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	// 上記のreturnでは下記のコードは実行されない
	fmt.Println("###############")
}
