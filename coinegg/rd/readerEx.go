package rd

import (
	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/nt"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// ReaderEx : coinegg reader define
type ReaderEx struct {
	Reader
	cltCle    *nt.ClientCycle
	cntClient int
}

// Init init parameters
func (ths *ReaderEx) Init(m, c *cap.CoinCapacity, clt *nt.ClientCycle, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.baseInit()
	ths.cntClient = 10
	if clt == nil {
		ths.initClient()
	} else {
		ths.cltCle = clt
	}
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *ReaderEx) initClient() {
	ths.cltCle = new(nt.ClientCycle)
	if ths.Proxy.UseProxy() {
		err := ths.cltCle.MakeProxy(ths.cntClient, ths.Proxy.Address, ths.Proxy.Port)
		if err != nil {
			_L.Error("Coinegg : initClient has error \n%+v", err)
		}
	} else {
		ths.cltCle.Make(ths.cntClient)
	}
}

// GetClientCycle : get current client cycle object
func (ths *ReaderEx) GetClientCycle() *nt.ClientCycle {
	return ths.cltCle
}

func (ths *ReaderEx) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
}

func (ths *ReaderEx) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
}
