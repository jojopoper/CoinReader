package rd

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from poloniex.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*POrderList), nil)
		ths.addBidsOrders(ret.(*POrderList), nil)
		return true
	}

	_L.Error("Poloniex : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(POrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Poloniex : decodeOrders has error :\n%+v", err)
		_L.Trace("Poloniex : decodeOrders orgdata [ %s ]", string(b))
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *POrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Asks {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0].(string), 64)
		itm.Amount, _ = val[1].(float64)
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *POrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0].(string), 64)
		itm.Amount, _ = val[1].(float64)
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
