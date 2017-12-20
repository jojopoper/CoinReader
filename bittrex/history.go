package bittrex

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from Bittrex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.(*HistoryResult))
		return true
	}

	_L.Error("Bittrex : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
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
