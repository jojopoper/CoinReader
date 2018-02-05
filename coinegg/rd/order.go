package rd

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from coinegg.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addSellOrders(ret.(*OrderList), nil)
		ths.addBuyOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Coinegg : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	// _L.Debug("%s", string(b))
	orders := new(OrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Coinegg : decodeOrders has error :\n%+v", err)
		_L.Trace("Coinegg : decodeOrders orgdata [ %s ]", string(b))
	}
	return orders, err
}

func (ths *Reader) addSellOrders(os *OrderList, w *sync.WaitGroup) {
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

func (ths *Reader) addBuyOrders(os *OrderList, w *sync.WaitGroup) {
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
