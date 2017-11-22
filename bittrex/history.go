package bittrex

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from poloniex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	address := fmt.Sprintf("https://%s/api/%s/public/getmarkethistory?market=%s-%s",
		ths.Address, ths.currentVer, ths.Monetary, ths.Coin)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Bittrex : GetProxyClient(history) has error \n%+v", err)
				return false
			}
			ths.historyClt.SetClient(client)
		} else {
			ths.historyClt.SetClient(ths.historyClt.GetClient(30))
		}
	}

	ret, err := ths.historyClt.ClientGet(address, rhttp.ReturnCustomType)
	if err == nil {
		ths.addHistorys(ret.(*HistoryResult))
		return true
	}

	_L.Error("Bittrex : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new(HistoryResult)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Bittrex : decodeHistory has error :\n%+v", err)
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs *HistoryResult) {
	ths.Datas.ClearHistorys()
	for _, val := range hs.Result {
		ob := &common.History{}
		ob.DateTime, _ = time.Parse("2006-01-02T15:04:05", val.TimeStamp)
		ob.Type = strings.ToLower(val.OrderType)
		ob.Amount = val.Quantity
		ob.Price = val.Price
		ob.Total = val.Total
		ths.Datas.AddHistory(ob)
	}
}
