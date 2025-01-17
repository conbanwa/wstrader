package binance

import (
	"fmt"
	"github.com/conbanwa/exws"
	"github.com/conbanwa/exws/cons"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var ba = NewWithConfig(
	&exws.APIConfig{
		HttpClient: http.DefaultClient,
		Endpoint:   "https://api.binance.com",
	})

func TestFeeMap(t *testing.T) {
	t.Log(ba.TradeFee())
}
func TestBinance_GetTicker(t *testing.T) {
	ticker, err := ba.GetTicker(cons.NewCurrencyPair2("USDT_USD"))
	t.Log(ticker, err)
}
func TestBinance_LimitBuy(t *testing.T) {
	order, err := ba.LimitBuy("3", "0.1", cons.LTC_USDT)
	t.Log(order, err)
}
func TestBinance_LimitSell(t *testing.T) {
	order, err := ba.LimitSell("1", "90000", cons.LTC_USDT)
	t.Log(order, err)
}
func TestBinance_CancelOrder(t *testing.T) {
	if r, err := ba.CancelOrder("3848718241", cons.BTC_USDT); !r {
		t.Log(err)
	}
}
func TestBinance_GetOneOrder(t *testing.T) {
	res, err := ba.GetOneOrder("3874087228", cons.BTC_USDT)
	assert.Nil(t, err)
	t.Log(res)
}
func TestBinance_GetDepth(t *testing.T) {
	dep, err := ba.GetDepth(5, cons.NewCurrencyPair2("BTC_USDT"))
	if assert.Nil(t, err) {
		t.Log(dep.AskList)
		t.Log(dep.BidList)
	}
}
func TestBinance_GetAccount(t *testing.T) {
	account, err := ba.GetAccount()
	t.Log(account, err)
}
func TestBinance_GetUnfinishedOrders(t *testing.T) {
	orders, err := ba.GetUnfinishedOrders(cons.NewCurrencyPair2("BTC_USDT"))
	t.Log(orders, err)
}
func TestBinance_GetKlineRecords(t *testing.T) {
	startTime := time.Now().Add(-24*time.Hour).Unix() * 1000
	endTime := time.Now().Add(-5*time.Hour).Unix() * 1000
	kline, _ := ba.GetKlineRecords(cons.ETH_BTC,
		cons.KLINE_PERIOD_5MIN, 100,
		exws.OptionalParameter{}.
			Optional("startTime", fmt.Sprint(startTime)).
			Optional("endTime", fmt.Sprint(endTime)))
	for _, k := range kline {
		tt := time.Unix(k.Timestamp, 0)
		t.Log(tt, k.Open, k.Close, k.High, k.Low, k.Vol)
	}
}
func TestBinance_GetTrades(t *testing.T) {
	t.Log(ba.GetTrades(cons.BTC_USDT, 0))
}
func TestBinance_GetTradeSymbols(t *testing.T) {
	t.Log(ba.GetTradeSymbol(cons.BTC_USDT))
}
func TestBinance_SetTimeOffset(t *testing.T) {
	t.Log(ba.setTimeOffset())
	t.Log(ba.timeOffset)
}
func TestBinance_GetOrderHistorys(t *testing.T) {
	t.Log(ba.GetOrderHistorys(cons.BTC_USDT,
		exws.OptionalParameter{}.
			Optional("startTime", "1607656034333").
			Optional("limit", "5")))
}
