package main

import (
	"fmt"

	_b "github.com/jojopoper/CoinReader/bittrex"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read aex ...")
	rd := _b.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "xrp")
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
