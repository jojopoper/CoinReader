package gate

import (
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common"
)

func (ths *Reader) addHistorys(hs [][]string) {
	ths.Datas.ClearHistorys()
	for _, val := range hs {
		ob := &common.History{}
		ob.DateTime, _ = time.Parse("15:04:05", val[0])
		ob.Price, _ = strconv.ParseFloat(val[1], 64)
		ob.Amount, _ = strconv.ParseFloat(val[2], 64)
		ob.Total, _ = strconv.ParseFloat(val[3], 64)
		ob.Type = val[5]
		ths.Datas.AddHistory(ob)
	}
}
