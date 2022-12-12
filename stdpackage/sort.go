package _stdpackage

import (
	"fmt"
	"sort"
)

func Sort() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}

	// structは特定の関数のみで使う場合はこのように簡略化することもできる
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Bob", 21},
		{"Mike", 23},
		{"Alice", 25},
	}
	fmt.Println(i, s, p)

	sort.Ints(i)
	sort.Strings(s)
	// レス関数？
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)
}
