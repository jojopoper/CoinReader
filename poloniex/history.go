package poloniex

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from poloniex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	address := fmt.Sprintf("https://%s/public?command=returnTradeHistory&currencyPair=%s_%s",
		ths.Address, ths.Monetary, ths.Coin)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Poloniex : GetProxyClient(history) has error \n%+v", err)
				return false
			}
			ths.historyClt.SetClient(client)
		} else {
			ths.historyClt.SetClient(ths.historyClt.GetClient(30))
		}
	}

	ret, err := ths.historyClt.ClientGet(address, rhttp.ReturnCustomType)
	if err == nil {
		ths.addHistorys(ret.([]*PHistory))
		return true
	}

	_L.Error("Poloniex : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := make([]*PHistory, 0)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Poloniex : decodeHistory has error :\n%+v", err)
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs []*PHistory) {
	ths.Datas.ClearHistorys()
	for _, val := range hs {
		ob := &common.History{}
		ob.DateTime, _ = time.Parse("2006-01-02 15:04:05", val.Date)
		ob.Type = val.Type
		ob.Amount, _ = strconv.ParseFloat(val.Amount, 64)
		ob.Price, _ = strconv.ParseFloat(val.Rate, 64)
		ob.Total, _ = strconv.ParseFloat(val.Total, 64)
		ths.Datas.AddHistory(ob)
	}
}
