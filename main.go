package main

import (
    "fmt"
    "flag"
    "regexp"
)

const WatchlistFlagName = "watchlist"
const TickerFlagName = "ticker"
const NYSEregex = `[a-zA-Z]{4}`
const NASDAQregex = `[a-zA-Z]{4}`

var watchListFlag *bool
var tickerFlag *string

func init() {
    watchListFlag = flag.Bool(WatchlistFlagName, false, "True or false.")
    tickerFlag = flag.String(TickerFlagName, "", "Symbol for the stock ticker you wish to query.")
}

func validateTickerSymbol(symbol string) bool {
    nyseRe, nasdaqRe := regexp.MustCompile(NYSEregex), regexp.MustCompile(NASDAQregex)
    return nyseRe.MatchString(symbol) || nasdaqRe.MatchString(symbol)
}

func validateTickers() {
}

func main() {
    flag.Parse()
    fmt.Println("Watchlist has value:", *watchListFlag)
    fmt.Println("Ticker has value:", *tickerFlag)
    ticker := NewTicker(*tickerFlag, "NASDAQ", "USD")
    fmt.Println(*ticker)
    fmt.Println(validateTickerSymbol(ticker.Symbol))
    WriteTickerToJsonFile(ticker)
}

