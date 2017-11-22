package bittrex

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) rdOrders() bool {
	address := fmt.Sprintf("https://%s/api/%s/public/getorderbook?market=%s-%s&type=both",
		ths.Address, ths.currentVer, ths.Monetary, ths.Coin)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Bittrex : GetProxyClient(order) has error \n%+v", err)
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
		ths.addSellOrders(ret.(*OrderBookAll), nil)
		ths.addBuyOrders(ret.(*OrderBookAll), nil)
		return true
	}

	_L.Error("Bittrex : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
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
		itm := &common.OrderBook{}
		itm.Price = val.Rate
		itm.Amount = val.Quantity
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBuyOrders(os *OrderBookAll, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result[OrderBuyKey] {
		itm := &common.OrderBook{}
		itm.Price = val.Rate
		itm.Amount = val.Quantity
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
