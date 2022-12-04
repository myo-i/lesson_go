package _range

import (
	"fmt"
)

func Range() {
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}

	// rangeでインデックスを使う場合
	for i, v := range l {
		fmt.Println(i, v)
	}

	// インデックスを使わない場合
	for _, v := range l {
		fmt.Println(v)
	}

	// ここまでスライスの場合
	// キーのみだけは変数を1つにしても取り出すことができる

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// vを省略した場合_で省略できる
	for k, _ := range m {
		fmt.Println(k)
	}
	// kを省略することもできるが、第一返り値のKは忘れないようにする
	for _, v := range m {
		fmt.Println(v)
	}

}
