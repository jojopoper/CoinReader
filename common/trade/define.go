package trade

// TradeType : sell or buy
type TradeType int

const (
	TradeBuy  TradeType = 1
	TradeSell TradeType = 2
)

// IOrderBase : order base interface define
type IOrderBase interface {
	SetDeci(int, int) IOrderBase
	GetReqParams(float64, float64) string
	GetTradeID() string
}

// IOrderInfo : order info interface define
type IOrderInfo interface {
	SetOrderBase(IOrderBase) IOrderInfo
	GetReqParams(float64, float64) string
	GetTradeID() string
}

// OrderListItem : order list item define
type OrderListItem struct {
	CoinName string
	TradeID  string
	Type     TradeType
	Price    float64
	Amount   float64
	DateTime string
}

// IAsetOrder : aset order interface define
type IAsetOrder interface {
	OrderList(string, string) []*OrderListItem
	SubmitOrder(IOrderInfo, float64, float64) *OrderReslut
	CancelOrder(IOrderBase) *OrderReslut
}

// OrderReslut : trade order result define
type OrderReslut struct {
	IsSuccess bool
	ID        string
	Reason    string
}

// SubmitOrderFunc : submit order function define
type SubmitOrderFunc func(IOrderInfo, float64, float64) *OrderReslut

// CancelOrderFunc : cancel order function define
type CancelOrderFunc func(IOrderBase) *OrderReslut

// GetOrderListFunc : get order list function define
type GetOrderListFunc func(string, string) []*OrderListItem
