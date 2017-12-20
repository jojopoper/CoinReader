package aex

import (
	"encoding/json"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from aex.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addSellOrders(ret.(*OrderList), nil)
		ths.addBuyOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Aex : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Aex : decodeOrders has error :\n%+v", err)
	}
	return orders, err
}

func (ths *Reader) addSellOrders(os *OrderList, w *sync.WaitGroup) {
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

func (ths *Reader) addBuyOrders(os *OrderList, w *sync.WaitGroup) {
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
