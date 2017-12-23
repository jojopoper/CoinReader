package main

import (
	"time"

	_b "github.com/jojopoper/CoinReader/aex/trade"
	// _c "github.com/jojopoper/CoinReader/common"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	_L.Info("Begin read aex ...")
	key := "api key"
	skey := "sky"
	uid := "userid"
	// trade1(key,skey,uid)
	// trade2(key,skey,uid)
	// trade3(key, skey, uid)
	trade4(key, skey, uid)
}

func trade1(k, sk, id string) {
	td := &_b.TradeOrder{}
	td.Init(id, []string{k, sk})
	o := &_b.AexOrderBase{}
	o.Init("bitCNY", "XLM", "")
	o.SetDeci(2, 6)
	oi := &_b.AexOrderInfo{}
	oi.SetOrderBase(o)

	tid := td.SubmitOrder(oi.GetSellInfo(), 1.9, 100.0)
	_L.Debug("%s", tid)
}

func trade2(k, sk, id string) {
	td := &_b.TradeOrder{}
	td.Init(id, []string{k, sk})
	o := &_b.AexOrderCancel{}
	o.Init("bitCNY", "XLM", "1787")
	// o.SetDeci(2, 6)

	tid := td.CancelOrder(o)
	_L.Debug("%s", tid)
}

func trade3(k, sk, id string) {
	td := &_b.TradeOrder{}
	td.Init(id, []string{k, sk})
	o := &_b.AexOrderBase{}
	o.Init("bitCNY", "XLM", "")
	o.SetDeci(2, 6)
	oi := &_b.AexOrderInfo{}
	oi.SetOrderBase(o)
	ret := td.SubmitOrder(oi.GetSellInfo(), 1.9, 100.0)
	_L.Trace("Submit order :%+v", ret)
	time.Sleep(2 * time.Second)
	if ret.IsSuccess && ret.ID != "" {
		o := &_b.AexOrderCancel{}
		o.Init("bitCNY", "XLM", ret.ID)
		// o.SetDeci(2, 6)

		ret = td.CancelOrder(o)
		_L.Debug("cancel order : %+v", ret)
	}
}

func trade4(k, sk, id string) {
	td := &_b.TradeOrder{}
	td.Init(id, []string{k, sk})
	ret := td.OrderList("bitCNY", "XLM")
	if ret != nil {
		_L.Debug("%+v", ret)
		for _, itm := range ret {
			o := &_b.AexOrderCancel{}
			o.Init("bitCNY", "XLM", itm.TradeID)

			cret := td.CancelOrder(o)
			_L.Debug("cancel order : %+v", cret)
			time.Sleep(2 * time.Second)
		}
	}
}
