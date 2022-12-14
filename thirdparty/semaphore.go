package _thirdparty

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 引数の数字をいじって走る数を制限できる
var s *semaphore.Weighted = semaphore.NewWeighted(2)

func longProcess(ctx context.Context) {
	// 指定した数のgoroutineが走っている場合は他のgoroutineが来ても必要ない場合使ったりする
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	//// ここでロックしてこのプロセスが終わってからリリースしている
	//if err := s.Acquire(ctx, 1); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// AcquireとTryAcquireはブロックするかそのまま処理を終わらせるか
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func Semaphore() {
	// goroutineが走る数を限定できる
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
