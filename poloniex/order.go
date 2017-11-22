package poloniex

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) rdOrders() bool {
	address := fmt.Sprintf("https://%s/public?command=returnOrderBook&depth=%d&currencyPair=%s_%s",
		ths.Address, ths.OrderDepth, ths.Monetary, ths.Coin)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Poloniex : GetProxyClient(order) has error \n%+v", err)
				return false
			}
			ths.orderClt.SetClient(client)
		} else {
			ths.orderClt.SetClient(ths.orderClt.GetClient(30))
		}
	}

	ret, err := ths.orderClt.ClientGet(address, rhttp.ReturnCustomType)
	if err == nil {
		ths.Datas.ClearOrderBook()
		ths.addAsksOrders(ret.(*POrderList), nil)
		ths.addBidsOrders(ret.(*POrderList), nil)
		return true
	}

	_L.Error("Poloniex : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(POrderList)
	err := json.Unmarshal(b, orders)
	if err != nil {
		_L.Error("Poloniex : decodeOrders has error :\n%+v", err)
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *POrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Asks {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0].(string), 64)
		itm.Amount, _ = val[1].(float64)
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *POrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0].(string), 64)
		itm.Amount, _ = val[1].(float64)
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
