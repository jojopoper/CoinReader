package rd

// OrderList : binance order list datas define
type OrderList struct {
	LastUpdateID int64           `json:"lastUpdateId"`
	Asks         [][]interface{} `json:"asks"`
	Bids         [][]interface{} `json:"bids"`
}

// HistoryItem : binance history list item data define
type HistoryItem struct {
	Tid          uint64 `json:"id"`
	Price        string `json:"price"`
	Amount       string `json:"qty"`
	Date         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}
