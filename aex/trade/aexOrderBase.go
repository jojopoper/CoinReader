package trade

import (
	"fmt"

	_t "github.com/jojopoper/CoinReader/common/trade"
)

// AexOrderBase : aex order basic infor define
type AexOrderBase struct {
	_t.OrderBaseInfo
}

// SetDeci : set deci for price and amount
func (ths *AexOrderBase) SetDeci(pd, ad int) _t.IOrderBase {
	ths.OrderBaseInfo.SetDeci(pd, ad)
	ths.FormatStr = fmt.Sprintf("mk_type=%%s&coinname=%%s&price=%s&amount=%s", ths.PriceFormat, ths.AmountFormat)
	return ths
}

// GetReqParams : get http request params
func (ths *AexOrderBase) GetReqParams(price float64, amount float64) string {
	return fmt.Sprintf(ths.FormatStr, ths.MK, ths.CoinName, price, amount)
}

// AexOrderCancel : aex order cancel define
type AexOrderCancel struct {
	_t.OrderBaseInfo
}

// GetReqParams : get http request params
func (ths *AexOrderCancel) GetReqParams(price float64, amount float64) string {
	return fmt.Sprintf("mk_type=%s&coinname=%s&order_id=%s", ths.MK, ths.CoinName, ths.TradeID)
}
