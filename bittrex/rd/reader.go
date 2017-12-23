package rd

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : bittrex reader define
type Reader struct {
	rd.ReaderDef
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
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
	ths.Monetary.Name = strings.ToUpper(ths.Monetary.Name)
	ths.Coin.Name = strings.ToUpper(ths.Coin.Name)
	ths.Address = "bittrex.com"
	ths.currentVer = "v1.1"

	ths.OrderAddr = fmt.Sprintf("https://%s/api/%s/public/getorderbook?market=%s-%s&type=both",
		ths.Address, ths.currentVer, ths.Monetary.Name, ths.Coin.Name)

	ths.HistoryAddr = fmt.Sprintf("https://%s/api/%s/public/getmarkethistory?market=%s-%s",
		ths.Address, ths.currentVer, ths.Monetary.Name, ths.Coin.Name)
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
