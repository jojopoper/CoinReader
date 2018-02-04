package main

import (
	"sync"
	"time"

	_r "github.com/jojopoper/CoinReader/aex/rd"
	"github.com/jojopoper/CoinReader/common/cap"
	_p "github.com/jojopoper/CoinReader/common/pfs"
	_cr "github.com/jojopoper/CoinReader/common/rd"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	_L.Info("Begin read aex ...")
	// thd1()
	// thd2()
	thd3()
}

func thd1() {
	for i := 0; i < 5; i++ {
		wt := new(sync.WaitGroup)
		wt.Add(3)
		xrpOdr, xrpHst := "", ""
		xlmOdr, xlmHst := "", ""
		btcOdr, btcHst := "", ""
		_L.Debug("[%d] Begin reading .... [%d]", i, time.Now().UnixNano())
		go reader1(wt, &xrpOdr, &xrpHst)
		go reader2(wt, &xlmOdr, &xlmHst)
		go reader3(wt, &btcOdr, &btcHst)
		wt.Wait()
		_L.Debug("[%d] End   reading .... [%d]", i, time.Now().UnixNano())
		_L.Info("%s", xrpOdr)
		_L.Info("%s", xrpHst)
		_L.Info("%s", xlmOdr)
		_L.Info("%s", xlmHst)
		_L.Info("%s", btcOdr)
		_L.Info("%s", btcHst)
		_L.Trace("Sleep 3 s...")
		time.Sleep(3 * time.Second)
	}
}

func thd2() {
	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	xrp := &cap.CoinCapacity{}
	xrp.Set("xrp", 7, 0.01)
	bitCny := &cap.CoinCapacity{}
	bitCny.Set("bitCNY", 2, 0.01)
	ltc := &cap.CoinCapacity{}
	ltc.Set("ltc", 5, 0.01)

	rd1 := _r.ReaderEx{}
	rd1.Init(btc, xrp, nil)
	rd2 := _r.ReaderEx{}
	rd2.Init(bitCny, btc, nil)
	rd3 := _r.ReaderEx{}
	rd3.Init(btc, ltc, nil)
	wt := new(sync.WaitGroup)
	var ts1, ts2 int64
	for i := 0; i < 10000; i++ {
		// xrpOdr, xrpHst := "", ""
		ts1 = time.Now().UnixNano()
		_L.Debug("[%d] Begin reading .... [%d]", i, ts1)
		// if rd.ReadAll() {
		// xrpOdr = rd.PrintOrders(5)
		// xrpHst = rd.PrintHistorys(5)
		// }
		wt.Add(3)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			rd1.ReadAll()
		}(wt)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			rd2.ReadAll()
		}(wt)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			rd3.ReadAll()
		}(wt)
		wt.Wait()
		ts2 = time.Now().UnixNano()
		_L.Debug("[%d] End   reading .... [%d]", i, ts2)
		_L.Info("%s", rd1.PrintOrders(3))
		_L.Info("%s", rd1.PrintHistorys(3))
		_L.Info("%s", rd2.PrintOrders(3))
		_L.Info("%s", rd2.PrintHistorys(3))
		_L.Info("%s", rd3.PrintOrders(3))
		_L.Info("%s", rd3.PrintHistorys(3))
		f := float64(ts2-ts1) / float64(1000000000.0)
		_L.Info("[%d] Timespan = %f", i, f)
		_L.Trace("Sleep 2 s...")
		time.Sleep(2 * time.Second)
	}
}

