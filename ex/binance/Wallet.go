package binance

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/conbanwa/num"
	"github.com/conbanwa/slice"
	. "github.com/conbanwa/wstrader"
	. "github.com/conbanwa/wstrader/cons"
	. "github.com/conbanwa/wstrader/web"
	"net/url"
)

type Wallet struct {
	ba   *Binance
	conf *APIConfig
}

func NewWallet(c *APIConfig) *Wallet {
	return &Wallet{ba: NewWithConfig(c), conf: c}
}
func (w *Wallet) GetAccount() (*Account, error) {
	return nil, errors.New("not implement")
}
func (w *Wallet) Withdrawal(param WithdrawParameter) (withdrawId string, err error) {
	return "", errors.New("not implement")
}
func (w *Wallet) Transfer(param TransferParameter) error {
	transferUrl := w.conf.Endpoint + "/sapi/v1/futures/transfer"
	postParam := url.Values{}
	postParam.Set("asset", param.Currency)
	postParam.Set("amount", fmt.Sprint(param.Amount))
	if param.From == SPOT && param.To == SWAP_USDT {
		postParam.Set("type", "1")
	}
	if param.From == SWAP_USDT && param.To == SPOT {
		postParam.Set("type", "2")
	}
	if param.From == SPOT && param.To == FUTURE {
		postParam.Set("type", "3")
	}
	if param.From == FUTURE && param.To == SPOT {
		postParam.Set("type", "4")
	}
	w.ba.buildParamsSigned(&postParam)
	resp, err := HttpPostForm2(w.ba.httpClient, transferUrl, postParam,
		map[string]string{"X-MBX-APIKEY": w.ba.accessKey})
	if err != nil {
		return err
	}
	respMap := make(map[string]any)
	err = json.Unmarshal(resp, &respMap)
	if err != nil {
		return err
	}
	if respMap["tranId"] != nil && num.ToInt[int64](respMap["tranId"]) > 0 {
		return nil
	}
	return errors.New(slice.Bytes2String(resp))
}
func (w *Wallet) GetWithDrawHistory(currency *Currency) ([]DepositWithdrawHistory, error) {
	//historyUrl := w.conf.Endpoint + "/wapi/v3/withdrawHistory.html"
	historyUrl := w.conf.Endpoint + "/sapi/v1/accountSnapshot"
	postParam := url.Values{}
	postParam.Set("type", "SPOT")
	w.ba.buildParamsSigned(&postParam)
	resp, err := HttpGet5(w.ba.httpClient, historyUrl+"?"+postParam.Encode(),
		map[string]string{"X-MBX-APIKEY": w.ba.accessKey})
	if err != nil {
		return nil, err
	}
	log.Debug().Bytes("response data", resp).Send()
	respMap := make(map[string]any)
	err = json.Unmarshal(resp, &respMap)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (w *Wallet) GetDepositHistory(currency *Currency) ([]DepositWithdrawHistory, error) {
	historyUrl := w.conf.Endpoint + "/wapi/v3/depositHistory.html"
	postParam := url.Values{}
	postParam.Set("asset", currency.Symbol)
	w.ba.buildParamsSigned(&postParam)
	resp, err := HttpGet5(w.ba.httpClient, historyUrl+"?"+postParam.Encode(),
		map[string]string{"X-MBX-APIKEY": w.ba.accessKey})
	if err != nil {
		return nil, err
	}
	log.Debug().Bytes("response data", resp).Send()
	respMap := make(map[string]any)
	err = json.Unmarshal(resp, &respMap)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
