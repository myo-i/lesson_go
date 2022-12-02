package _const

import "fmt"

const Pi = 3.14

const (
	Username = "test_name"
	Password = "test_pass"
)

func Const() {
	fmt.Println(Pi, Username, Password)
}
