package rd

import (
	"encoding/json"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from Huobi.pro, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.(*HistoryResult))
		return true
	}

	_L.Error("Huobi : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
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
	for _, val := range hs.Data {
		for _, itm := range val.Data {
			ob := &rd.History{}

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
