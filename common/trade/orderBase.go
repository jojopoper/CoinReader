package trade

import (
	"fmt"
)

// OrderBaseInfo : order base information define
type OrderBaseInfo struct {
	MK           string
	CoinName     string
	PriceFormat  string
	AmountFormat string
	TradeID      string
	FormatStr    string
}

// Init : init struct info
func (ths *OrderBaseInfo) Init(mk, cn, id string) {
	ths.MK = mk
	ths.CoinName = cn
	ths.TradeID = id
}

// SetDeci : set deci for price and amount
func (ths *OrderBaseInfo) SetDeci(pd, ad int) IOrderBase {
	ths.PriceFormat = fmt.Sprintf("%%.%df", pd)
	ths.AmountFormat = fmt.Sprintf("%%.%df", ad)
	return ths
}

// GetReqParams : get http request params
func (ths *OrderBaseInfo) GetReqParams(price float64, amount float64) string {
	panic("A specific method must be implemented by the inheritance class")
}

// GetTradeID : get trade id
func (ths *OrderBaseInfo) GetTradeID() string {
	return ths.TradeID
}
