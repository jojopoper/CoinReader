package rd

import (
	"encoding/json"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from allcoin.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.([]*HistoryItem))
		return true
	}

	_L.Error("Allcoin : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
	return false
}

func (ths *Reader) decodeHistory(b []byte) (interface{}, error) {
	historys := make([]*HistoryItem, 0)
	err := json.Unmarshal(b, &historys)
	if err != nil {
		_L.Error("Allcoin : decodeHistory has error :\n%+v", err)
		_L.Trace("Allcoin : decodeHistory orgdata [ %s ]", string(b))
	}
	return historys, err
}

func (ths *Reader) addHistorys(hs []*HistoryItem) {
	// ths.R(hs)
	for _, val := range hs {
		ob := &rd.History{}
		ob.DateTime = time.Unix(val.Date, 0)
		ob.Type = val.Type
		ob.Amount = val.Amount
		ob.Price = val.Price
		ob.Total = ob.Amount * ob.Price
		ths.Datas.AddHistory(ob)
	}
}
