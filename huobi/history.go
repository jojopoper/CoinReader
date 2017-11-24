package huobi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from Huobi.pro, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	address := fmt.Sprintf("https://api.%s/market/history/trade?symbol=%s%s&size=%d",
		ths.Address, ths.Coin, ths.Monetary, ths.Size)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Huobi : GetProxyClient(history) has error \n%+v", err)
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

	_L.Error("Poloniex : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new(HistoryResult)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Huobi : decodeHistory has error :\n%+v", err)
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs *HistoryResult) {
	ths.Datas.ClearHistorys()
	for _, val := range hs.Data {
		for _, itm := range val.Data {
			ob := &common.History{}

			ob.DateTime = time.Unix(ths.splitTime(itm.Ts))
			ob.Type = itm.Direction
			ob.Amount = itm.Amount
			ob.Price = itm.Price
			ob.Total = ob.Amount * ob.Price
			ths.Datas.AddHistory(ob)
		}
	}
}

func (ths *Reader) splitTime(t int64) (int64, int64) {
	sec := t / 1000
	nsec := t % 1000 * 1e6
	return sec, nsec
}
