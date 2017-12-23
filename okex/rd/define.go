package rd

// OrderList : okex order book list data define
type OrderList struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// HistoryItem : okex history item data define
type HistoryItem struct {
	Date   int64   `json:"date"`
	DateMs int64   `json:"date_ms"`
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	Type   string  `json:"type"`
	Tid    uint64  `json:"tid"`
}
