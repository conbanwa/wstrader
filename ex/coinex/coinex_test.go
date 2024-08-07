package coinex

import (
	"fmt"
	"github.com/conbanwa/wstrader/cons"
	"net/http"
	"testing"
)

var coinex = New(http.DefaultClient, "", "")

func TestCoinEx_GetTicker(t *testing.T) {
	ticker, err := coinex.GetTicker(cons.LTC_BTC)
	t.Log(err)
	t.Log(ticker)
}
func TestCoinEx_GetDepth(t *testing.T) {
	dep, err := coinex.GetDepth(5, cons.LTC_BTC)
	t.Log(err)
	t.Log(dep.AskList)
	t.Log(dep.BidList)
}
func TestCoinEx_GetAccount(t *testing.T) {
	acc, err := coinex.GetAccount()
	t.Log(err)
	t.Log(acc)
}
func TestCoinEx_LimitBuy(t *testing.T) {
}
func TestCoinEx_LimitSell(t *testing.T) {
	ord, err := coinex.LimitSell("100", "0.0000601", cons.NewCurrencyPair2("CET_BCH"))
	t.Log(err)
	t.Log(ord)
}
func TestCoinEx_GetUnfinishOrders(t *testing.T) {
	ords, err := coinex.GetUnfinishedOrders(cons.NewCurrencyPair2("CET_BCH"))
	t.Log(err)
	if len(ords) > 0 {
		t.Log(fmt.Sprint(ords[0].OrderID))
	}
}
func TestCoinEx_CancelOrder(t *testing.T) {
	r, err := coinex.CancelOrder("37504128", cons.NewCurrencyPair2("CET_BCH"))
	t.Log(r, err)
}
func TestCoinEx_GetOneOrder(t *testing.T) {
	ord, err := coinex.GetOneOrder("37504128", cons.NewCurrencyPair2("CET_BCH"))
	t.Log(err)
	t.Log(ord)
}
