package aex

// OrderList : aex order list datas define
type OrderList struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// HistoryItem : aex history list item data define
type HistoryItem struct {
	Date   int64   `json:"date"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    uint64  `json:"tid"`
	Type   string  `json:"type"`
}
