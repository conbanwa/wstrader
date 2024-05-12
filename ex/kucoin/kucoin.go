package kucoin

import (
	"github.com/conbanwa/num"
	. "qa3/wstrader"
	. "qa3/wstrader/cons"
	. "qa3/wstrader/q"
	"qa3/wstrader/util"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/conbanwa/logs"
)

func New(apiKey, apiSecret, apiPassphrase string) *KuCoin {
	return NewWithConfig(&APIConfig{
		Endpoint:      "https://api.kucoin.com",
		ApiKey:        apiKey,
		ApiSecretKey:  apiSecret,
		ApiPassphrase: apiPassphrase,
	})
}
func NewWithConfig(config *APIConfig) *KuCoin {
	if config.Endpoint == "" {
		config.Endpoint = "https://api.kucoin.com"
	}
	kc := &KuCoin{
		baseUrl:       config.Endpoint,
		apiKey:        config.ApiKey,
		apiSecret:     config.ApiSecretKey,
		apiPassphrase: config.ApiPassphrase,
	}
	kc.service = kucoin.NewApiService(
		kucoin.ApiBaseURIOption(kc.baseUrl),
		kucoin.ApiKeyOption(kc.apiKey),
		kucoin.ApiSecretOption(kc.apiSecret),
		kucoin.ApiPassPhraseOption(kc.apiPassphrase),
	)
	return kc
}

type KuCoin struct {
	apiKey        string
	apiSecret     string
	baseUrl       string
	apiPassphrase string
	service       *kucoin.ApiService
}

var inernalKlinePeriodConverter = map[KlinePeriod]string{
	KLINE_PERIOD_1MIN:  "1min",
	KLINE_PERIOD_3MIN:  "3min",
	KLINE_PERIOD_5MIN:  "5min",
	KLINE_PERIOD_15MIN: "15min",
	KLINE_PERIOD_30MIN: "30min",
	KLINE_PERIOD_60MIN: "1hour",
	KLINE_PERIOD_1H:    "1hour",
	KLINE_PERIOD_2H:    "2hour",
	KLINE_PERIOD_4H:    "4hour",
	KLINE_PERIOD_6H:    "6hour",
	KLINE_PERIOD_8H:    "8hour",
	KLINE_PERIOD_12H:   "12hour",
	KLINE_PERIOD_1DAY:  "1day",
	KLINE_PERIOD_1WEEK: "1week",
}

