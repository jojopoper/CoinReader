package main

import (
	"fmt"

	_b "github.com/jojopoper/CoinReader/bitfinex"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read bitfinex.com ...")
	rd := _b.Reader{}
	// rd.Init("btc", "eth", "127.0.0.1", "18181")
	rd.Init("btc", "eth")
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
