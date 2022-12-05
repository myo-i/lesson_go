package exercises

import "fmt"

func problem1() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min int
	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}
		if min > v {
			min = v
		}
	}

	fmt.Println(min)
}

func problem2() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var sum int
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}

func Exercise() {
	problem1()

	problem2()
}
