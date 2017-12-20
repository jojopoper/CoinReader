package aex

import (
	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// ReaderEx : aex reader define
type ReaderEx struct {
	Reader
	common.ClientCycle
	cntClient int
}

// Init init parameters
func (ths *ReaderEx) Init(m, c string, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.baseInit()
	ths.cntClient = 10
	ths.initClient()
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *ReaderEx) initClient() {
	if ths.Proxy.UseProxy() {
		err := ths.MakeProxy(ths.cntClient, ths.Proxy.Address, ths.Proxy.Port)
		if err != nil {
			_L.Error("Aex : initClient has error \n%+v", err)
		}
	} else {
		ths.Make(ths.cntClient)
	}
}

func (ths *ReaderEx) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
}

func (ths *ReaderEx) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
}
