package main

import (
	"time"

	_b "github.com/jojopoper/CoinReader/aex/bls"
	// _c "github.com/jojopoper/CoinReader/common"
	_L "github.com/jojopoper/xlog"
)

func main() {
	_L.LogInstance = _L.CreateDefaultLogInstance("ex.log", 5)
	_L.Info("Begin read aex ...")
	balance1()
}

func balance1() {
	b := &_b.Balances{}
	key := "api key"
	skey := "sky"
	b.Init("userid", []string{key, skey})
	for i := 0; i < 30; i++ {
		b.Update()
		_L.Info(" $$$$ XLM balance $$$$")
		_L.Info(" $ Unlock > %.8f", b.GetBalance("xlm"))
		_L.Info(" $ Locked > %.8f", b.GetLockBalance("xlm"))
		_L.Trace("waiting 2s...")
		time.Sleep(2 * time.Second)
	}
}
