package rd

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : bitfinex reader define
type Reader struct {
	rd.ReaderDef
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
	limit      int
	currentVer string
}

// Init init parameters
func (ths *Reader) Init(m, c *cap.CoinCapacity, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.baseInit()
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary.Name = strings.ToLower(ths.Monetary.Name)
	ths.Coin.Name = strings.ToLower(ths.Coin.Name)
	ths.Address = "bitfinex.com"
	ths.limit = 100
	ths.currentVer = "v1"

	ths.OrderAddr = fmt.Sprintf("https://api.%s/%s/book/%s%s?limit_bid=%d&limit_asks=%d",
		ths.Address, ths.currentVer, ths.Coin.Name, ths.Monetary.Name, ths.limit, ths.limit)

	ths.HistoryAddr = fmt.Sprintf("https://api.%s/%s/trades/%s%s?limit_trades=%d",
		ths.Address, ths.currentVer, ths.Coin.Name, ths.Monetary.Name, ths.limit)
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
