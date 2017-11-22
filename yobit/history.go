package yobit

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
	address := fmt.Sprintf("https://%s/api/%s/trades/%s_%s?limit=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.depth)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("Yobit : GetProxyClient(history) has error \n%+v", err)
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

	_L.Error("Yobit : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := make([]*HistoryItem, 0)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		_L.Error("Yobit : decodeHistory (0) has error :\n%+v", err)
	}
	for _, val := range tmp {
		bys, err := json.Marshal(val)
		if err != nil {
			_L.Error("Yobit : decodeHistory (1) has error :\n%+v", err)
		} else {
			err = json.Unmarshal(bys, &historys)
			if err != nil {
				_L.Error("Yobit : decodeHistory (2) has error :\n%+v", err)
			}
		}
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs []*HistoryItem) {
	ths.Datas.ClearHistorys()
	for _, val := range hs {
		ob := &common.History{}
		ob.DateTime = time.Unix(val.Timestamp, 0)
		if val.Type == "bid" {
			ob.Type = "buy"
		} else if val.Type == "ask" {
			ob.Type = "sell"
		}
		ob.Amount = val.Amount
		ob.Price = val.Price
		ob.Total = ob.Amount * ob.Price
		ths.Datas.AddHistory(ob)
	}
}
