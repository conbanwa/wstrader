package huobi

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
	apiSecretkey = ""
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   config.Proxy}, nil
		},
	},
	Timeout: 10 * time.Second,
}
var hbpro *HuoBiPro

func init() {
	hbpro = NewHuoBiProSpot(httpProxyClient, apiKey, apiSecretkey)
}
func skipKey(t *testing.T) {
	if apiKey == "" {
		t.Skip("Skipping testing without apiKey")
	}
}
func TestHuobiPro_GetTicker(t *testing.T) {
	ticker, err := hbpro.GetTicker(cons.XRP_BTC)
	assert.Nil(t, err)
	t.Log(ticker)
}
func TestHuobiPro_GetDepth(t *testing.T) {
	dep, err := hbpro.GetDepth(2, cons.LTC_USDT)
	assert.Nil(t, err)
	t.Log(dep.AskList)
	t.Log(dep.BidList)
}
func TestHuobiPro_GetAccountInfo(t *testing.T) {
	skipKey(t)
	info, err := hbpro.GetAccountInfo("point")
	assert.Nil(t, err)
	t.Log(info)
}

// 获取点卡剩余
func TestHuoBiPro_GetPoint(t *testing.T) {
	skipKey(t)
	point := NewHuoBiProPoint(httpProxyClient, apiKey, apiSecretkey)
	acc, _ := point.GetAccount()
	t.Log(acc.SubAccounts[HBPOINT])
}

// 获取现货资产信息
func TestHuobiPro_GetAccount(t *testing.T) {
	skipKey(t)
	acc, err := hbpro.GetAccount()
	assert.Nil(t, err)
	t.Log(acc.SubAccounts)
}
func TestHuobiPro_LimitBuy(t *testing.T) {
	skipKey(t)
	ord, err := hbpro.LimitBuy("", "0.09122", cons.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}
func TestHuobiPro_LimitSell(t *testing.T) {
	skipKey(t)
	ord, err := hbpro.LimitSell("1", "0.212", cons.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}
func TestHuobiPro_MarketSell(t *testing.T) {
	skipKey(t)
	ord, err := hbpro.MarketSell("0.1738", "0.212", cons.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}
func TestHuobiPro_MarketBuy(t *testing.T) {
	skipKey(t)
	ord, err := hbpro.MarketBuy("0.02", "", cons.BCC_BTC)
	assert.Nil(t, err)
	t.Log(ord)
}
func TestHuobiPro_GetUnfinishOrders(t *testing.T) {
	skipKey(t)
	ords, err := hbpro.GetUnfinishedOrders(cons.ETC_USDT)
	assert.Nil(t, err)
	t.Log(ords)
}
func TestHuobiPro_CancelOrder(t *testing.T) {
	skipKey(t)
	r, err := hbpro.CancelOrder("600329873", cons.ETH_USDT)
	assert.Nil(t, err)
	t.Log(r)
	t.Log(err)
}
func TestHuobiPro_GetOneOrder(t *testing.T) {
	skipKey(t)
	ord, err := hbpro.GetOneOrder("165062634284339", cons.BTC_USDT)
	assert.Nil(t, err)
	t.Log(ord)
}
func TestHuobiPro_GetOrderHistorys(t *testing.T) {
	ords, err := hbpro.GetOrderHistorys(
		cons.NewCurrencyPair2("BTC_USDT"),
		exws.OptionalParameter{}.Optional("start-date", "2020-11-30"))
	t.Log(err)
	t.Log(ords)
}
func TestHuobiPro_GetCurrenciesList(t *testing.T) {
	hbpro.GetCurrenciesList()
}
func TestHuobiPro_GetCurrenciesPrecision(t *testing.T) {
	res, err := hbpro.GetCurrenciesPrecision()
	assert.Nil(t, err)
	t.Log("total: ", len(res), res[0], res[1], res[2])
}
