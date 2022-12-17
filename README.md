# lesson_go

## 学び
goは1つのパッケージが１つのファイルのようなもので、同じ関数名や変数名が使えない。  
goは1つのディレクトリに1つのパッケージしか存在できない。

## GO testing
Ginkgo
https://pkg.go.dev/github.com/onsi/ginkgo

GOMEGA
https://onsi.github.io/gomega/

## gofmt
gofmt [対象のファイル名] でコードのフォーマットを表示  
gofmt -w [対象のファイル名] でフォーマットを適用

## パッケージの検索
https://pkg.go.dev/

## Go Webアプリ入門
https://go.dev/doc/articles/wiki/

## 様々な関数の使い方
https://www.wakuwakubank.com/posts/778-go-func/

## ファイルの読み書き
Openは書き込み権限なしで読み込み、Createは書き込み権限ありで読み込み（すでに存在するファイルを指定した場合は空にする）  
https://zenn.dev/hsaki/books/golang-io-package/viewer/file
