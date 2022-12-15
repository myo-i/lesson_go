package app

import (
	"fmt"
	"io/ioutil"
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

func Ioutil() {
	p1 := &Page{Title: "test", Body: []byte("Sample Data")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Sprintf(string(p2.Body))
}
