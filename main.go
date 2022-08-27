package main

import (
    "fmt"
    "flag"
)

const WatchlistFlagName = "watchlist"
const TickerFlagName = "ticker"

var watchListFlag *bool
var tickerFlag *string

func init() {
    watchListFlag = flag.Bool(WatchlistFlagName, false, "True or false.")
    tickerFlag = flag.String(TickerFlagName, "", "Symbol for the stock ticker you wish to query.")
}

func main() {
    flag.Parse()
    fmt.Println("Watchlist has value:", *watchListFlag)
    fmt.Println("Ticker has value:", *tickerFlag)
    ticker := NewTicker(*tickerFlag, "NASDAQ", "USD")
    fmt.Println(*ticker)
    WriteTickerToJsonFile(ticker)
}

