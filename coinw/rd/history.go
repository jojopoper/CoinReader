package rd

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from coinw.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientPostForm(ths.HistoryAddr, rhttp.ReturnCustomType, ths.historyParams)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.(*CoinWHistoryResp))
		return true
	}

	_L.Error("Coinw : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new(CoinWHistoryResp)
	err := json.Unmarshal(b, historys)
	if err != nil {
		_L.Error("Coinw : decodeHistory has error :\n%+v", err)
		_L.Trace("Coinw : org data is = [\n%s\n]", string(b))
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs *CoinWHistoryResp) {
	if hs.Code == 200 {
		for _, val := range hs.Datas {
			ob := &rd.History{}
			ob.DateTime, _ = time.Parse("15:04:05", val.Time)
			if val.EnType == "bid" {
				ob.Type = "buy"
			} else if val.EnType == "ask" {
				ob.Type = "sell"
			}
			ob.Amount, _ = strconv.ParseFloat(val.Amount, 64)
			ob.Price, _ = strconv.ParseFloat(val.Price, 64)
			ob.Total = ob.Price * ob.Amount
			ths.Datas.AddHistory(ob)
		}
	} else {
		_L.Error("Coinw history has error: \n > Code = %d\n > Msg = %s", hs.Code, hs.Msg)
	}
}
