package rd

import (
	"encoding/json"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from Cex.io, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*OrderList), nil)
		ths.addBidsOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Cex.io : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Cex.io : decodeOrders has error :\n%+v", err)
		_L.Trace("Cex.io : decodeOrders orgdata [ %s ]", string(b))
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Asks {
		itm := &rd.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &rd.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
