package rd

import (
	"encoding/json"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from Bittrex.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addSellOrders(ret.(*OrderBookAll), nil)
		ths.addBuyOrders(ret.(*OrderBookAll), nil)
		return true
	}

	_L.Error("Bittrex : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderBookAll)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Bittrex : decodeOrders has error :\n%+v", err)
	}
	return orders, err
}

func (ths *Reader) addSellOrders(os *OrderBookAll, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result[OrderSellKey] {
		itm := &rd.OrderBook{}
		itm.Price = val.Rate
		itm.Amount = val.Quantity
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBuyOrders(os *OrderBookAll, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result[OrderBuyKey] {
		itm := &rd.OrderBook{}
		itm.Price = val.Rate
		itm.Amount = val.Quantity
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
