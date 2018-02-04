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

// Reader bcex.ca reader struct
type Reader struct {
	rd.ReaderDef
	tool.ReverseSlice
	orderClt      *rhttp.CHttp
	orderParams   string
	historyClt    *rhttp.CHttp
	historyParams string
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
	ths.Address = "www.bcex.ca"

	ths.OrderAddr = fmt.Sprintf("https://%s/Api_Order/depth", ths.Address)
	ths.orderParams = fmt.Sprintf("symbol=%s2%s", ths.Coin.Name, ths.Monetary.Name)

	ths.HistoryAddr = fmt.Sprintf("https://%s/Api_Order/marketOrder", ths.Address)
	ths.historyParams = fmt.Sprintf("symbol=%s2%s", ths.Coin.Name, ths.Monetary.Name)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Bcex.ca : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Bcex.ca : initHistParams set http client has error \n%+v", err)
	}
}
