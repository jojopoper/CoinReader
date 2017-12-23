package rd

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/CoinReader/common/tool"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader Okex reader struct
type Reader struct {
	rd.ReaderDef
	tool.ReverseSlice
	OrderDepth int
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
	ths.Address = "www.okex.com"
	ths.OrderDepth = 30
	ths.currentVer = "v1"

	ths.OrderAddr = fmt.Sprintf("https://%s/api/%s/depth.do?symbol=%s_%s&size=%d",
		ths.Address, ths.currentVer, ths.Coin.Name, ths.Monetary.Name, ths.OrderDepth)

	ths.HistoryAddr = fmt.Sprintf("https://%s/api/%s/trades.do?symbol=%s_%s",
		ths.Address, ths.currentVer, ths.Coin.Name, ths.Monetary.Name)
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
