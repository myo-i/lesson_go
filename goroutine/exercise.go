package _goroutine

import "fmt"

func goroutineE(s []string, c chan string) {
	sum := ""
	defer close(c)
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func Exercise() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutineE(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
