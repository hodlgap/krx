package main

import (
	"github.com/hodlgap/krx"
	"time"
)

/*
	Download csv file and parse to krx.Stock
*/

func main() {
	yesterday := time.Now().AddDate(0, 0, -1)

	code, err := krx.GetDownloadCode(yesterday)
	if err != nil {
		panic(err)
	}

	buf, err := krx.GetStockFile(code)
	if err != nil {
		panic(err)
	}

	stocks, err := krx.ParseStockBuf(buf)
	if err != nil {
		panic(err)
	}

	println(len(stocks))
}
