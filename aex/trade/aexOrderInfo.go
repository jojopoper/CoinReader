package trade

import (
	"fmt"

	_t "github.com/jojopoper/CoinReader/common/trade"
)

// AexOrderInfo : aex order information define
type AexOrderInfo struct {
	_t.OrderInfo
}

// GetBuyInfo : get buy trade order
func (ths *AexOrderInfo) GetBuyInfo() _t.IOrderInfo {
	ths.OrderType = _t.TradeBuy
	return ths
}

// GetSellInfo : get sell trade order
func (ths *AexOrderInfo) GetSellInfo() _t.IOrderInfo {
	ths.OrderType = _t.TradeSell
	return ths
}

// GetReqParams : get http request parameters
func (ths *AexOrderInfo) GetReqParams(price, amount float64) string {
	if ths.BaseInfo == nil {
		return ""
	}
	return fmt.Sprintf("type=%d&%s", ths.OrderType, ths.BaseInfo.GetReqParams(price, amount))
}
