package bitfinex

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
	address := fmt.Sprintf("https://api.%s/%s/book/%s%s?limit_bid=%d&limit_asks=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.limit, ths.limit)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Bitfinex : GetProxyClient(order) has error \n%+v", err)
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
		ths.addAsksOrders(ret.(*OrderList), nil)
		ths.addBidsOrders(ret.(*OrderList), nil)
		return true
	}

	_L.Error("Bitfinex : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderList)
	err := json.Unmarshal(b, &orders)
	if err != nil {
		_L.Error("Bitfinex : decodeOrders has error :\n%+v", err)
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Asks {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val.Price, 64)
		itm.Amount, _ = strconv.ParseFloat(val.Amount, 64)
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderList, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Bids {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val.Price, 64)
		itm.Amount, _ = strconv.ParseFloat(val.Amount, 64)
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
