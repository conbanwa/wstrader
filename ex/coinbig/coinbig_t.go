package coinbig

import (
	"errors"
	"qa3/wstrader/q"
	. "qa3/wstrader/web"
	"strings"
	"sync"

	"github.com/conbanwa/logs"
)

func (cb *CoinBig) PairArray() (map[string]q.D, map[q.D]q.P, error) {
	return nil, nil, nil
}
func Sym2duo(pair string) q.D {
	parts := strings.Split(pair, "_")
	var res q.D
	if len(parts) == 2 {
		res = q.D{Base: parts[0], Quote: parts[1]}
	} else {
		logs.F("FATAL: DIV ERR!", pair)
	}
	return res
}
func (cb *CoinBig) Fee() float64 {
	return 0.001
}
func (cb *CoinBig) PlaceOrders(places [3]q.Order) (orders [3]q.Order, err error) {
	orders = places
	for _, p := range places {
		go func(p q.Order) {
		}(p)
	}
	return
}
func (cb *CoinBig) Balances() (availables, frozens *sync.Map, err error) {
	availables, frozens = new(sync.Map), new(sync.Map)
	return
}
func (cb *CoinBig) Test() bool {
	return true
}
func unify(local string) string {
	global := strings.ToUpper(local)
	return global
}
func (cb *CoinBig) GetAttr() (a q.Attr) {
	a.MultiOrder = false
	return a
}

func (cb *CoinBig) TradeFee() (map[string]q.TradeFee, error) {
	return nil, nil
}
func (cb *CoinBig) WithdrawFee() (sf []q.NetworkWithdraw, err error) {

	return
}
func (cb *CoinBig) OneTicker(d q.D) (ticker q.Bbo, err error) {
	return
}
func (cb *CoinBig) AllTicker(SymPair map[string]q.D) (mdt *sync.Map, err error) {
	url := "CoinBig.baseurl" + "wait_market_tickers"
	respMap, err := HttpGet(nil, url)
	if err != nil {
		return
	}
	if respMap["status"].(string) == "error" {
		return mdt, errors.New(respMap["err-msg"].(string))
	}
	ticker := mdt
	//ticker.Pair = "currencyPair"
	return ticker, nil
}
