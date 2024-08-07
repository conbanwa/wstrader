package exx

import (
	"errors"
	"github.com/conbanwa/wstrader/q"
	. "github.com/conbanwa/wstrader/web"
	"strings"
	"sync"
)

func (exx *Exx) PairArray() (map[string]q.D, map[q.D]q.P, error) {
	return nil, nil, nil
}
func Sym2duo(pair string) q.D {
	parts := strings.Split(pair, "_")
	if len(parts) != 2 {
		panic("FATAL: SPLIT ERR! " + pair)
	}
	if parts[0] == parts[1] {
		panic("FATAL: SAME CURRENCY ERR! " + pair)
	}
	return q.D{Base: unify(parts[0]), Quote: unify(parts[1])}
}
func (exx *Exx) Fee() float64 {
	return 0.001
}
func (exx *Exx) PlaceOrders(places [3]q.Order) (orders [3]q.Order, err error) {
	orders = places
	for _, p := range places {
		go func(p q.Order) {
		}(p)
	}
	return
}
func (exx *Exx) Balances() (availables, frozens *sync.Map, err error) {
	availables, frozens = new(sync.Map), new(sync.Map)
	return
}
func (exx *Exx) Test() bool {
	return true
}
func unify(local string) string {
	global := strings.ToUpper(local)
	return global
}
func (exx *Exx) GetAttr() (a q.Attr) {
	a.MultiOrder = false
	return a
}

func (exx *Exx) TradeFee() (map[string]q.TradeFee, error) {
	return nil, nil
}
func (exx *Exx) WithdrawFee() (sf []q.NetworkWithdraw, err error) {

	return
}
func (exx *Exx) OneTicker(d q.D) (ticker q.Bbo, err error) {
	return
}
func (exx *Exx) AllTicker(SymPair map[string]q.D) (mdt *sync.Map, err error) {
	url := "Exx.baseurl" + "wait_market_tickers"
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
