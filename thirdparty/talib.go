package _thirdparty

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func Talib() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	// RSIは株価の指標の1つ
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

	// EMAも指標の1つ。14は過去14日間
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
