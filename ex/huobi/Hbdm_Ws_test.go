package huobi

import (
	"log"
	"qa3/wstrader"
	"qa3/wstrader/config"
	"qa3/wstrader/cons"
	"qa3/wstrader/q"
	"testing"
	"time"
)

func TestNewHbdmWs(t *testing.T) {
	ws := NewHbdmWs()
	ws.ProxyUrl("socks5://" + config.Proxy)
	ws.SetCallbacks(func(ticker *wstrader.FutureTicker) {
		log.Println(ticker.Ticker)
	}, func(depth *wstrader.Depth) {
		log.Println(">>>>>>>>>>>>>>>")
		log.Println(depth.ContractType, depth.Pair)
		log.Println(depth.BidList)
		log.Println(depth.AskList)
		log.Println("<<<<<<<<<<<<<<")
	}, func(trade *q.Trade, s string) {
		log.Println(s, trade)
	})
	t.Log(ws.SubscribeTicker(cons.BTC_USD, cons.QUARTER_CONTRACT))
	t.Log(ws.SubscribeDepth(cons.BTC_USD, cons.NEXT_WEEK_CONTRACT))
	t.Log(ws.SubscribeTrade(cons.LTC_USD, cons.THIS_WEEK_CONTRACT))
	time.Sleep(time.Minute)
}