package trade

// OrderInfo : order information define
type OrderInfo struct {
	OrderType TradeType
	BaseInfo  IOrderBase
}

// SetOrderBase : set order base interface
func (ths *OrderInfo) SetOrderBase(b IOrderBase) IOrderInfo {
	ths.BaseInfo = b
	return ths
}

// GetReqParams : get http request parameters
func (ths *OrderInfo) GetReqParams(price float64, amount float64) string {
	panic("A specific method must be implemented by the inheritance class")
}

// GetTradeID : get trade id
func (ths *OrderInfo) GetTradeID() string {
	return ths.BaseInfo.GetTradeID()
}
