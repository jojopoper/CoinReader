package rd

// YOrderList : yobit order list datas define
type YOrderList struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// HistoryItem : yobit history list item datas define
type HistoryItem struct {
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Tid       uint64  `json:"tid"`
	Timestamp int64   `json:"timestamp"`
}
