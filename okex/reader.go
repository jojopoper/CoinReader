package okex

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader Okex reader struct
type Reader struct {
	common.ReaderDef
	OrderDepth int
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
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
	ths.Address = "www.okex.com"
	ths.OrderDepth = 30
	ths.currentVer = "v1"

	ths.OrderAddr = fmt.Sprintf("https://%s/api/%s/depth.do?symbol=%s_%s&size=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.OrderDepth)

	ths.HistoryAddr = fmt.Sprintf("https://%s/api/%s/trades.do?symbol=%s_%s",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Okex : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Okex : initHistParams set http client has error \n%+v", err)
	}
}
