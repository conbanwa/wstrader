package build

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"qa3/wstrader"
	"qa3/wstrader/config"
	"qa3/wstrader/cons"
	"qa3/wstrader/ex/atop"
	"qa3/wstrader/ex/bigone"
	"qa3/wstrader/ex/binance"
	"qa3/wstrader/ex/bitfinex"
	"qa3/wstrader/ex/bithumb"
	"qa3/wstrader/ex/bitmex"
	"qa3/wstrader/ex/bitstamp"
	"qa3/wstrader/ex/bittrex"
	"qa3/wstrader/ex/coinbene"
	"qa3/wstrader/ex/coinex"
	"qa3/wstrader/ex/ftx"
	"qa3/wstrader/ex/gateio"
	"qa3/wstrader/ex/gdax"
	"qa3/wstrader/ex/hitbtc"
	"qa3/wstrader/ex/huobi"
	"qa3/wstrader/ex/kraken"
	"qa3/wstrader/ex/kucoin"
	"qa3/wstrader/ex/okx"
	"qa3/wstrader/ex/okx/okex"
	"qa3/wstrader/ex/poloniex"
	"qa3/wstrader/ex/zb"
	"time"

	"github.com/conbanwa/logs"
)

type APIBuilder struct {
	HttpClientConfig *config.HttpClientConfig
	client           *http.Client
	httpTimeout      time.Duration
	apiKey           string
	secretKey        string
	clientId         string
	apiPassphrase    string
	futuresEndPoint  string
	endPoint         string
	futuresLever     float64
}

var DefaultAPIBuilder = NewAPIBuilder()

func init() {
	config.SetProxy()
}

func NewAPIBuilder() (builder *APIBuilder) {
	return &APIBuilder{
		HttpClientConfig: config.DefaultHttpClientConfig,
		client: &http.Client{
			Timeout: config.DefaultHttpClientConfig.HttpTimeout,
			Transport: &http.Transport{
				Proxy: func(request *http.Request) (*url.URL, error) {
					return config.DefaultHttpClientConfig.Proxy, nil
				},
				MaxIdleConns:          config.DefaultHttpClientConfig.MaxIdleConns,
				IdleConnTimeout:       5 * config.DefaultHttpClientConfig.HttpTimeout,
				MaxConnsPerHost:       2,
				MaxIdleConnsPerHost:   2,
				TLSHandshakeTimeout:   config.DefaultHttpClientConfig.HttpTimeout,
				ResponseHeaderTimeout: config.DefaultHttpClientConfig.HttpTimeout,
				ExpectContinueTimeout: config.DefaultHttpClientConfig.HttpTimeout,
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, config.DefaultHttpClientConfig.HttpTimeout)
				}},
		}}
}

func NewCustomAPIBuilder(client *http.Client) (builder *APIBuilder) {
	return &APIBuilder{client: client}
}

func (builder *APIBuilder) BuildSpotWs(exName string) (wstrader.SpotWsApi, error) {
	switch exName {
	case cons.OKEX_V3, cons.OKEX:
		return okex.NewOKExSpotV3Ws(nil), nil
	case cons.FTX:
		return ftx.NewWs(), nil
	case cons.HUOBI_PRO, cons.HUOBI:
		return huobi.NewSpotWs(), nil
	case cons.BINANCE:
		return binance.NewSpotWs(), nil
	case cons.BITFINEX:
		return bitfinex.NewWs(), nil
	case cons.GATEIO:
		return gateio.NewWs(), nil
	}
	return nil, errors.New("not support the exchange " + exName)
}

