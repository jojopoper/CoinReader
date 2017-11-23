package bleutrade

// OrderItem : order item data define
type OrderItem struct {
	Quantity string `json:"Quantity"`
	Rate     string `json:"Rate"`
}

// OrderResult : order book result define
type OrderResult struct {
	Buy  []*OrderItem `json:"buy"`
	Sell []*OrderItem `json:"sell"`
}

// OrderDatas : order book datas define
type OrderDatas struct {
	Success string       `json:"success"`
	Message string       `json:"message"`
	Result  *OrderResult `json:"result"`
}

// HistoryItem : history item data define
type HistoryItem struct {
	TimeStamp string `json:"TimeStamp"`
	Quantity  string `json:"Quantity"`
	Price     string `json:"Price"`
	Total     string `json:"Total"`
	OrderType string `json:"OrderType"`
}

// HistoryDatas : history datas define
type HistoryDatas struct {
	Success string         `json:"success"`
	Message string         `json:"message"`
	Result  []*HistoryItem `json:"result"`
}
