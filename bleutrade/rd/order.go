package rd

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from bleutrade.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*OrderDatas), nil)
		ths.addBidsOrders(ret.(*OrderDatas), nil)
		return true
	}

	_L.Error("BleuTrade : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderDatas)
	err := json.Unmarshal(b, &orders)
	if err != nil {
		_L.Error("BleuTrade : decodeOrders  has error :\n%+v", err)
		_L.Trace("BleuTrade : decodeOrders orgdata [ %s ]", string(b))
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderDatas, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result.Sell {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val.Rate, 64)
		itm.Amount, _ = strconv.ParseFloat(val.Quantity, 64)
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderDatas, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result.Buy {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val.Rate, 64)
		itm.Amount, _ = strconv.ParseFloat(val.Quantity, 64)
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
