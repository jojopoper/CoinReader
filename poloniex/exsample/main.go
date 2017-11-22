package main

import (
	"fmt"

	_p "github.com/jojopoper/CoinReader/poloniex"
	// _L "github.com/jojopoper/xlog"
)

func main() {
	// _L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read poloniex ...")
	rd := _p.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "xrp")
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
