package bleutrade

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from poloniex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	address := fmt.Sprintf("https://%s/api/%s/public/getmarkethistory?market=%s_%s&count=%d",
		ths.Address, ths.currentVer, ths.Coin, ths.Monetary, ths.depth)
	if ths.historyClt == nil {
		ths.historyClt = new(rhttp.CHttp)
		ths.historyClt.SetDecodeFunc(ths.decodeHistory)

		if ths.Proxy.UseProxy() {
			client, err := ths.historyClt.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
			if err != nil {
				_L.Error("BleuTrade : GetProxyClient(history) has error \n%+v", err)
				return false
			}
			ths.historyClt.SetClient(client)
		} else {
			ths.historyClt.SetClient(ths.historyClt.GetClient(30))
		}
	}

	ret, err := ths.historyClt.ClientGet(address, rhttp.ReturnCustomType)
	if err == nil {
		ths.addHistorys(ret.(*HistoryDatas))
		return true
	}

	_L.Error("BleuTrade : Client get(history) has error :\n%+v", err)
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new(HistoryDatas)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("BleuTrade : decodeHistory (0) has error :\n%+v", err)
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs *HistoryDatas) {
	ths.Datas.ClearHistorys()
	for _, val := range hs.Result {
		ob := &common.History{}
		ob.DateTime, _ = time.Parse("2006-01-02 15:04:05", val.TimeStamp)
		ob.Price, _ = strconv.ParseFloat(val.Price, 64)
		ob.Amount, _ = strconv.ParseFloat(val.Quantity, 64)
		ob.Total, _ = strconv.ParseFloat(val.Total, 64)
		ob.Type = strings.ToLower(val.OrderType)
		ths.Datas.AddHistory(ob)
	}
}
