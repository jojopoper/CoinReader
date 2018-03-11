package main

import (
	"fmt"

	_r "github.com/jojopoper/CoinReader/cex/rd"
	"github.com/jojopoper/CoinReader/common/cap"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read cex.io ...")

	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	coin := &cap.CoinCapacity{}
	coin.Set("xlm", 8, 0.01)

	rd := _r.Reader{}
	rd.Init(btc, coin, "127.0.0.1", "8801")
	// rd.Init(btc,coin)
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
