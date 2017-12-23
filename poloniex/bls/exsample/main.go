package main

import (
	"fmt"

	"github.com/jojopoper/CoinReader/poloniex/bls"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	fmt.Println("Begin read poloniex ...")
	thd2()
}

func thd2() {
	b := &bls.Balances{}
	key := "Api Key"
	skey := "Secret Key"
	userid := "xxxx" // 可以为空
	b.Init(userid, []string{key, skey}, "127.0.0.1", "18181")
	// b.Init(userid,[]string{key, skey})

	if b.Update() {
		_L.Info("     BTC Balance : %.8f", b.GetBalance("btc"))
		_L.Info("BTC Balance Lock : %.8f", b.GetLockBalance("btc"))
	}
}
