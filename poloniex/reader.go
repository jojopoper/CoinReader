package poloniex

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader Poloniex reader struct
type Reader struct {
	common.ReaderDef
	OrderDepth int
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
}

// Init init parameters
func (ths *Reader) Init(m, c string, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.baseInit()
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary = strings.ToUpper(ths.Monetary)
	ths.Coin = strings.ToUpper(ths.Coin)
	ths.Address = "poloniex.com"
	ths.OrderDepth = 20

	ths.OrderAddr = fmt.Sprintf("https://%s/public?command=returnOrderBook&depth=%d&currencyPair=%s_%s",
		ths.Address, ths.OrderDepth, ths.Monetary, ths.Coin)

	ths.HistoryAddr = fmt.Sprintf("https://%s/public?command=returnTradeHistory&currencyPair=%s_%s",
		ths.Address, ths.Monetary, ths.Coin)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Poloniex : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Poloniex : initHistParams set http client has error \n%+v", err)
	}
}
