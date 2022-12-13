package _stdpackage

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}
func Context() {
	ch := make(chan string)

	// 処理に時間がかかった場合、キャンセルするよう要請できる
	ctx := context.Background()
	// 3秒かかっても終わらなかったら強制終了（あくまでgoroutineに対して）
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	// contextだけ設定してタイムアウトも何もしたくない場合、下記を渡せばよい
	//ctx := context.TODO()

	defer cancel()
	go longProcess(ctx, ch)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-ch:
			fmt.Println("success")
			return
		}
	}
}
