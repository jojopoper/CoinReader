package bitfinex

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : bitfinex reader define
type Reader struct {
	common.ReaderDef
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
	limit      int
	currentVer string
}

// Init init parameters
func (ths *Reader) Init(m, c string, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.baseInit()
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary = strings.ToLower(ths.Monetary)
	ths.Coin = strings.ToLower(ths.Coin)
	ths.Address = "bitfinex.com"
	ths.limit = 100
	ths.currentVer = "v1"

	ths.OrderAddr = fmt.Sprintf("https://api.%s/%s/book/%s%s?limit_bid=%d&limit_asks=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.limit, ths.limit)

	ths.HistoryAddr = fmt.Sprintf("https://api.%s/%s/trades/%s%s?limit_trades=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.limit)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Bitfinex : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Bitfinex : initHistParams set http client has error \n%+v", err)
	}
}
