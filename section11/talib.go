package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Println(spy.CSV())
	// Rsi = 株価が売られているか買われているかの指標、グラフ
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	// moving average 過去14日間
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
