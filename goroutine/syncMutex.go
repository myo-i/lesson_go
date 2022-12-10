package _goroutine

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

// 値を変更するのでポインタ
func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key] += 1
}

// 読み込み
func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func SyncMutex() {
	// 1つのマップに対して2つのgoroutineが書き込もうとするとエラーが発生するときがある
	// チャネルを使わずに、この問題を回避するためにsync.Mutexを用いる
	//c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	// 1つのマップに対して2つのgoroutineが同時に書き込む可能性もあるし
	// key:1の時に同時に読み込むと同じ数字に対して数字を入れてしまう時も出てくる
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"] += 1
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"] += 1
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("key"))
}
