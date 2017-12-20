package yobit

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader Yobit reader struct
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
	ths.baseInit()
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary = strings.ToLower(ths.Monetary)
	ths.Coin = strings.ToLower(ths.Coin)
	ths.Address = "yobit.net"
	ths.currentVer = "3"
	ths.depth = 20

	ths.OrderAddr = fmt.Sprintf("https://%s/api/%s/depth/%s_%s?limit=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.depth)

	ths.HistoryAddr = fmt.Sprintf("https://%s/api/%s/trades/%s_%s?limit=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.depth)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Yobit : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Yobit : initHistParams set http client has error \n%+v", err)
	}
}
