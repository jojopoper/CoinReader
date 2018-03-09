package rd

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/CoinReader/common/tool"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : gate website order and history datas reader
type Reader struct {
	sync.Mutex
	rd.ReaderDef
	tool.ReverseSlice
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
	ths.Monetary.Name = strings.ToLower(ths.Monetary.Name)
	ths.Coin.Name = strings.ToLower(ths.Coin.Name)
	ths.Address = "gate.io"
	ths.currentVer = "api2"

	ths.OrderAddr = fmt.Sprintf("http://data.%s/%s/1/orderBook/%s_%s",
		ths.Address, ths.currentVer, ths.Coin.Name, ths.Monetary.Name)

	ths.HistoryAddr = fmt.Sprintf("http://data.%s/%s/1/tradeHistory/%s_%s",
		ths.Address, ths.currentVer, ths.Coin.Name, ths.Monetary.Name)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Gate : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Gate : initHistParams set http client has error \n%+v", err)
	}
}
