package common

const (
	OrderBuyStringKey  = "Buy"
	OrderSellStringKey = "Sell"
)

//OrderBook buy&sell data struct
type OrderBook struct {
	Price  float64 // 单价
	Amount float64 // 数量
	Total  float64 // 总额
}

// Calc calculate Total value
func (ths *OrderBook) Calc() *OrderBook {
	ths.Total = ths.Price * ths.Amount
	return ths
}
