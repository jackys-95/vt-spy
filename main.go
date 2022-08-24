package main

import (
    "fmt"
    "flag"
)

const Watchlist = "watchlist"
const Ticker = "ticker"

var watchListFlag *bool
var tickerFlag *string

func init() {
    watchListFlag = flag.Bool(Watchlist, false, "True or false.")
    tickerFlag = flag.String(Ticker, "", "Symbol for the stock ticker you wish to query.")
}

func main() {
    flag.Parse()
    fmt.Println("Watchlist has value: ", *watchListFlag)
    fmt.Println("Ticker has value: ", *tickerFlag)
}