func thd3() {
	iCnt := 0

	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	bts := &cap.CoinCapacity{}
	bts.Set("bts", 7, 0.01)
	bitCny := &cap.CoinCapacity{}
	bitCny.Set("bitCNY", 2, 0.01)

	rd1 := _r.ReaderEx{}
	rd1.Init(btc, bts, nil)
	rd2 := _r.ReaderEx{}
	rd2.Init(bitCny, bts, rd1.GetClientCycle())
	rd3 := _r.ReaderEx{}
	rd3.Init(bitCny, btc, rd1.GetClientCycle())
	wt := new(sync.WaitGroup)
	var ts1, ts2 int64
	for i := 0; i < 100000; i++ {
		ts1 = time.Now().UnixNano()
		// _L.Debug("[%d] Begin reading .... [%d]", i, ts1)
		wt.Add(3)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			rd1.ReadOrders()
		}(wt)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			rd2.ReadOrders()
		}(wt)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			rd3.ReadOrders()
		}(wt)
		wt.Wait()
		ts2 = time.Now().UnixNano()
		// _L.Debug("[%d] End   reading .... [%d]", i, ts2)
		// _L.Info("%s", rd1.PrintOrders(3))
		// _L.Info("%s", rd1.PrintHistorys(3))
		// _L.Info("%s", rd2.PrintOrders(3))
		// _L.Info("%s", rd2.PrintHistorys(3))
		// _L.Info("%s", rd3.PrintOrders(3))
		// _L.Info("%s", rd3.PrintHistorys(3))
		f := float64(ts2-ts1) / float64(1000000000.0)
		_L.Info("[%d] Timespan = %f", i, f)

		pf := &_p.Profits3Element{}
		if len(rd1.Datas.Orders[_cr.OrderSellStringKey]) > 0 &&
			len(rd2.Datas.Orders[_cr.OrderBuyStringKey]) > 0 &&
			len(rd3.Datas.Orders[_cr.OrderSellStringKey]) > 0 {
			ret := pf.DcEcBcGain(pf.NewElement(rd1.Datas.Orders[_cr.OrderSellStringKey][0], 0.002, "BTS/BTC"),
				pf.NewElement(rd2.Datas.Orders[_cr.OrderBuyStringKey][0], 0.0, "BTS/CNY"),
				pf.NewElement(rd3.Datas.Orders[_cr.OrderSellStringKey][0], 0.0, "BTC/CNY"))
			if ret.Percent > 0.0 {
				iCnt++
				_L.Trace("BTC -> CNY")
				// _L.Debug("\nDcEcBc Flow = [%f , %f]\nDcEcBc gain = [%f%% , %f]", ret.Span, ret.Get, ret.Percent, ret.Profit)
				_L.Debug("%s", ret.PrintString())
			} else {
				_L.Trace("BTC -> CNY")
				_L.Trace("\n> DcEcBc Flow = [%f , %f]\n> DcEcBc gain = [%f%% , %f]", ret.Span, ret.Get, ret.Percent, ret.Profit)
			}
		}
		if len(rd1.Datas.Orders[_cr.OrderBuyStringKey]) > 0 &&
			len(rd2.Datas.Orders[_cr.OrderSellStringKey]) > 0 &&
			len(rd3.Datas.Orders[_cr.OrderBuyStringKey]) > 0 {
			ret := pf.BcEcDcGain(pf.NewElement(rd3.Datas.Orders[_cr.OrderBuyStringKey][0], 0.0, "CNY/BTC"),
				pf.NewElement(rd2.Datas.Orders[_cr.OrderSellStringKey][0], 0.0, "CNY/BTS"),
				pf.NewElement(rd1.Datas.Orders[_cr.OrderBuyStringKey][0], 0.002, "BTC/BTS"))
			if ret.Percent > 0.0 {
				iCnt++
				_L.Trace("CNY -> BTC")
				// _L.Debug("\nDcEcBc Flow = [%f , %f]\nDcEcBc gain = [%f%% , %f]", ret.Span, ret.Get, ret.Percent, ret.Profit)
				_L.Debug("%s", ret.PrintString())
			} else {
				_L.Trace("CNY -> BTC")
				_L.Trace("\n> BcEcDc Flow = [%f , %f]\n> BcEcDc gain = [%f%% , %f]", ret.Span, ret.Get, ret.Percent, ret.Profit)
			}
		}
		_L.Trace("[%d] Sleep 5 s...", iCnt)
		time.Sleep(5 * time.Second)
	}
}

func reader1(w *sync.WaitGroup, ord, hst *string) {
	if w != nil {
		defer w.Done()
	}
	rd := _r.Reader{}
	mt := &cap.CoinCapacity{}
	mt.Set("btc", 8, 0.01)
	c := &cap.CoinCapacity{}
	c.Set("xrp", 7, 0.01)
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init(mt, c)
	if rd.ReadAll() {
		*ord = rd.PrintOrders(3)
		*hst = rd.PrintHistorys(3)
	}
}

func reader2(w *sync.WaitGroup, ord, hst *string) {
	if w != nil {
		defer w.Done()
	}
	rd := _r.Reader{}
	mt := &cap.CoinCapacity{}
	mt.Set("btc", 8, 0.01)
	c := &cap.CoinCapacity{}
	c.Set("xlm", 7, 1)
	// rd.Init(mt, c, "127.0.0.1", "18181")
	rd.Init(mt, c)
	if rd.ReadAll() {
		*ord = rd.PrintOrders(3)
		*hst = rd.PrintHistorys(3)
	}
}

func reader3(w *sync.WaitGroup, ord, hst *string) {
	if w != nil {
		defer w.Done()
	}
	rd := _r.Reader{}
	mt := &cap.CoinCapacity{}
	mt.Set("bitCNY", 2, 0.1)
	c := &cap.CoinCapacity{}
	c.Set("btc", 8, 0.01)
	// rd.Init(mt, c, "127.0.0.1", "18181")
	rd.Init(mt, c)
	if rd.ReadAll() {
		*ord = rd.PrintOrders(3)
		*hst = rd.PrintHistorys(3)
	}
}
