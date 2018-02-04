package rd

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from bcex.ca, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientPostForm(ths.HistoryAddr, rhttp.ReturnCustomType, ths.historyParams)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.(*BcexHistoryResp))
		return true
	}

	_L.Error("Bcex.ca : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := new(BcexHistoryResp)
	err := json.Unmarshal(b, historys)
	if err != nil {
		_L.Error("Bcex.ca : decodeHistory has error :\n%+v", err)
		_L.Trace("Bcex.ca : org data is = [\n%s\n]", string(b))
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs *BcexHistoryResp) {
	if hs.Code == 0 {
		for _, val := range hs.Datas {
			ob := &rd.History{}
			tmpTime, _ := strconv.ParseInt(val.DateTime, 10, 64)
			ob.DateTime = time.Unix(tmpTime, 0)
			ob.Type = val.Type
			ob.Amount, _ = strconv.ParseFloat(val.Amount, 64)
			ob.Price = val.Price
			ob.Total = ob.Price * ob.Amount
			ths.Datas.AddHistory(ob)
		}
	} else {
		_L.Error("Bcex.ca history has error: \n > Code = %d\n > Msg = %s", hs.Code, hs.Msg)
	}
}
