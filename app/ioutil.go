package app

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

// 既に開いた状態にしておいて必要な時に呼び出していく
var templates = template.Must(template.ParseFiles("app/"+"edit.html", "app/"+"view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//t, _ := template.ParseFiles("app/" + tmpl + ".html")
	//t.Execute(w, p)
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RequestはURLで打ち込んだ/view/testの部分のこと
// ResponseWriterはレスポンス
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// test.txtを読み込んでその中身を返している
	p, err := loadPage(title)
	if err != nil {
		// loadPageでエラーが発生したらeditにリダイレクト
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
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

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// わかりにくい...
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// 下記のfnがsaveHandlerとかeditHandlerのこと
		fn(w, r, m[2])
	}
}

func Ioutil() {
	// /view/を受け取ったらviewHandlerへ飛ばす
	// Handlefuncの第二引数はfunc(ResponseWriter, Request)しか受け付けない（定義みろ！！）
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// Handlerを登録したいのであれば以下を実行してサーバーを立てる前に上記のようにハンドラーを登録する必要がある
	// 上記のハンドラーが登録出来たら、以下のメソッドでレスポンスを返す（サーバーも起動）
	log.Fatal(http.ListenAndServe(":8080", nil))
}