func (builder *APIBuilder) GetHttpClientConfig() *config.HttpClientConfig {
	return builder.HttpClientConfig
}
func (builder *APIBuilder) GetHttpClient() *http.Client {
	return builder.client
}
func (builder *APIBuilder) HttpProxy(proxyUrl string) (_builder *APIBuilder) {
	if proxyUrl == "" {
		return builder
	}
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return builder
	}
	builder.HttpClientConfig.Proxy = proxy
	transport := builder.client.Transport.(*http.Transport)
	transport.Proxy = http.ProxyURL(proxy)
	return builder
}
func (builder *APIBuilder) HttpTimeout(timeout time.Duration) (_builder *APIBuilder) {
	builder.HttpClientConfig.HttpTimeout = timeout
	builder.httpTimeout = timeout
	builder.client.Timeout = timeout
	if transport := builder.client.Transport.(*http.Transport); transport != nil {
		//transport.ResponseHeaderTimeout = timeout
		//transport.TLSHandshakeTimeout = timeout
		transport.IdleConnTimeout = timeout
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		}
	}
	return builder
}
func (builder *APIBuilder) APIKey(key string) (_builder *APIBuilder) {
	builder.apiKey = key
	return builder
}
func (builder *APIBuilder) APISecretkey(key string) (_builder *APIBuilder) {
	builder.secretKey = key
	return builder
}
func (builder *APIBuilder) ClientID(id string) (_builder *APIBuilder) {
	builder.clientId = id
	return builder
}
func (builder *APIBuilder) ApiPassphrase(apiPassphrase string) (_builder *APIBuilder) {
	builder.apiPassphrase = apiPassphrase
	return builder
}
func (builder *APIBuilder) FuturesEndpoint(endpoint string) (_builder *APIBuilder) {
	builder.futuresEndPoint = endpoint
	return builder
}
func (builder *APIBuilder) Endpoint(endpoint string) (_builder *APIBuilder) {
	builder.endPoint = endpoint
	return builder
}
func (builder *APIBuilder) FuturesLever(lever float64) (_builder *APIBuilder) {
	builder.futuresLever = lever
	return builder
}
func (builder *APIBuilder) Build(exName string) (api wstrader.API) {
	var _api wstrader.API
	switch exName {
	case cons.KUCOIN:
		_api = kucoin.New(builder.apiKey, builder.secretKey, builder.apiPassphrase)
	//case OKCOIN_CN:
	//	_api = okcoin.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.POLONIEX:
		_api = poloniex.New(builder.client, builder.apiKey, builder.secretKey)
	//case OKCOIN_COM:
	//	_api = okcoin.NewCOM(builder.client, builder.apiKey, builder.secretKey)
	case cons.BITSTAMP:
		_api = bitstamp.NewBitstamp(builder.client, builder.apiKey, builder.secretKey, builder.clientId)
	case cons.HUOBI_PRO:
		//_api = huobi.NewHuoBiProSpot(builder.client, builder.apiKey, builder.secretKey)
		_api = huobi.NewHuobiWithConfig(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.endPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey})
	case cons.OKEX_V3:
		_api = okex.NewOKEx(&wstrader.APIConfig{
			HttpClient:    builder.client,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretKey,
			ApiPassphrase: builder.apiPassphrase,
			Endpoint:      builder.endPoint,
		})
	case cons.OKEX:
		_api = okx.NewOKExV5Spot(&wstrader.APIConfig{
			HttpClient:    builder.client,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretKey,
			ApiPassphrase: builder.apiPassphrase,
			Endpoint:      builder.endPoint,
		})
	case cons.BITFINEX:
		_api = bitfinex.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.KRAKEN:
		_api = kraken.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.BINANCE:
		//_api = binance.New(builder.client, builder.apiKey, builder.secretKey)
		_api = binance.NewWithConfig(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.endPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey})
	case cons.BITTREX:
		_api = bittrex.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.BITHUMB:
		_api = bithumb.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.GDAX:
		_api = gdax.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.ZB:
		_api = zb.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.COINEX:
		_api = coinex.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.BIGONE:
		_api = bigone.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.HITBTC:
		_api = hitbtc.New(builder.client, builder.apiKey, builder.secretKey)
	case cons.ATOP:
		_api = atop.New(builder.client, builder.apiKey, builder.secretKey)
	default:
		logs.W("exchange name error [" + exName + "].")
	}
	return _api
}
func (builder *APIBuilder) BuildFuture(exName string) (api wstrader.FutureRestAPI) {
	switch exName {
	case cons.BITMEX:
		return bitmex.New(&wstrader.APIConfig{
			//Endpoint:     "https://www.bitmex.com/",
			Endpoint:     builder.futuresEndPoint,
			HttpClient:   builder.client,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey})
	case cons.BITMEX_TEST:
		return bitmex.New(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     "https://testnet.bitmex.com",
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
		})
	case cons.OKEX_FUTURE, cons.OKEX_V3:
		//return okcoin.NewOKEx(builder.client, builder.apiKey, builder.secretKey)
		return okex.NewOKEx(&wstrader.APIConfig{
			HttpClient: builder.client,
			//	Endpoint:      "https://www.okx.com",
			Endpoint:      builder.futuresEndPoint,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretKey,
			ApiPassphrase: builder.apiPassphrase,
			Lever:         builder.futuresLever}).OKExFuture
	case cons.HBDM:
		return huobi.NewHbdm(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.futuresEndPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
			Lever:        builder.futuresLever})
	case cons.HBDM_SWAP:
		return huobi.NewHbdmSwap(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.endPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
			Lever:        builder.futuresLever,
		})
	case cons.OKEX_SWAP:
		return okex.NewOKEx(&wstrader.APIConfig{
			HttpClient:    builder.client,
			Endpoint:      builder.futuresEndPoint,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretKey,
			ApiPassphrase: builder.apiPassphrase,
			Lever:         builder.futuresLever}).OKExSwap
	case cons.COINBENE:
		return coinbene.NewCoinbeneSwap(wstrader.APIConfig{
			HttpClient: builder.client,
			//	Endpoint:     "http://openapi-contract.coinbene.com",
			Endpoint:     builder.futuresEndPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
			Lever:        builder.futuresLever,
		})
	case cons.BINANCE_SWAP:
		return binance.NewBinanceSwap(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.futuresEndPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
			Lever:        builder.futuresLever,
		})
	case cons.BINANCE, cons.BINANCE_FUTURES:
		return binance.NewBinanceFutures(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.futuresEndPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
			Lever:        builder.futuresLever,
		})
	default:
		println(fmt.Sprintf("%s not support future", exName))
		return nil
	}
}
func (builder *APIBuilder) BuildFuturesWs(exName string) (wstrader.FuturesWsApi, error) {
	switch exName {
	case cons.OKEX_V3, cons.OKEX, cons.OKEX_FUTURE:
		return okex.NewOKExV3FuturesWs(okex.NewOKEx(&wstrader.APIConfig{
			HttpClient: builder.client,
			Endpoint:   builder.futuresEndPoint,
		})), nil
	case cons.HBDM:
		return huobi.NewHbdmWs(), nil
	case cons.HBDM_SWAP:
		return huobi.NewHbdmSwapWs(), nil
	case cons.BINANCE, cons.BINANCE_FUTURES, cons.BINANCE_SWAP:
		return binance.NewFuturesWs(), nil
	case cons.BITMEX:
		return bitmex.NewSwapWs(), nil
	}
	return nil, errors.New("not support the exchange " + exName)
}
func (builder *APIBuilder) BuildWallet(exName string) (wstrader.WalletApi, error) {
	switch exName {
	case cons.OKEX_V3, cons.OKEX:
		return okex.NewOKEx(&wstrader.APIConfig{
			HttpClient:    builder.client,
			ApiKey:        builder.apiKey,
			ApiSecretKey:  builder.secretKey,
			ApiPassphrase: builder.apiPassphrase,
		}).OKExWallet, nil
	case cons.HUOBI_PRO:
		return huobi.NewWallet(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.endPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
		}), nil
	case cons.BINANCE:
		return binance.NewWallet(&wstrader.APIConfig{
			HttpClient:   builder.client,
			Endpoint:     builder.endPoint,
			ApiKey:       builder.apiKey,
			ApiSecretKey: builder.secretKey,
		}), nil
	}
	return nil, errors.New("not support the wallet api for  " + exName)
}