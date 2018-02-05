package rd

// OrderList : coinegg order list datas define
type OrderList struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

// HistoryItem : coinegg history list item data define
type HistoryItem struct {
	Date   string `json:"date"`
	Price  string `json:"price"`
	Amount string `json:"amount"`
	Tid    string `json:"tid"`
	Type   string `json:"type"`
}
