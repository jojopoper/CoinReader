package main

import (
	"fmt"
	"sync"
	"time"

	_b "github.com/jojopoper/CoinReader/aex"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read aex ...")
	thd1()
	// thd2()
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
	rd1 := _b.ReaderEx{}
	rd1.Init("btc", "xrp")
	rd2 := _b.ReaderEx{}
	rd2.Init("bitcny", "btc")
	rd3 := _b.ReaderEx{}
	rd3.Init("btc", "ltc")
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

func reader1(w *sync.WaitGroup, ord, hst *string) {
	if w != nil {
		defer w.Done()
	}
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "xrp")
	if rd.ReadAll() {
		*ord = rd.PrintOrders(5)
		*hst = rd.PrintHistorys(5)
	}
}

func reader2(w *sync.WaitGroup, ord, hst *string) {
	if w != nil {
		defer w.Done()
	}
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "xlm")
	if rd.ReadAll() {
		*ord = rd.PrintOrders(5)
		*hst = rd.PrintHistorys(5)
	}
}

func reader3(w *sync.WaitGroup, ord, hst *string) {
	if w != nil {
		defer w.Done()
	}
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("bitcny", "btc")
	if rd.ReadAll() {
		*ord = rd.PrintOrders(5)
		*hst = rd.PrintHistorys(5)
	}
}
