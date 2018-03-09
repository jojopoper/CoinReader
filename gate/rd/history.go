package rd

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from gate.io, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.(*HistoryData))
		return true
	}

	_L.Error("Gate : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new(HistoryData)
	err := json.Unmarshal(b, historys)
	if err != nil {
		_L.Error("Gate : decodeHistory has error :\n%+v", err)
		_L.Trace("Gate : decodeHistory orgdata [ %s ]", string(b))
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs *HistoryData) {
	if hs.Result == "true" {
		ths.R(hs.Datas)
		for _, val := range hs.Datas {
			ob := &rd.History{}
			ts, _ := strconv.ParseInt(val.Timestamp, 10, 64)
			ob.DateTime = time.Unix(ts, 0)
			ob.Type = val.Type
			ob.Amount = val.Amount
			ob.Price = val.Rate
			ob.Total = val.Total
			ths.Datas.AddHistory(ob)
		}
	}
}
