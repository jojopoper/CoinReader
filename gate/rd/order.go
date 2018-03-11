package rd

import (
	"encoding/json"
	"reflect"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from gate.io, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*OrderList), nil)
		ths.addBidsOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Gate : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Gate : decodeOrders has error :\n%+v", err)
		_L.Trace("Gate : decodeOrders orgdata [ %s ]", string(b))
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	ths.R(os.Asks)
	for _, val := range os.Asks {
		itm := &rd.OrderBook{}
		if reflect.TypeOf(val[0]).Kind() == reflect.Float64 {
			itm.Price = val[0].(float64)
		} else {
			itm.Price, _ = strconv.ParseFloat(val[0].(string), 64)
		}
		if reflect.TypeOf(val[1]).Kind() == reflect.Float64 {
			itm.Amount = val[1].(float64)
		} else {
			itm.Amount, _ = strconv.ParseFloat(val[1].(string), 64)
		}
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &rd.OrderBook{}
		if reflect.TypeOf(val[0]).Kind() == reflect.Float64 {
			itm.Price = val[0].(float64)
		} else {
			itm.Price, _ = strconv.ParseFloat(val[0].(string), 64)
		}
		if reflect.TypeOf(val[1]).Kind() == reflect.Float64 {
			itm.Amount = val[1].(float64)
		} else {
			itm.Amount, _ = strconv.ParseFloat(val[1].(string), 64)
		}
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
