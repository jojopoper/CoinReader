package gate

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

func (ths *Reader) endReading() {
	ths.Lock()
	defer ths.Unlock()
	ths.isReading = false
}

func (ths *Reader) reading() bool {
	ths.Lock()
	if ths.isReading {
		defer ths.Unlock()
		return true
	}
	ths.isReading = true
	ths.Unlock()
	defer ths.endReading()

	address := fmt.Sprintf("https://%s/json_svr/query/?u=11&c=%d&type=ask_bid_list_table&symbol=%s_%s",
		ths.Address, ths.r.Intn(65535), ths.Coin, ths.Monetary)
	if ths.orderClt == nil {
		ths.orderClt = new(rhttp.CHttp)
		ths.orderClt.SetDecodeFunc(ths.decodeOrders)

		if ths.Proxy.UseProxy() {
			client, err := ths.orderClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Gate : GetProxyClient has error \n%+v", err)
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
		rtdata := ret.(*ResultData)
		ths.addAsksOrders(rtdata.AskList, nil)
		ths.addBidsOrders(rtdata.BidList, nil)
		ths.addHistorys(rtdata.TradeList)
		return true
	}

	_L.Error("Gate : Client get has error :\n%+v", err)
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	datas := new(ResultData)
	err := json.Unmarshal(b, &datas)
	if err != nil {
		_L.Error("Bittrex : decodeOrders has error :\n%+v", err)
	}
	return datas, err
}

func (ths *Reader) addAsksOrders(os [][]string, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0], 64)
		itm.Amount, _ = strconv.ParseFloat(val[1], 64)
		ths.Datas.AddOrder(common.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os [][]string, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os {
		itm := &common.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0], 64)
		itm.Amount, _ = strconv.ParseFloat(val[1], 64)
		ths.Datas.AddOrder(common.OrderBuyStringKey, itm.Calc())
	}
}
