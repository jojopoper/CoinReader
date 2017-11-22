package yobit

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) rdOrders() bool {
	address := fmt.Sprintf("https://%s/api/%s/depth/%s_%s?limit=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.depth)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Yobit : GetProxyClient(order) has error \n%+v", err)
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
		ths.addAsksOrders(ret.(*YOrderList), nil)
		ths.addBidsOrders(ret.(*YOrderList), nil)
		return true
	}

	_L.Error("Yobit : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(YOrderList)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		_L.Error("Yobit : decodeOrders (0) has error :\n%+v", err)
	}
	for _, val := range tmp {
		bys, err := json.Marshal(val)
		if err != nil {
			_L.Error("Yobit : decodeOrders (1) has error :\n%+v", err)
		} else {
			err = json.Unmarshal(bys, &orders)
			if err != nil {
				_L.Error("Yobit : decodeOrders (2) has error :\n%+v", err)
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
		itm := &common.OrderBook{}
		itm.Price = val[0]
		itm.Amount = val[1]
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *YOrderList, w *sync.WaitGroup) {
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
