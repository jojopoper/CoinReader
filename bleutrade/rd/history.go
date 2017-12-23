package rd

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// rdHistory : readout histroy datas from bleutrade.com, datas saved in History datas
func (ths *Reader) rdHistory() bool {
	ret, err := ths.historyClt.ClientGet(ths.HistoryAddr, rhttp.ReturnCustomType)
	ths.Datas.ClearHistorys()
	if err == nil {
		ths.addHistorys(ret.(*HistoryDatas))
		return true
	}

	_L.Error("BleuTrade : Client get(history) has error :\n%+v", err)
	ths.initHistParams()
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
	for _, val := range hs.Result {
		ob := &rd.History{}
		ob.DateTime, _ = time.Parse("2006-01-02 15:04:05", val.TimeStamp)
		ob.Price, _ = strconv.ParseFloat(val.Price, 64)
		ob.Amount, _ = strconv.ParseFloat(val.Quantity, 64)
		ob.Total, _ = strconv.ParseFloat(val.Total, 64)
		ob.Type = strings.ToLower(val.OrderType)
		ths.Datas.AddHistory(ob)
	}
}
