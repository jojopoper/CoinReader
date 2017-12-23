package main

import (
	"fmt"

	"github.com/jojopoper/CoinReader/common/cap"
	_r "github.com/jojopoper/CoinReader/okex/rd"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read okex ...")

	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	ltc := &cap.CoinCapacity{}
	ltc.Set("ltc", 8, 0.01)

	rd := _r.Reader{}
	rd.Init(btc, ltc, "127.0.0.1", "18181")
	// rd.Init(btc,ltc)
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
