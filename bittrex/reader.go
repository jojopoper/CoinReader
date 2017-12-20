package bittrex

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : bittrex reader define
type Reader struct {
	common.ReaderDef
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
	ths.Monetary = strings.ToUpper(ths.Monetary)
	ths.Coin = strings.ToUpper(ths.Coin)
	ths.Address = "bittrex.com"
	ths.currentVer = "v1.1"

	ths.OrderAddr = fmt.Sprintf("https://%s/api/%s/public/getorderbook?market=%s-%s&type=both",
		ths.Address, ths.currentVer, ths.Monetary, ths.Coin)

	ths.HistoryAddr = fmt.Sprintf("https://%s/api/%s/public/getmarkethistory?market=%s-%s",
		ths.Address, ths.currentVer, ths.Monetary, ths.Coin)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Bittrex : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Bittrex : initHistParams set http client has error \n%+v", err)
	}
}
