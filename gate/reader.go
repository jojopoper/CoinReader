package gate

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
)

// Reader : gate website order and history datas reader
type Reader struct {
	sync.Mutex
	common.ReaderDef
	orderClt  *rhttp.CHttp
	isReading bool
	r         *rand.Rand
}

// Init init parameters
func (ths *Reader) Init(m, c string, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.reading, ths.reading, v...)
	ths.Monetary = strings.ToLower(ths.Monetary)
	ths.Coin = strings.ToLower(ths.Coin)
	ths.Address = "gate.io"
	ths.isReading = false
	ths.r = rand.New(rand.NewSource(time.Now().Unix()))
}
