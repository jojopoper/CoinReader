package rd

import (
	"strconv"
	"time"

	"github.com/jojopoper/CoinReader/common/rd"
)

func (ths *Reader) addHistorys(hs *ResponseData) {
	// ths.R(hs)
	for _, val := range hs.Hist {
		ob := &rd.History{}
		ob.DateTime, _ = time.Parse("15:04:05", val[0].(string))
		if val[1].(float64) == 1 {
			ob.Type = "buy"
		} else if val[1].(float64) == 2 {
			ob.Type = "sell"
		}
		ob.Price, _ = strconv.ParseFloat(val[2].(string), 64)
		ob.Amount, _ = strconv.ParseFloat(val[3].(string), 64)
		ob.Total = ob.Amount * ob.Price
		ths.Datas.AddHistory(ob)
	}
}
