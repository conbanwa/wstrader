package exx

import (
	"github.com/conbanwa/wstrader/cons"
	"net/http"
	"net/url"
	"testing"
)

var (
	apiKey       = ""
	apiSecretkey = "yourSecretKey"
	exx          = New(http.DefaultClient, apiKey, apiSecretkey)
)

func skipKey(t *testing.T) {
	if apiKey == "" {
		t.Skip("Skipping testing without TestKey")
	}
}
func TestExx_Signed(t *testing.T) {
	params := url.Values{}
	exx.accessKey = apiKey
	exx.secretKey = apiSecretkey
	exx.buildPostForm(&params)
	t.Log(params)
}
func TestExx_GetAccount(t *testing.T) {
	skipKey(t)
	acc, err := exx.GetAccount()
	t.Log(acc, err)
	t.Log(acc.SubAccounts[cons.BTC])
}
func TestExx_GetTicker(t *testing.T) {
	skipKey(t)
	ticker, err := exx.GetTicker(cons.BTC_USD)
	t.Log(ticker, err)
}
func TestExx_GetDepth(t *testing.T) {
	dep, _ := exx.GetDepth(2, cons.BTC_USDT)
	t.Log(dep)
	t.Log(dep.AskList[0])
	t.Log(dep.BidList[0])
}
func TestExx_LimitSell(t *testing.T) {
	skipKey(t)
	ord, err := exx.LimitSell("0.001", "75000", cons.NewCurrencyPair2("BTC_QC"))
	t.Log(err)
	t.Log(ord)
}
func TestExx_LimitBuy(t *testing.T) {
	skipKey(t)
	ord, err := exx.LimitBuy("2", "4", cons.NewCurrencyPair2("1ST_QC"))
	t.Log(err)
	t.Log(ord)
}
func TestExx_CancelOrder(t *testing.T) {
	skipKey(t)
	r, err := exx.CancelOrder("201802014255365", cons.NewCurrencyPair2("BTC_QC"))
	t.Log(err)
	t.Log(r)
}
func TestExx_GetUnfinishOrders(t *testing.T) {
	skipKey(t)
	ords, err := exx.GetUnfinishedOrders(cons.NewCurrencyPair2("1ST_QC"))
	t.Log(err)
	t.Log(ords)
}
func TestExx_GetOneOrder(t *testing.T) {
	skipKey(t)
	ord, err := exx.GetOneOrder("20180201341043", cons.NewCurrencyPair2("1ST_QC"))
	t.Log(err)
	t.Log(ord)
}
