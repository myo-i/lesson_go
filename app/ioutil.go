package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// ストラクトのbodyのみファイルに書き込んでsave
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// saveしたデータが入っているファイルの読み込み
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// RequestはURLで打ち込んだ/view/testの部分のこと
// ResponseWriterはレスポンス
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/testを受け取ったらtestを抽出
	title := r.URL.Path[len("/view/"):]
	// test.txtを読み込んでその中身を返している
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func Ioutil() {
	// /view/を受け取ったらviewHandlerへ飛ばす
	// Handlefuncの第二引数はfunc(ResponseWriter, Request)（定義みろ！！）
	http.HandleFunc("/view/", viewHandler)
	// Handlerを登録したいのであれば以下を実行してサーバーを立てる前に上記のようにハンドラーを登録する必要がある
	// 上記のハンドラーが登録出来たら、以下のメソッドでレスポンスを返す（サーバーも起動）
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
