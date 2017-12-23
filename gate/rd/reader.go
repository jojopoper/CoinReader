package rd

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : gate website order and history datas reader
type Reader struct {
	sync.Mutex
	rd.ReaderDef
	orderClt  *rhttp.CHttp
	isReading bool
	r         *rand.Rand
}

// Init init parameters
func (ths *Reader) Init(m, c *cap.CoinCapacity, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.reading, ths.reading, v...)
	ths.baseInit()
	ths.initOrderParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary.Name = strings.ToLower(ths.Monetary.Name)
	ths.Coin.Name = strings.ToLower(ths.Coin.Name)
	ths.Address = "gate.io"
	ths.isReading = false
	ths.r = rand.New(rand.NewSource(time.Now().Unix()))

	ths.OrderAddr = fmt.Sprintf("https://%s/json_svr/query/?u=11&type=ask_bid_list_table&symbol=%s_%s&c=",
		ths.Address, ths.Coin.Name, ths.Monetary.Name)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Gate : InitOrderParams set http client has error \n%+v", err)
	}
}
