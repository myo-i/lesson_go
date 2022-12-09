package _structs

import "fmt"

type Num struct {
	X, Y int
}

func (n Num) Plus() int {
	return n.X + n.Y
}

func (n Num) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", n.X, n.Y)
}

func Exercise() {
	v := Num{3, 4}
	fmt.Println(v.Plus())
	fmt.Println(v)
}
