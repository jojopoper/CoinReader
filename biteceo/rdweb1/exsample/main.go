package main

import (
	"fmt"

	_r "github.com/jojopoper/CoinReader/biteceo/rdweb1"
	"github.com/jojopoper/CoinReader/common/cap"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	_L.Info("Begin read Bite.ceo ...")

	cny := &cap.CoinCapacity{}
	cny.Set("cny", 2, 0.01)
	eth := &cap.CoinCapacity{}
	eth.Set("eth", 8, 0.01)

	rd := _r.Reader{}
	// rd.Init(btc, btg, "127.0.0.1", "18181")
	rd.Init(cny, eth)
	if rd.ReadOrders() {
		ordstr := rd.PrintOrders(20)
		fmt.Println(ordstr)
	}
}
