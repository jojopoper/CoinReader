package rd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) endReading() {
	ths.Lock()
	defer ths.Unlock()
	ths.isReading = false
}

// rdOrders : readout order datas from bite.ceo, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ths.Lock()
	if ths.isReading {
		defer ths.Unlock()
		return true
	}
	ths.isReading = true
	ths.Unlock()
	defer ths.endReading()

	addr := fmt.Sprintf("%s%d", ths.OrderAddr, time.Now().UnixNano())
	ret, err := ths.orderClt.ClientGet(addr, rhttp.ReturnCustomType)
	// ret, err := ths.orderClt.Get(addr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addSellOrders(ret.(*ResponseData), nil)
		ths.addBuyOrders(ret.(*ResponseData), nil)
		ths.addHistorys(ret.(*ResponseData))
		return true
	}

	_L.Error("Bite.ceo : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	// _L.Debug("%s", string(b))
	orders := new(ResponseData)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Bite.ceo : decodeOrders has error :\n%+v", err)
		_L.Trace("Bite.ceo : decodeOrders orgdata [ %s ]", string(b))
	}
	return orders, err
}

func (ths *Reader) addSellOrders(os *ResponseData, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	ths.R(os.Asks)
	for _, val := range os.Asks {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0], 64)
		itm.Amount, _ = strconv.ParseFloat(val[1], 64)
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBuyOrders(os *ResponseData, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0], 64)
		itm.Amount, _ = strconv.ParseFloat(val[1], 64)
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
