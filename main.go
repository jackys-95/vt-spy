package main

import (
    "fmt"
    "flag"
    "strings"
)

const WatchlistFlagName = "watchlist"
const TickerFlagName = "ticker"
const TickersFlagName = "tickers"
var watchListFlag *bool
var tickerFlag *string
var tickersFlag *string

func init() {
    watchListFlag = flag.Bool(WatchlistFlagName, false, "True or false.")
    tickerFlag = flag.String(TickerFlagName, "", "Symbol for the stock ticker you wish to query.")
    tickersFlag = flag.String(TickersFlagName, "", "A list of stock symbols.")
}

func parseTickersFlag(tickers string) []Ticker {
    tickerStrings := strings.Split(tickers, " ")
    var tickerStructs []Ticker
    for _, ticker := range tickerStrings {
        // TODO: dynamically determine NASDAQ/NYSE and currency.
        if validateTickerSymbol(ticker) {
            tickerStructs = append(tickerStructs, *NewTicker(ticker, "NASDAQ", "USD"))
        } else {
            fmt.Printf("%v is not a valid ticker symbol.\n", ticker)
        }
    }
    fmt.Println(tickers)
    return tickerStructs
}

func main() {
    flag.Parse()

    fmt.Println("Watchlist has value:", *watchListFlag)
    fmt.Println("Ticker has value:", *tickerFlag)
    fmt.Println("Tickers has value:", *tickersFlag)

    ticker := NewTicker(*tickerFlag, "NASDAQ", "USD")
    fmt.Println(*ticker)
    fmt.Println(validateTickerSymbol(ticker.Symbol))
    WriteTickerToJsonFile(ticker)

    fmt.Println(parseTickersFlag(*tickersFlag))
}

