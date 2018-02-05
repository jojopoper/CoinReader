package rd1

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/CoinReader/common/tool"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Reader : bite.ceo reader define
type Reader struct {
	sync.Mutex
	rd.ReaderDef
	tool.ReverseSlice
	orderClt    *rhttp.CHttp
	orderParams string
	isReading   bool
}

// Init init parameters
func (ths *Reader) Init(m, c *cap.CoinCapacity, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdOrders, v...)
	ths.baseInit()
	ths.initOrderParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary.Name = strings.ToLower(ths.Monetary.Name)
	ths.Coin.Name = strings.ToLower(ths.Coin.Name)
	ths.Address = "www.bite.ceo"
	ths.isReading = false

	ths.OrderAddr = fmt.Sprintf("https://%s/trade/index_json?t=0.", ths.Address)
	ths.orderParams = fmt.Sprintf("market=%s_%s", ths.Coin.Name, ths.Monetary.Name)
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Bite.ceo : InitOrderParams set http client has error \n%+v", err)
	}
}
