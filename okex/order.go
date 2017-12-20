package okex

import (
	"encoding/json"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from okex.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*OrderList), nil)
		ths.addBidsOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Okex : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Okex : decodeOrders has error :\n%+v", err)
	}
	r := &common.ReverseSlice{}
	r.R(orders.Asks)
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Asks {
		itm := &common.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &common.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
