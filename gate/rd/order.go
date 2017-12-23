package rd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/jojopoper/CoinReader/common/rd"
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

	addr := fmt.Sprintf("%s%d", ths.OrderAddr, ths.r.Intn(65535))
	ret, err := ths.orderClt.ClientGet(addr, rhttp.ReturnCustomType)
	ths.Datas.ClearOrderBook()
	ths.Datas.ClearHistorys()
	if err == nil {
		rtdata := ret.(*ResultData)
		ths.addAsksOrders(rtdata.AskList, nil)
		ths.addBidsOrders(rtdata.BidList, nil)
		ths.addHistorys(rtdata.TradeList)
		return true
	}

	_L.Error("Gate : Client get has error :\n%+v", err)
	ths.initOrderParams()
	return false
}

func (ths *Reader) decodeOrders(b []byte) (interface{}, error) {
	datas := new(ResultData)
	err := json.Unmarshal(b, &datas)
	if err != nil {
		_L.Error("Gate : decodeOrders has error :\n%+v", err)
	}
	return datas, err
}

func (ths *Reader) addAsksOrders(os [][]string, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0], 64)
		itm.Amount, _ = strconv.ParseFloat(val[1], 64)
		ths.Datas.AddOrder(rd.OrderSellStringKey, itm.Calc())
	}
}

func (ths *Reader) addBidsOrders(os [][]string, w *sync.WaitGroup) {
	if w != nil {
		defer w.Done()
	}
	for _, val := range os {
		itm := &rd.OrderBook{}
		itm.Price, _ = strconv.ParseFloat(val[0], 64)
		itm.Amount, _ = strconv.ParseFloat(val[1], 64)
		ths.Datas.AddOrder(rd.OrderBuyStringKey, itm.Calc())
	}
}