func (kc *KuCoin) String() string {
	return KUCOIN
}
func (kc *KuCoin) GetTicker(currency CurrencyPair) (*Ticker, error) {
	resp, err := kc.service.TickerLevel1(currency.ToSymbol("-"))
	if err != nil {
		logs.E("KuCoin GetTicker error:", err)
		return nil, err
	}
	var model kucoin.TickerLevel1Model
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetTicker error:", err)
		return nil, err
	}
	var ticker Ticker
	ticker.Pair = currency
	ticker.Date = uint64(model.Time / 1000)
	ticker.Last = num.ToFloat64(model.Price)
	ticker.Buy = num.ToFloat64(model.BestBid)
	ticker.Sell = num.ToFloat64(model.BestAsk)
	return &ticker, nil
}
func (kc *KuCoin) LimitBuy(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error) {
	clientID := util.GenerateOrderClientId(32)
	in := kucoin.CreateOrderModel{
		ClientOid: clientID,
		Side:      "buy",
		Symbol:    currency.ToSymbol("-"),
		Type:      "limit",
		Price:     price,
		Size:      amount,
	}
	resp, err := kc.service.CreateOrder(&in)
	if err != nil {
		logs.E("KuCoin LimitBuy error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin LimitBuy error:", err)
		return nil, err
	}
	var order Order
	order.OrderID2 = model.Id
	order.Cid = clientID
	return &order, nil
}
func (kc *KuCoin) LimitSell(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error) {
	clientID := util.GenerateOrderClientId(32)
	in := kucoin.CreateOrderModel{
		ClientOid: clientID,
		Side:      "sell",
		Symbol:    currency.ToSymbol("-"),
		Type:      "limit",
		Price:     price,
		Size:      amount,
	}
	resp, err := kc.service.CreateOrder(&in)
	if err != nil {
		logs.E("KuCoin LimitSell error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin LimitSell error:", err)
		return nil, err
	}
	var order Order
	order.OrderID2 = model.Id
	order.Cid = clientID
	return &order, nil
}
func (kc *KuCoin) MarketBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	clientID := util.GenerateOrderClientId(32)
	in := kucoin.CreateOrderModel{
		ClientOid: clientID,
		Side:      "buy",
		Symbol:    currency.ToSymbol("-"),
		Type:      "market",
		Price:     price,
		Size:      amount,
	}
	resp, err := kc.service.CreateOrder(&in)
	if err != nil {
		logs.E("KuCoin MarketBuy error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin MarketBuy error:", err)
		return nil, err
	}
	var order Order
	order.OrderID2 = model.Id
	order.Cid = clientID
	return &order, nil
}
func (kc *KuCoin) MarketSell(amount, price string, currency CurrencyPair) (*Order, error) {
	clientID := util.GenerateOrderClientId(32)
	in := kucoin.CreateOrderModel{
		ClientOid: clientID,
		Side:      "sell",
		Symbol:    currency.ToSymbol("-"),
		Type:      "market",
		Price:     price,
		Size:      amount,
	}
	resp, err := kc.service.CreateOrder(&in)
	if err != nil {
		logs.E("KuCoin MarketSell error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin MarketSell error:", err)
		return nil, err
	}
	var order Order
	order.OrderID2 = model.Id
	order.Cid = clientID
	return &order, nil
}
func (kc *KuCoin) CancelOrder(orderId string, currency CurrencyPair) (bool, error) {
	var resp *kucoin.ApiResponse
	var err error
	if orderId != "" {
		resp, err = kc.service.CancelOrder(orderId)
	} else {
		resp, err = kc.service.CancelOrder(currency.ToSymbol("-"))
	}
	if err != nil {
		logs.E("KuCoin CancelOrder error:", err)
		return false, err
	}
	var model kucoin.CancelOrderResultModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin CancelOrder error:", err)
		return false, err
	}
	return true, nil
}
func (kc *KuCoin) GetOneOrder(orderId string, currency CurrencyPair) (*Order, error) {
	resp, err := kc.service.Order(orderId)
	if err != nil {
		logs.E("KuCoin GetOneOrder error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetOneOrder error:", err)
		return nil, err
	}
	var order Order
	order.Price = num.ToFloat64(model.Price)
	order.Amount = num.ToFloat64(model.Size)
	order.AvgPrice = num.ToFloat64(model.DealFunds) / num.ToFloat64(model.DealSize)
	order.DealAmount = num.ToFloat64(model.DealSize)
	order.Fee = num.ToFloat64(model.Fee)
	order.Cid = model.ClientOid
	order.OrderID2 = model.Id
	order.OrderTime = int(model.CreatedAt / 1000)
	order.Type = model.Type
	order.Currency = NewCurrencyPair3(model.Symbol, "-")
	if model.Side == "buy" {
		if model.Type == "limit" {
			order.Side = BUY
		} else {
			order.Side = BUY_MARKET
		}
	} else {
		if model.Type == "limit" {
			order.Side = SELL
		} else {
			order.Side = SELL_MARKET
		}
	}
	if model.CancelExist {
		if model.DealSize == "0" {
			order.Status = ORDER_UNFINISH
		} else {
			order.Status = ORDER_PART_FINISH
		}
	} else {
		if model.DealSize == "0" {
			order.Status = ORDER_CANCEL
		} else if model.DealSize == model.Side {
			order.Status = ORDER_FINISH
		}
	}
	return &order, nil
}
func (kc *KuCoin) GetUnfinishedOrders(currency CurrencyPair) ([]Order, error) {
	params := map[string]string{
		"status": "active",
		"symbol": currency.ToSymbol("-"),
	}
	resp, err := kc.service.Orders(params, nil)
	if err != nil {
		logs.E("KuCoin GetUnfinishedOrders error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetUnfinishedOrders error:", err)
		return nil, err
	}
	var orders []Order
	return orders, nil
}
func (kc *KuCoin) GetOrderHistorys(currency CurrencyPair, optional ...OptionalParameter) ([]Order, error) {
	params := map[string]string{
		"status": "done",
		"symbol": currency.ToSymbol("-"),
	}
	pagination := kucoin.PaginationParam{}
	if len(optional) > 0 {
		pagination.CurrentPage = num.ToInt[int64](optional[0]["currentPage"])
		pagination.PageSize = num.ToInt[int64](optional[0]["pageSize"])
	}
	resp, err := kc.service.Orders(params, &pagination)
	if err != nil {
		logs.E("KuCoin GetOrderHistorys error:", err)
		return nil, err
	}
	var model kucoin.OrderModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetOrderHistorys error:", err)
		return nil, err
	}
	var orders []Order
	return orders, nil
}
func (kc *KuCoin) GetAccount() (*Account, error) {
	accs, err := kc.Accounts("", "")
	if err != nil {
		logs.E("KuCoin GetAccount error:", err)
		return nil, err
	}
	account := Account{}
	account.Exchange = kc.String()
	account.SubAccounts = make(map[Currency]SubAccount)
	for _, v := range accs {
		currency := NewCurrency(v.Currency, "").AdaptBccToBch()
		// KuCoin同一币种可能有多种账户类型
		if sub, exist := account.SubAccounts[currency]; !exist {
			account.SubAccounts[currency] = SubAccount{
				Currency:     currency,
				Amount:       num.ToFloat64(v.Available),
				ForzenAmount: num.ToFloat64(v.Holds),
			}
		} else {
			account.SubAccounts[currency] = SubAccount{
				Currency:     currency,
				Amount:       sub.Amount + num.ToFloat64(v.Available),
				ForzenAmount: sub.ForzenAmount + num.ToFloat64(v.Holds),
			}
		}
	}
	return &account, nil
}
func (kc *KuCoin) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	dep := 20
	if size > 20 {
		dep = 100
	}
	resp, err := kc.service.AggregatedPartOrderBook(currency.ToSymbol("-"), int64(dep))
	if err != nil {
		logs.E("KuCoin GetDepth error:", err)
		return nil, err
	}
	var model kucoin.PartOrderBookModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetDepth error:", err)
		return nil, err
	}
	var depth Depth
	depth.Pair = currency
	depth.UTime = time.Now()
	for i, ask := range model.Asks {
		if i < size {
			depth.AskList = append(depth.AskList, DepthRecord{
				Price:  num.ToFloat64(ask[0]),
				Amount: num.ToFloat64(ask[1]),
			})
		}
	}
	for j, bid := range model.Bids {
		if j < size {
			depth.BidList = append(depth.BidList, DepthRecord{
				Price:  num.ToFloat64(bid[0]),
				Amount: num.ToFloat64(bid[1]),
			})
		}
	}
	return &depth, nil
}
func (kc *KuCoin) GetKlineRecords(currency CurrencyPair, period KlinePeriod, size int, optional ...OptionalParameter) ([]Kline, error) {
	resp, err := kc.service.KLines(currency.ToSymbol("-"), inernalKlinePeriodConverter[period], 0, 0)
	if err != nil {
		logs.E("KuCoin GetKlineRecords error:", err)
		return nil, err
	}
	var model kucoin.KLinesModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetKlineRecords error:", err)
		return nil, err
	}
	var kLines []Kline
	for i, item := range model {
		if i < size {
			kLines = append(kLines, Kline{
				Pair:      currency,
				Timestamp: num.ToInt[int64]((*item)[0]),
				Open:      num.ToFloat64((*item)[1]),
				Close:     num.ToFloat64((*item)[2]),
				High:      num.ToFloat64((*item)[3]),
				Low:       num.ToFloat64((*item)[4]),
				Vol:       num.ToFloat64((*item)[6]),
			})
		}
	}
	return kLines, nil
}
func (kc *KuCoin) GetTrades(currency CurrencyPair, since int64) ([]Trade, error) {
	resp, err := kc.service.TradeHistories(currency.ToSymbol("-"))
	if err != nil {
		logs.E("KuCoin GetTrades error:", err)
		return nil, err
	}
	var model kucoin.TradeHistoriesModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin GetTrades error:", err)
		return nil, err
	}
	var trades []Trade
	for _, item := range model {
		typo := BUY
		if (*item).Side == "sell" {
			typo = SELL
		}
		trades = append(trades, Trade{
			Pair:   currency,
			Tid:    num.ToInt[int64]((*item).Sequence),
			Type:   typo,
			Amount: num.ToFloat64((*item).Size),
			Price:  num.ToFloat64((*item).Price),
			Date:   num.ToInt[int64]((*item).Time / 1000),
		})
	}
	return trades, nil
}

