package app

import (
	"html/template"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("app/" + tmpl + ".html")
	t.Execute(w, p)
}

// RequestはURLで打ち込んだ/view/testの部分のこと
// ResponseWriterはレスポンス
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/testを受け取ったらtestを抽出
	title := r.URL.Path[len("/view/"):]
	// test.txtを読み込んでその中身を返している
	p, err := loadPage(title)
	if err != nil {
		// loadPageでエラーが発生したらeditにリダイレクト
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	// edit.htmlにname="body"があるのでそこから取得
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// wとrを渡して/view/+titleにリダイレクト
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func Ioutil() {
	// /view/を受け取ったらviewHandlerへ飛ばす
	// Handlefuncの第二引数はfunc(ResponseWriter, Request)（定義みろ！！）
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	// Handlerを登録したいのであれば以下を実行してサーバーを立てる前に上記のようにハンドラーを登録する必要がある
	// 上記のハンドラーが登録出来たら、以下のメソッドでレスポンスを返す（サーバーも起動）
	log.Fatal(http.ListenAndServe(":8080", nil))
}
