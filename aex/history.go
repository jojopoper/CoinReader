package aex

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from poloniex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	address := fmt.Sprintf("https://api.%s/trades.php?c=%s&mk_type=%s",
		ths.Address, ths.Coin, ths.Monetary)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Aex : GetProxyClient(history) has error \n%+v", err)
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

	_L.Error("Aex : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new([]*HistoryItem)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Aex : decodeHistory has error :\n%+v", err)
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs []*HistoryItem) {
	ths.Datas.ClearHistorys()
	for _, val := range hs {
		ob := &common.History{}
		ob.DateTime = time.Unix(val.Date, 0)
		ob.Type = val.Type
		ob.Amount = val.Amount
		ob.Price = val.Price
		ob.Total = ob.Amount * ob.Price
		ths.Datas.AddHistory(ob)
	}
}
