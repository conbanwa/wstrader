package bitfinex

import (
	"log"
	"qa3/wstrader"
	"qa3/wstrader/cons"
	"qa3/wstrader/q"
	"testing"
	"time"
)

func TestNewBitfinexWs(t *testing.T) {
	bitfinexWs := NewWs()
	handleTicker := func(ticker *wstrader.Ticker) {
		log.Printf("Ticker: %+v: ", ticker)
	}
	handleTrade := func(trade *q.Trade) {
		log.Printf("Trade: %+v: ", trade)
	}
	handleCandle := func(candle *wstrader.Kline) {
		log.Printf("Candle: %+v: ", candle)
	}
	bitfinexWs.SetCallbacks(handleTicker, handleTrade, handleCandle)
	//Ticker
	t.Log(bitfinexWs.SubscribeTicker(cons.BTC_USD))
	t.Log(bitfinexWs.SubscribeTicker(cons.LTC_USD))
	//Trades
	t.Log(bitfinexWs.SubscribeTrade(cons.BTC_USD))
	//Candles
	t.Log(bitfinexWs.SubscribeCandle(cons.BTC_USD, cons.KLINE_PERIOD_1MIN))
	time.Sleep(time.Minute)
}