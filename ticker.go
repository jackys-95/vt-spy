package main 

type Ticker struct {
    symbol, exchange, currency string
}

func NewTicker(symbol, exchange, currency string) *Ticker {
    ticker := new(Ticker)
    ticker.symbol = symbol
    ticker.exchange = exchange
    ticker.currency = currency
    return ticker
}

