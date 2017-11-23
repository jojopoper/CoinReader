package bleutrade

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
	address := fmt.Sprintf("https://%s/api/%s/public/getorderbook?market=%s_%s&type=all&depth=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.depth)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("BleuTrade : GetProxyClient(order) has error \n%+v", err)
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
		ths.addAsksOrders(ret.(*OrderDatas), nil)
		ths.addBidsOrders(ret.(*OrderDatas), nil)
		return true
	}

	_L.Error("BleuTrade : Client get(order) has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	orders := new(OrderDatas)
	err := json.Unmarshal(b, &orders)
	if err != nil {
		_L.Error("BleuTrade : decodeOrders  has error :\n%+v", err)
	}
	return orders, err
}

func (ths *Reader) addAsksOrders(os *OrderDatas, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result.Sell {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val.Rate, 64)
		itm.Amount, _ = strconv.ParseFloat(val.Quantity, 64)
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os *OrderDatas, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os.Result.Buy {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val.Rate, 64)
		itm.Amount, _ = strconv.ParseFloat(val.Quantity, 64)
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
