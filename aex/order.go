package aex

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) rdOrders() bool {
	address := fmt.Sprintf("https://api.%s/depth.php?c=%s&mk_type=%s",
		ths.Address, ths.Coin, ths.Monetary)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Aex : GetProxyClient(order) has error \n%+v", err)
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
		ths.addSellOrders(ret.(*OrderList), nil)
		ths.addBuyOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Aex : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
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
