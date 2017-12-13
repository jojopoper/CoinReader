package aex

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : aex reader define
type Reader struct {
	common.ReaderDef
	orderClt    *rhttp.CHttp
	historyClt  *rhttp.CHttp
	currentVer  string
	orderAddr   string
	historyAddr string
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
	ths.Address = "aex.com"
	ths.currentVer = "Update time 2017-02-17"
}

func (ths *Reader) initOrderParams() {
	ths.orderAddr = fmt.Sprintf("https://api.%s/depth.php?c=%s&mk_type=%s",
		ths.Address, ths.Coin, ths.Monetary)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	ths.setClient(ths.orderClt)
}

func (ths *Reader) initHistParams() {
	ths.historyAddr = fmt.Sprintf("https://api.%s/trades.php?c=%s&mk_type=%s",
		ths.Address, ths.Coin, ths.Monetary)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	ths.setClient(ths.historyClt)
}

func (ths *Reader) setClient(c *rhttp.CHttp) {
	if ths.Proxy.UseProxy() {
		client, err := c.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
		if err != nil {
			_L.Error("Aex : GetProxyClient(history) has error \n%+v", err)
			return
		}
		c.SetClient(client)
	} else {
		c.SetClient(c.GetClient(30))
	}
}
