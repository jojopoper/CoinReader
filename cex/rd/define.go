package rd

// OrderList : cex.io order book list data define
type OrderList struct {
	Timestamp uint64      `json:"timestamp"`
	Asks      [][]float64 `json:"asks"`
	Bids      [][]float64 `json:"bids"`
	Pair      string      `json:"pair"`
	ID        uint        `json:"id"`
	SellTotal string      `json:"sell_total"`
	BuyTotal  string      `json:"buy_total"`
}

// HistoryItem : cex.io history item data define
type HistoryItem struct {
	Type   string `json:"type"`
	Date   string `json:"date"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
	Tid    string `json:"tid"`
}
