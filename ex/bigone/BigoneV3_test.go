package bigone

import (
	. "github.com/conbanwa/exws"
	"github.com/conbanwa/exws/config"
	. "github.com/conbanwa/exws/cons"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(config.GetProxy(true))
		},
	},
	Timeout: 10 * time.Second,
}
var (
	apikey    = ""
	secretkey = ""
	b1        = NewV3(httpProxyClient, apikey, secretkey)
)

func TestNewV3(t *testing.T) {
	return
	b1.setTimeOffset()
}
func TestBigoneV3_GetTicker(t *testing.T) {
	return
	t.Log(b1.GetTicker(ETH_BTC))
}
func TestBigoneV3_GetDepth(t *testing.T) {
	return
	t.Log(b1.GetDepth(1, ETH_BTC))
}
func TestBigoneV3_GetAccount(t *testing.T) {
	return
	t.Log(b1.GetAccount())
}
func TestBigoneV3_GetUnfinishOrders(t *testing.T) {
	return
	t.Log(b1.GetUnfinishedOrders(BTC_USDT))
}
func TestBigoneV3_GetOrderHistorys(t *testing.T) {
	return
	t.Log(b1.GetOrderHistorys(BTC_USDT, OptionalParameter{"test": 1}))
}
func TestBigoneV3_LimitSell(t *testing.T) {
	return
	t.Log(b1.LimitSell("322", "1", BTC_USDT))
}
func TestBigoneV3_LimitBuy(t *testing.T) {
	return
	t.Log(b1.LimitBuy("10", "1", BTC_USDT))
}
func TestBigoneV3_CancelOrder(t *testing.T) {
	return
	t.Log(b1.CancelOrder("570658164", BTC_USDT))
}
func TestBigoneV3_GetOneOrder(t *testing.T) {
	return
	t.Log(b1.GetOneOrder("570658164", BTC_USDT))
}
func TestBigoneV3_GetKlineRecords(t *testing.T) {
	return
	t.Log(b1.GetKlineRecords(ETH_BTC, KLINE_PERIOD_1MIN, 1, 1))
}