// Account
// Accounts returns a list of accounts.
func (kc *KuCoin) Accounts(currency, typo string) (kucoin.AccountsModel, error) {
	resp, err := kc.service.Accounts(currency, typo)
	if err != nil {
		logs.E("KuCoin Accounts error:", err)
		return nil, err
	}
	var model kucoin.AccountsModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin Accounts error:", err)
		return nil, err
	}
	return model, nil
}

// Account returns an account when you know the accountId.
func (kc *KuCoin) Account(accountId string) (*kucoin.AccountModel, error) {
	resp, err := kc.service.Account(accountId)
	if err != nil {
		logs.E("KuCoin Accounts error:", err)
		return nil, err
	}
	var model *kucoin.AccountModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin Accounts error:", err)
		return nil, err
	}
	return model, nil
}

// SubAccountUsers returns a list of sub-account user.
func (kc *KuCoin) SubAccountUsers() (kucoin.SubAccountUsersModel, error) {
	resp, err := kc.service.SubAccountUsers()
	if err != nil {
		logs.E("KuCoin SubAccountUsers error:", err)
		return nil, err
	}
	var model kucoin.SubAccountUsersModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin SubAccountUsers error:", err)
		return nil, err
	}
	return model, nil
}

// SubAccounts returns the aggregated balance of all sub-accounts of the current user.
func (kc *KuCoin) SubAccounts() (kucoin.SubAccountsModel, error) {
	resp, err := kc.service.SubAccounts()
	if err != nil {
		logs.E("KuCoin SubAccounts error:", err)
		return nil, err
	}
	var model kucoin.SubAccountsModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin SubAccounts error:", err)
		return nil, err
	}
	return model, nil
}

