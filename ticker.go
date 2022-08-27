package main

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
)

type Ticker struct {
    Symbol, Exchange, Currency string
}

func NewTicker(symbol, exchange, currency string) *Ticker {
    ticker := new(Ticker)
    ticker.Symbol = symbol
    ticker.Exchange = exchange
    ticker.Currency = currency
    return ticker
}

func WriteTickerToJsonFile(ticker *Ticker) {
    jsonArray := ToJson(ticker)
    err := os.WriteFile("tickers.json", jsonArray, 0666)
    if err != nil {
        log.Fatal(err)
    }
}

func ToJson(ticker *Ticker) []byte {
    fmt.Println(*ticker)
    jsonArray, err := json.MarshalIndent(*ticker, "", " ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(jsonArray)
    return jsonArray
}

func FromJson(jsonArray []byte) Ticker { 
    var ticker Ticker
    err := json.Unmarshal(jsonArray, &ticker)
    if err != nil {
        log.Fatal(err)
    }
    return ticker
}

