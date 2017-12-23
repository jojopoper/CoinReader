package main

import (
	"fmt"

	"github.com/jojopoper/CoinReader/common/cap"
	_g "github.com/jojopoper/CoinReader/gate/rd"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read gate.io ...")

	btc := &cap.CoinCapacity{}
	btc.Set("btc", 8, 0.01)
	btg := &cap.CoinCapacity{}
	btg.Set("btg", 8, 0.01)

	rd := _g.Reader{}
	rd.Init(btc, btg, "127.0.0.1", "18181")
	// rd.Init(btc, btg)
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
