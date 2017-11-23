package bleutrade

import (
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
)

type Reader struct {
	common.ReaderDef
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
	currentVer string
	depth      int
}

// Init init parameters
func (ths *Reader) Init(m, c string, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.Monetary = strings.ToUpper(ths.Monetary)
	ths.Coin = strings.ToUpper(ths.Coin)
	ths.Address = "bleutrade.com"
	ths.currentVer = "v2"
	ths.depth = 50
}
