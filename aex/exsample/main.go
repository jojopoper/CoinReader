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
	wt := new(sync.WaitGroup)
	wt.Add(3)
	xrpOdr, xrpHst := "", ""
	xlmOdr, xlmHst := "", ""
	btcOdr, btcHst := "", ""
	_L.Debug("Begin reading .... [%d]", time.Now().Unix())
	go reader1(wt, &xrpOdr, &xrpHst)
	go reader2(wt, &xlmOdr, &xlmHst)
	go reader3(wt, &btcOdr, &btcHst)
	wt.Wait()
	_L.Debug("End reading ....   [%d]", time.Now().Unix())
	_L.Info("%s", xrpOdr)
	_L.Info("%s", xrpHst)
	_L.Info("%s", xlmOdr)
	_L.Info("%s", xlmHst)
	_L.Info("%s", btcOdr)
	_L.Info("%s", btcHst)
}

func reader1(w *sync.WaitGroup, ord, hst *string) {
	defer w.Done()
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "xrp")
	if rd.ReadAll() {
		*ord = rd.PrintOrders(20)
		*hst = rd.PrintHistorys(20)
	}
}

func reader2(w *sync.WaitGroup, ord, hst *string) {
	defer w.Done()
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "xlm")
	if rd.ReadAll() {
		*ord = rd.PrintOrders(20)
		*hst = rd.PrintHistorys(20)
	}
}

func reader3(w *sync.WaitGroup, ord, hst *string) {
	defer w.Done()
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("bitcny", "btc")
	if rd.ReadAll() {
		*ord = rd.PrintOrders(20)
		*hst = rd.PrintHistorys(20)
	}
}
