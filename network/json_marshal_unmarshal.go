package _network

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	// `json:"name"`はjsonで返す時は"name"にしてねという意味
	// 隠したいときはキーを"-"にする
	Name string `json:"name"`
	// `json:"age,string"`とすれば返す時の型も指定できる
	// `json:"age,omitempty"`で0や空になるときは隠すこともできる
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
	// ストラクトの場合はポインタにする必要がある
	T *T `json:"T,omitempty"`
}

// Unmarshalもカスタマイズできる

// 下記のようなメソッドを作ることでMarshalのカスタマイズができる
//func (p Person) MarshalJSON() ([]byte, error) {
//	v, err := json.Marshal(&struct {
//		Name string
//	}{
//		Name: "Mr." + p.Name,
//	})
//	return v, err
//}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func JsonUnmarshalMarshal() {
	// 下記のjsonがネットワークから来たと仮定する
	b := []byte(`{"name":"Mike", "age":0, "nickname":["a", "b", "c"]}`)
	var p Person
	// ネットワークから来たjsonをストラクトのキーに入れてくれる
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// ストラクトのキーに入れたものをjsonに戻したい場合（Unmarshalしたものをjsonとして返したい場合）
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
