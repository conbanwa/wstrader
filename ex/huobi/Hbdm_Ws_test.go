package huobi

import (
	"github.com/conbanwa/exws"
	"github.com/conbanwa/exws/cons"
	"github.com/conbanwa/exws/q"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHbdmWs(t *testing.T) {
	ws := NewHbdmWs()
	ws.SetCallbacks(func(ticker *exws.FutureTicker) {
		t.Log(ticker.Ticker)
	}, func(depth *exws.Depth) {
		t.Log(">>>>>>>>>>>>>>>")
		t.Log(depth.ContractType, depth.Pair)
		t.Log(depth.BidList)
		t.Log(depth.AskList)
	}, func(trade *q.Trade, s string) {
		t.Log(s, trade)
	})
	assert.Nil(t, ws.SubscribeTicker(cons.BTC_USD, cons.QUARTER_CONTRACT))
	assert.Nil(t, ws.SubscribeDepth(cons.BTC_USD, cons.NEXT_WEEK_CONTRACT))
	assert.Nil(t, ws.SubscribeTrade(cons.LTC_USD, cons.THIS_WEEK_CONTRACT))
	time.Sleep(time.Second)
}
