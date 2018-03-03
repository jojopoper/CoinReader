package main

import (
	"fmt"

	_r "github.com/jojopoper/CoinReader/binance/rd"
	"github.com/jojopoper/CoinReader/common/cap"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	_L.Info("Begin read Binance ...")

	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	lsk := &cap.CoinCapacity{}
	lsk.Set("lsk", 8, 0.01)

	rd := _r.Reader{}
	// rd.Init(btc, btg, "127.0.0.1", "18181")
	rd.Init(btc, lsk)
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
