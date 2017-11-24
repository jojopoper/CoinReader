package main

import (
	"fmt"

	_h "github.com/jojopoper/CoinReader/huobi"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read huobi.pro ...")
	rd := _h.Reader{}
	rd.Init("btc", "bt1", "127.0.0.1", "18181")
	// rd.Init("btc", "eth")
	if rd.ReadAll() {
		ordstr := rd.PrintOrders(20)
		hisstr := rd.PrintHistorys(20)
		fmt.Println(ordstr)
		fmt.Println(hisstr)
	}
}
