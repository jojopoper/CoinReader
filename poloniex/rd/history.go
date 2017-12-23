package rd

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from poloniex.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.([]*PHistory))
		return true
	}

	_L.Error("Poloniex : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
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
	for _, val := range hs {
		ob := &rd.History{}
		ob.DateTime, _ = time.Parse("2006-01-02 15:04:05", val.Date)
		ob.Type = val.Type
		ob.Amount, _ = strconv.ParseFloat(val.Amount, 64)
		ob.Price, _ = strconv.ParseFloat(val.Rate, 64)
		ob.Total, _ = strconv.ParseFloat(val.Total, 64)
		ths.Datas.AddHistory(ob)
	}
}
