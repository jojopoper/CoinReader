package bitfinex

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from Huobi.pro, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	address := fmt.Sprintf("https://api.%s/%s/trades/%s%s?limit_trades=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.limit)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Bitfinex : GetProxyClient(history) has error \n%+v", err)
				return false
			}
			ths.historyClt.SetClient(client)
		} else {
			ths.historyClt.SetClient(ths.historyClt.GetClient(30))
		}
	}

	ret, err := ths.historyClt.ClientGet(address, rhttp.ReturnCustomType)
	if err == nil {
		ths.addHistorys(ret.([]*HistoryItem))
		return true
	}

	_L.Error("Bitfinex : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := make([]*HistoryItem, 0)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Bitfinex : decodeHistory has error :\n%+v", err)
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs []*HistoryItem) {
	ths.Datas.ClearHistorys()
	for _, val := range hs {
		ob := &common.History{}
		ob.DateTime = time.Unix(val.TimeStamp, 0)
		ob.Type = val.Type
		ob.Amount, _ = strconv.ParseFloat(val.Amount, 64)
		ob.Price, _ = strconv.ParseFloat(val.Price, 64)
		ob.Total = ob.Amount * ob.Price
		ths.Datas.AddHistory(ob)
	}
}
