package krx

type Stock struct {
	StockCode            string  `csv:"종목코드"`  // 종목코드
	StockName            string  `csv:"종목명"`   // 종목명
	MarketType           string  `csv:"시장구분"`  // E.g) KOSDAQ, KOSPI
	OrgName              string  `csv:"소속부"`   // E.g) 중견기업부, 우량기업부
	ClosingPrice         int64   `csv:"종가"`    // 종가
	DayOverDay           int64   `csv:"대비"`    // 전일 대비 상승가
	DayOverDayRate       float32 `csv:"등락률"`   // 전일 대비 상승률
	OpeningPrice         int64   `csv:"시가"`    // 시가
	MostExpensivePrice   int64   `csv:"고가"`    // 고가
	MostCheapestPrice    int64   `csv:"저가"`    // 저가
	TradingVolume        int64   `csv:"거래량"`   // 거래량
	TradingValue         int64   `csv:"거래대금"`  // 거래대금
	MarketCapitalization int64   `csv:"시가총액"`  // 시가총액
	TotalShares          int64   `csv:"상장주식수"` // 상장주식수
}
