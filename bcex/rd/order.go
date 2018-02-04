package rd

import (
	"encoding/json"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from bcex.ca, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	// API文档写使用POST请求，但是使用Get也可以得到正确信息
	// ret, err := ths.orderClt.ClientPostForm(ths.OrderAddr, rhttp.ReturnCustomType, ths.orderParams)
	addr := ths.OrderAddr + "?" + ths.orderParams
	ret, err := ths.orderClt.ClientGet(addr, rhttp.ReturnCustomType)

	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*BcexOrderResp), nil)
		ths.addBidsOrders(ret.(*BcexOrderResp), nil)
		return true
	}

	_L.Error("Bcex.ca : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(BcexOrderResp)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Bcex.ca : decodeOrders (0) has error :\n%+v", err)
		_L.Trace("Bcex.ca : org data is = [\n%s\n]", string(b))
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *BcexOrderResp, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	if os.Code == 0 {
		ths.R(os.Datas.Asks)
		for _, val := range os.Datas.Asks {
			itm := &rd.OrderBook{}
			itm.Price = val[0]
			itm.Amount = val[1]
			ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
		}
	} else {
		_L.Error("Bcex.ca order has error: \n > Code = %d\n > Msg = %s", os.Code, os.Msg)
	}
}

func (ths *Reader) addBidsOrders(os *BcexOrderResp, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	if os.Code == 0 {
		for _, val := range os.Datas.Bids {
			itm := &rd.OrderBook{}
			itm.Price = val[0]
			itm.Amount = val[1]
			ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
		}
	} else {
		_L.Error("Bcex.ca order has error: \n > Code = %d\n > Msg = %s", os.Code, os.Msg)
	}
}
