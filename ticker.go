package main

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
    "regexp"
    "time"
)

const watchlistFileName = "watchlist.json"
const nyseRegex = `[a-zA-Z]{4}`
const nasdaqRegex = `[a-zA-Z]{4}`

// TODO: These should be singleton
var watchlist Watchlist
var watchlistTickers map[string]bool

type Watchlist struct {
    Tickers  []Ticker `json:tickers`
}

type Ticker struct {
    Symbol      string `json:symbol`
    Exchange    string `json:exchange`
    Currency    string `json:currency`
    DateAdded   int64  `json:dateAdded`
}

func init() {
    watchlist = readWatchlistFromFile()
    watchlistTickers = buildWatchlistTickers(watchlist)
    fmt.Println(watchlist)
    fmt.Println(watchlistTickers)
}

func buildWatchlistTickers(watchlist Watchlist) map[string]bool {
    tickerMap := make(map[string]bool)
    for _, ticker := range watchlist.Tickers {
        tickerMap[ticker.Symbol] = true
    }
    return tickerMap
}

func NewTicker(symbol, exchange, currency string) *Ticker {
    ticker := new(Ticker)
    ticker.Symbol = symbol
    ticker.Exchange = exchange
    ticker.Currency = currency
    ticker.DateAdded = time.Now().Unix()
    return ticker
}

func readWatchlistFromFile() Watchlist {
    file, err := os.ReadFile(watchlistFileName)
    if err != nil {
        fmt.Println("Watchlist doesn't exist.")
        return Watchlist{make([]Ticker, 0)}
    }

    var fileWatchlist Watchlist
    unmarshalErr := json.Unmarshal(file, &fileWatchlist)
    if unmarshalErr != nil {
        log.Fatal(unmarshalErr)
    }

    return fileWatchlist
}

func addTickerToWatchlist(ticker Ticker) {
    if _, exists := watchlistTickers[ticker.Symbol]; exists {
        fmt.Printf("%v has already been added to the watchlist.\n",
                   ticker.Symbol)
    } else {
        watchlist.Tickers = append(watchlist.Tickers, ticker)
        watchlistTickers[ticker.Symbol] = true
    }

    writeWatchlistToJson(watchlist)
}

func writeWatchlistToJson(watchlist Watchlist) {
    watchlistJson := watchlistToJson(watchlist)
    err := os.WriteFile(watchlistFileName, watchlistJson, 0666)
    if err != nil {
        log.Fatal(err)
    }
}

func watchlistToJson(watchlist Watchlist) []byte {
    watchlistJson, err := json.MarshalIndent(watchlist, "", " ")
    if err != nil {
        log.Fatal(err)
    }
    return watchlistJson
}

func validateTickerSymbol(symbol string) bool {
    nyseRe, nasdaqRe := regexp.MustCompile(nyseRegex), regexp.MustCompile(nasdaqRegex)
    return nyseRe.MatchString(symbol) || nasdaqRe.MatchString(symbol)
}

