package rd

import (
	"encoding/json"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from yobit.net, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientGet(ths.OrderAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*YOrderList), nil)
		ths.addBidsOrders(ret.(*YOrderList), nil)
		return true
	}

	_L.Error("Yobit : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(YOrderList)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		_L.Error("Yobit : decodeOrders (0) has error :\n%+v", err)
		_L.Trace("Yobit : decodeOrders (0) orgdata [ %s ]", string(b))
	}
	for _, val := range tmp {
		bys, err := json.Marshal(val)
		if err != nil {
			_L.Error("Yobit : decodeOrders (1) has error :\n%+v", err)
		} else {
			err = json.Unmarshal(bys, &orders)
			if err != nil {
				_L.Error("Yobit : decodeOrders (2) has error :\n%+v", err)
				_L.Trace("Yobit : decodeOrders (2) orgdata [ %s ]", string(b))
			}
		}
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *YOrderList, w *sync.WaitGroup) {
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

func (ths *Reader) addBidsOrders(os *YOrderList, w *sync.WaitGroup) {
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
