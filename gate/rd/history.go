package rd

import (
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
)

func (ths *Reader) addHistorys(hs [][]string) {
	for _, val := range hs {
		ob := &rd.History{}
		ob.DateTime, _ = time.Parse("15:04:05", val[0])
		ob.Price, _ = strconv.ParseFloat(val[1], 64)
		ob.Amount, _ = strconv.ParseFloat(val[2], 64)
		ob.Total, _ = strconv.ParseFloat(val[3], 64)
		ob.Type = val[5]
		ths.Datas.AddHistory(ob)
	}
}
