package rd

// OrderList : allcoin order list datas define
type OrderList struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// HistoryItem : allcoin history list item data define
type HistoryItem struct {
	Date   int64   `json:"date"`
	DateMs int64   `json:"date_ms"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    string  `json:"tid"`
	Type   string  `json:"type"`
}
