package rd

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

type Reader struct {
	rd.ReaderDef
	orderClt   *rhttp.CHttp
	historyClt *rhttp.CHttp
	Size       int
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
	ths.Address = "huobi.pro"
	ths.Size = 100

	ths.OrderAddr = fmt.Sprintf("https://api.%s/market/depth?symbol=%s%s&type=step0",
		ths.Address, ths.Coin.Name, ths.Monetary.Name)

	ths.HistoryAddr = fmt.Sprintf("https://api.%s/market/history/trade?symbol=%s%s&size=%d",
		ths.Address, ths.Coin.Name, ths.Monetary.Name, ths.Size)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("huobi.pro : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("huobi.pro : initHistParams set http client has error \n%+v", err)
	}
}
