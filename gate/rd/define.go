package rd

// HistoryItem : history data item define
type HistoryItem struct {
	TradeID   string  `json:"tradeID"`
	DateTime  string  `json:"date"`
	Timestamp string  `json:"timestamp"`
	Type      string  `json:"type"`
	Rate      float64 `json:"rate"`
	Amount    float64 `json:"amount"`
	Total     float64 `json:"total"`
}

// HistoryData : history data define
type HistoryData struct {
	Result  string         `json:"result"`
	Elapsed string         `json:"elapsed"`
	Datas   []*HistoryItem `json:"data"`
}

// OrderList : order book datas
type OrderList struct {
	Result string      `json:"result"`
	Asks   [][]float64 `json:"asks"`
	Bids   [][]float64 `json:"bids"`
}
