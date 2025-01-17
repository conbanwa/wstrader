package build

import (
	"github.com/conbanwa/exws"
	"github.com/conbanwa/exws/cons"
	"github.com/conbanwa/exws/ex/binance"
	"github.com/conbanwa/exws/q"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	apiKey       = ""
	apiSecretkey = "YOUR_KEY_SECRET"
)

var api = DefaultAPIBuilder.APIKey(apiKey).APISecretkey(apiSecretkey)

func skipKey(t *testing.T) {
	if apiKey == "" {
		t.Skip("Skipping testing without apiKey")
	}
}

func TestFetchFutureDepthAndIndex(t *testing.T) {
	skipKey(t)
	binanceApi := api.Endpoint(binance.TestnetSpotWsBaseUrl).BuildFuture(cons.BINANCE_SWAP)
	depth, err := binanceApi.GetFutureDepth(cons.BTC_USD, cons.SWAP_USDT_CONTRACT, 100)
	assert.Nil(t, err)
	askTotalAmount, bidTotalAmount := 0.0, 0.0
	askTotalVol, bidTotalVol := 0.0, 0.0
	for _, v := range depth.AskList {
		askTotalAmount += v.Amount
		askTotalVol += v.Price * v.Amount
	}
	for _, v := range depth.BidList {
		bidTotalAmount += v.Amount
		bidTotalVol += v.Price * v.Amount
	}
	markPrice, err := binanceApi.GetFutureIndex(cons.BTC_USD)
	assert.Nil(t, err)
	t.Logf("CURRENT mark price: %f", markPrice)
	t.Logf("ContractType: %s ContractId: %s Pair: %s UTime: %s AmountTickSize: %d\n", depth.ContractType, depth.ContractId, depth.Pair, depth.UTime.String(), depth.Pair.AmountTickSize)
	t.Logf("askTotalAmount: %f, bidTotalAmount: %f, askTotalVol: %f, bidTotalVol: %f", askTotalAmount, bidTotalAmount, askTotalVol, bidTotalVol)
	t.Logf("ask price averge: %f, bid price averge: %f,", askTotalVol/askTotalAmount, bidTotalVol/bidTotalAmount)
	t.Logf("ask-bid spread: %f%%,", 100*(depth.AskList[0].Price-depth.BidList[0].Price)/markPrice)
}
func TestSubscribeSpotMarketData(t *testing.T) {
	skipKey(t)
	binanceWs, err := api.Endpoint(binance.TestnetFutureUsdBaseUrl).BuildSpotWs(cons.BINANCE)
	assert.Nil(t, err)
	binanceWs.TickerCallback(func(ticker *exws.Ticker) {
		t.Logf("%+v\n", *ticker)
	})
	binanceWs.SubscribeTicker(cons.BTC_USDT)
	binanceWs.DepthCallback(func(depth *exws.Depth) {
		t.Logf("%+v\n", *depth)
	})
	binanceWs.SubscribeDepth(cons.BTC_USDT)
	binanceWs.TradeCallback(func(trade *q.Trade) {
		t.Logf("%+v\n", *trade)
	})
	binanceWs.SubscribeTrade(cons.BTC_USDT)
	select {}
}

func TestSubscribeFutureMarketData(t *testing.T) {
	skipKey(t)
	binanceWs, err := api.Endpoint(binance.TestnetFutureUsdWsBaseUrl).BuildFuturesWs(cons.BINANCE_FUTURES)
	assert.Nil(t, err)
	binanceWs.TickerCallback(func(ticker *exws.FutureTicker) {
		//t.Logf("%+v\n", *ticker.Ticker)
	})
	binanceWs.SubscribeTicker(cons.BTC_USD, cons.SWAP_USDT_CONTRACT)
	binanceWs.DepthCallback(func(depth *exws.Depth) {
		t.Logf("%+v\n", *depth)
	})
	binanceWs.SubscribeDepth(cons.BTC_USDT, cons.SWAP_USDT_CONTRACT)
	binanceWs.TradeCallback(func(trade *q.Trade, contractType string) {
		t.Logf("%+v\n", *trade)
	})
	binanceWs.SubscribeTrade(cons.BTC_USDT, cons.SWAP_USDT_CONTRACT)
	select {}
}