// SubAccount returns the detail of a sub-account.
func (kc *KuCoin) SubAccount(subUserId string) (*kucoin.SubAccountModel, error) {
	resp, err := kc.service.SubAccount(subUserId)
	if err != nil {
		logs.E("KuCoin SubAccount error:", err)
		return nil, err
	}
	var model *kucoin.SubAccountModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin SubAccount error:", err)
		return nil, err
	}
	return model, nil
}

// CreateAccount creates an account according to type(main|trade) and currency
// Parameter #1 typo is type of account.
func (kc *KuCoin) CreateAccount(typo, currency string) (*kucoin.AccountModel, error) {
	resp, err := kc.service.CreateAccount(typo, currency)
	if err != nil {
		logs.E("KuCoin CreateAccount error:", err)
		return nil, err
	}
	var model *kucoin.AccountModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin CreateAccount error:", err)
		return nil, err
	}
	return model, nil
}

// The inner transfer interface is used for transferring assets between the accounts of a user and is free of charges.
// For example, a user could transfer assets from their main account to their trading account on the platform.
func (kc *KuCoin) InnerTransfer(currency, from, to, amount string) (string, error) {
	resp, err := kc.service.InnerTransferV2(util.GenerateOrderClientId(32), currency, from, to, amount)
	if err != nil {
		logs.E("KuCoin InnerTransfer error:", err)
		return "", err
	}
	var model *kucoin.InnerTransferResultModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin InnerTransfer error:", err)
		return "", err
	}
	return model.OrderId, nil
}

// SubTransfer transfers between master account and sub-account.
func (kc *KuCoin) SubTransfer(currency, amount, direction, subUserId, accountType, subAccountType string) (string, error) {
	params := map[string]string{
		"clientOid":      util.GenerateOrderClientId(32),
		"currency":       currency,
		"amount":         amount,
		"direction":      direction,      // IN or OUT
		"subUserId":      subUserId,      // the user ID of a sub-account
		"accountType":    accountType,    // The account type of the master user: MAIN
		"subAccountType": subAccountType, //The account type of the sub user: MAIN, TRADE or MARGIN
	}
	resp, err := kc.service.SubTransfer(params)
	if err != nil {
		logs.E("KuCoin SubTransfer error:", err)
		return "", err
	}
	var model *kucoin.InnerTransferResultModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin SubTransfer error:", err)
		return "", err
	}
	return model.OrderId, nil
}

