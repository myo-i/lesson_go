package _stdpackage

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func Ioutil() {
	//// ioutilは非推奨（os.Openでも同じように使える）
	//// 実行するのはmain.goなのでmain.goとパスを渡してもよい
	//content, err := ioutil.ReadFile("main.go")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(string(content))
	//
	//if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	//	log.Fatalln(err)
	//}

	// ネットワークから来たデータを一時的に保存しておいたりする記憶方法？
	r := bytes.NewReader([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
