package huobi

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) rdOrders() bool {
	address := fmt.Sprintf("https://api.%s/market/depth?symbol=%s%s&type=step0",
		ths.Address, ths.Coin, ths.Monetary)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Huobi : GetProxyClient(order) has error \n%+v", err)
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
		ths.addAsksOrders(ret.(*OrderResult), nil)
		ths.addBidsOrders(ret.(*OrderResult), nil)
		return true
	}

	_L.Error("Huobi : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderResult)
	err := json.Unmarshal(b, &orders)
	if err != nil {
		_L.Error("Huobi : decodeOrders has error :\n%+v", err)
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderResult, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Tick.Asks {
		itm := &common.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderResult, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Tick.Bids {
		itm := &common.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
