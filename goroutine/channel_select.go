package _goroutine

import (
	"fmt"
	"time"
)

func goroutineS1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutineS2(ch chan int) {
	for {
		ch <- 100
		time.Sleep(1 * time.Second)
	}
}

func ChannelSelect() {
	c1 := make(chan string)
	c2 := make(chan int)
	go goroutineS1(c1)
	go goroutineS2(c2)

	for {
		// 同時にチャネルの受信を行える
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
