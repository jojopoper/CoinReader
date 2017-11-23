package main

import (
	"fmt"

	_y "github.com/jojopoper/CoinReader/okex"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read okex ...")
	rd := _y.Reader{}
	rd.Init("btc", "ltc", "127.0.0.1", "18181")
	// rd.Init("btc", "ltc")
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
