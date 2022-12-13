# krx
![CodeQL](https://github.com/hodlgap/krx/actions/workflows/codeql.yml/badge.svg?branch=main&event=push)


arch/actions/workflows/codeql.yml) 
From download to entity

### Usage
```go
package main

import (
	"time"
	
	"github.com/hodlgap/krx"
)

func main() {
	// KRX only provides data up to the previous business day.
	yesterday := time.Now().AddDate(0, 0, -1)

	// Issuing temporary tokens for download
	code, err := krx.GetDownloadCode(yesterday)
	if err != nil {
		panic(err)
	}

	buf, err := krx.GetStockFile(code)
	if err != nil {
		panic(err)
	}

	// Parse to entity
	stocks, err := krx.ParseStockBuf(buf)
	if err != nil {
		panic(err)
	}
	
	for _, stock := range stocks {
        println(stock.StockName)
    }
}
```

### CSV
http://data.krx.co.kr/contents/MDC/MDI/mdiLoader/index.cmd?menuId=MDC0201020101
