package bittrex

const (
	OrderBuyKey = "buy"
	OrderSellKey = "sell"
)

// OrderItem : order item define
type OrderItem struct {
	Quantity float64 `json:"Quantity"`
	Rate float64 `json:"Rate"`
}

// OrderBookAll : buy and sell order book define
type OrderBookAll struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Result map[string][]*OrderItem `json:"result"`
}

// HistoryItem : history data item define
type HistoryItem struct {
	ID uint64 `json:"Id"`
	TimeStamp string `json:"TimeStamp"`
	Quantity float64 `json:"Quantity"`
	Price float64 `json:"Price"`
	Total float64 `json:"Total"`
	FillType string `json:"FillType"`
	OrderType string `json:"OrderType"`
}

// HistoryResult : history data result define
type HistoryResult struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Result []*HistoryItem `json:"result"`
}