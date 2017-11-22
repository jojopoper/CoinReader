package main

import (
	"fmt"

	_y "github.com/jojopoper/CoinReader/yobit"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read yobit ...")
	rd := _y.Reader{}
	// rd.Init("btc", "xrp", "127.0.0.1", "18181")
	rd.Init("btc", "btg")
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
