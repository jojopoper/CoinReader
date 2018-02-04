package rd

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdOrders : readout order datas from coinw.com, datas saved in order datas
func (ths *Reader) rdOrders() bool {
	ret, err := ths.orderClt.ClientPostForm(ths.OrderAddr, rhttp.ReturnCustomType, ths.orderParams)
	ths.Datas.ClearOrderBook()
	if err == nil {
		ths.addAsksOrders(ret.(*CoinWOrderResp), nil)
		ths.addBidsOrders(ret.(*CoinWOrderResp), nil)
		return true
	}

	_L.Error("Coinw : Client get(order) has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(CoinWOrderResp)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Coinw : decodeOrders (0) has error :\n%+v", err)
		_L.Trace("Coinw : org data is = [\n%s\n]", string(b))
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *CoinWOrderResp, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	if os.Code == 200 {
		for _, val := range os.Datas.Asks {
			itm := &rd.OrderBook{}
			itm.Price, _ = strconv.ParseFloat(val.Price, 64)
			itm.Amount, _ = strconv.ParseFloat(val.Amount, 64)
			ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
		}
	} else {
		_L.Error("Coinw order has error: \n > Code = %d\n > Msg = %s", os.Code, os.Msg)
	}
}

func (ths *Reader) addBidsOrders(os *CoinWOrderResp, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	if os.Code == 200 {
		for _, val := range os.Datas.Bids {
			itm := &rd.OrderBook{}
			itm.Price, _ = strconv.ParseFloat(val.Price, 64)
			itm.Amount, _ = strconv.ParseFloat(val.Amount, 64)
			ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
		}
	}
}
