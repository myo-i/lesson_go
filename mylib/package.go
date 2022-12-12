package mylib

import (
	"fmt"
	"github.com/markcheno/go-quote"
	// 名前を変えることもできる
	changename "github.com/markcheno/go-talib"

	// 先頭に_を付けることで使われなくても得r－が出ないようにできる
	_ "lesson_go/goroutine"
)

func Packages() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := changename.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
