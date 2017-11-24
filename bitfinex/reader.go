package bitfinex

import (
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
)

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
	ths.Monetary = strings.ToLower(ths.Monetary)
	ths.Coin = strings.ToLower(ths.Coin)
	ths.Address = "bitfinex.com"
	ths.limit = 100
	ths.currentVer = "v1"
}
