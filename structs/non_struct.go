package _structs

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)

	// typeで宣言するとmain.MyInt型になるため、intにキャストする必要がある
	// 返り値をMyIntとすればキャストしなくてもよい
	return int(i * 2)
}

func NonStruct() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())

}
