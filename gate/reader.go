package gate

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
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
	ths.baseInit()
	ths.initOrderParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary = strings.ToLower(ths.Monetary)
	ths.Coin = strings.ToLower(ths.Coin)
	ths.Address = "gate.io"
	ths.isReading = false
	ths.r = rand.New(rand.NewSource(time.Now().Unix()))

	ths.OrderAddr = fmt.Sprintf("https://%s/json_svr/query/?u=11&type=ask_bid_list_table&symbol=%s_%s&c=",
		ths.Address, ths.Coin, ths.Monetary)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Gate : InitOrderParams set http client has error \n%+v", err)
	}
}