// Deposits
// CreateDepositAddress creates a deposit address.
func (kc *KuCoin) CreateDepositAddress(currency, chain string) (*kucoin.DepositAddressModel, error) {
	resp, err := kc.service.CreateDepositAddress(currency, chain)
	if err != nil {
		logs.E("KuCoin CreateDepositAddress error:", err)
		return nil, err
	}
	var model *kucoin.DepositAddressModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin CreateDepositAddress error:", err)
		return nil, err
	}
	return model, nil
}

// DepositAddresses returns the deposit address of currency for deposit.
// If return data is empty, you may need create a deposit address first.
func (kc *KuCoin) DepositAddresses(currency, chain string) (*kucoin.DepositAddressModel, error) {
	resp, err := kc.service.DepositAddresses(currency, chain)
	if err != nil {
		logs.E("KuCoin DepositAddresses error:", err)
		return nil, err
	}
	var model *kucoin.DepositAddressModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin DepositAddresses error:", err)
		return nil, err
	}
	return model, nil
}

// Deposits returns a list of deposit.
func (kc *KuCoin) Deposits(currency, startAt, endAt, status string) (*kucoin.DepositsModel, error) {
	params := map[string]string{
		"currency": currency,
		"startAt":  startAt,
		"endAt":    endAt,
		"status":   status,
	}
	resp, err := kc.service.Deposits(params, nil)
	if err != nil {
		logs.E("KuCoin Deposits error:", err)
		return nil, err
	}
	var model *kucoin.DepositsModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin Deposits error:", err)
		return nil, err
	}
	return model, nil
}

// Withdrawals
// Deposits returns a list of deposit.
func (kc *KuCoin) Withdrawals(currency, startAt, endAt, status string) (*kucoin.WithdrawalsModel, error) {
	params := map[string]string{
		"currency": currency,
		"startAt":  startAt,
		"endAt":    endAt,
		"status":   status,
	}
	resp, err := kc.service.Withdrawals(params, nil)
	if err != nil {
		logs.E("KuCoin Withdrawals error:", err)
		return nil, err
	}
	var model *kucoin.WithdrawalsModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin Withdrawals error:", err)
		return nil, err
	}
	return model, nil
}

// ApplyWithdrawal applies a withdrawal.
func (kc *KuCoin) ApplyWithdrawal(currency, address, amount, memo, isInner, remark, chain string) (string, error) {
	resp, err := kc.service.ApplyWithdrawal(currency, address, amount, map[string]string{
		"memo":    memo,
		"remark":  remark,
		"chain":   chain,
		"isInner": isInner,
	})
	if err != nil {
		logs.E("KuCoin ApplyWithdrawal error:", err)
		return "", err
	}
	var model *kucoin.ApplyWithdrawalResultModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin ApplyWithdrawal error:", err)
		return "", err
	}
	return model.WithdrawalId, nil
}

// WithdrawalQuotas returns the quotas of withdrawal.
func (kc *KuCoin) WithdrawalQuotas(currency, chain string) (*kucoin.WithdrawalQuotasModel, error) {
	resp, err := kc.service.WithdrawalQuotas(currency, chain)
	if err != nil {
		logs.E("KuCoin WithdrawalQuotas error:", err)
		return nil, err
	}
	var model *kucoin.WithdrawalQuotasModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin WithdrawalQuotas error:", err)
		return nil, err
	}
	return model, nil
}

// CancelWithdrawal cancels a withdrawal by withdrawalId.
func (kc *KuCoin) CancelWithdrawal(withdrawalId string) (*kucoin.CancelWithdrawalResultModel, error) {
	resp, err := kc.service.CancelWithdrawal(withdrawalId)
	if err != nil {
		logs.E("KuCoin CancelWithdrawal error:", err)
		return nil, err
	}
	var model *kucoin.CancelWithdrawalResultModel
	err = resp.ReadData(&model)
	if err != nil {
		logs.E("KuCoin CancelWithdrawal error:", err)
		return nil, err
	}
	return model, nil
}