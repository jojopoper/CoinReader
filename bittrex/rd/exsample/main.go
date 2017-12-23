package main

import (
	"fmt"

	_r "github.com/jojopoper/CoinReader/bittrex/rd"
	"github.com/jojopoper/CoinReader/common/cap"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read bittrex ...")

	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	xrp := &cap.CoinCapacity{}
	xrp.Set("xrp", 8, 0.01)

	rd := _r.Reader{}
	// rd.Init(btc, xrp, "127.0.0.1", "18181")
	rd.Init(btc, xrp)
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
