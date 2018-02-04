package rd

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from bitfinex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.([]*HistoryItem))
		return true
	}

	_L.Error("Bitfinex : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := make([]*HistoryItem, 0)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Bitfinex : decodeHistory has error :\n%+v", err)
		_L.Trace("Bitfinex : decodeHistory orgdata [ %s ]", string(b))
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs []*HistoryItem) {
	for _, val := range hs {
		ob := &rd.History{}
		ob.DateTime = time.Unix(val.TimeStamp, 0)
		ob.Type = val.Type
		ob.Amount, _ = strconv.ParseFloat(val.Amount, 64)
		ob.Price, _ = strconv.ParseFloat(val.Price, 64)
		ob.Total = ob.Amount * ob.Price
		ths.Datas.AddHistory(ob)
	}
}
