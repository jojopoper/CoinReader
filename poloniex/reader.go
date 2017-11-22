package poloniex

import (
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
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
	ths.Monetary = strings.ToUpper(ths.Monetary)
	ths.Coin = strings.ToUpper(ths.Coin)
	ths.Address = "poloniex.com"
	ths.OrderDepth = 20
}
