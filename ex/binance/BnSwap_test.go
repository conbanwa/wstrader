package binance

import (
	"github.com/conbanwa/exws"
	"github.com/conbanwa/exws/config"
	"github.com/conbanwa/exws/cons"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const (
	apiKey       = ""
	apiSecretkey = "YOUR_KEY_SECRET"
)

func skipKey(t *testing.T) {
	if apiKey == "" {
		t.Skip("Skipping testing without apiKey")
	}
}

var bs = NewBinanceSwap(&exws.APIConfig{
	Endpoint: "https://testnet.binancefuture.com",
	HttpClient: &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(config.GetProxy(true))
			},
		},
		Timeout: 10 * time.Second,
	},
	ApiKey:       apiKey,
	ApiSecretKey: apiSecretkey,
})

func TestBinanceSwap_Ping(t *testing.T) {
	bs.Ping()
}
func TestBinanceSwap_GetFutureDepth(t *testing.T) {
	res, err := bs.GetFutureDepth(cons.BTC_USDT, cons.SWAP_CONTRACT, 1)
	assert.Nil(t, err)
	t.Log(res)
}
func TestBinanceSwap_GetFutureIndex(t *testing.T) {
	res, err := bs.GetFutureIndex(cons.BTC_USDT)
	assert.Nil(t, err)
	t.Log(res)
}
func TestBinanceSwap_GetKlineRecords(t *testing.T) {
	skipKey(t)
	kline, err := bs.GetKlineRecords("", cons.BTC_USDT, cons.KLINE_PERIOD_4H, 1, exws.OptionalParameter{"test": 0})
	t.Log(err, kline[0].Kline)
}
func TestBinanceSwap_GetTrades(t *testing.T) {
	skipKey(t)
	t.Log(bs.GetTrades("", cons.BTC_USDT, 0))
}
func TestBinanceSwap_GetFutureUserinfo(t *testing.T) {
	skipKey(t)
	t.Log(bs.GetFutureUserinfo())
}
func TestBinanceSwap_PlaceFutureOrder(t *testing.T) {
	skipKey(t)
	t.Log(bs.PlaceFutureOrder(cons.BTC_USDT, "", "8322", "0.01", cons.OPEN_BUY, 0, 0))
}
func TestBinanceSwap_PlaceFutureOrder2(t *testing.T) {
	skipKey(t)
	t.Log(bs.PlaceFutureOrder(cons.BTC_USDT, "", "8322", "0.01", cons.OPEN_BUY, 1, 0))
}
func TestBinanceSwap_GetFutureOrder(t *testing.T) {
	skipKey(t)
	t.Log(bs.GetFutureOrder("1431689723", cons.BTC_USDT, ""))
}
func TestBinanceSwap_FutureCancelOrder(t *testing.T) {
	skipKey(t)
	t.Log(bs.FutureCancelOrder(cons.BTC_USDT, "", "1431554165"))
}
func TestBinanceSwap_GetFuturePosition(t *testing.T) {
	skipKey(t)
	t.Log(bs.GetFuturePosition(cons.BTC_USDT, ""))
}
