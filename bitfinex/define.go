package bitfinex

// OrderItem : bitfinex order book item data define
type OrderItem struct {
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	TimeStamp string `json:"timestamp"`
}

// OrderList : bitfinex order book datas define
type OrderList struct {
	Asks []*OrderItem `json:"asks"`
	Bids []*OrderItem `json:"bids"`
}

// HistoryItem : bitfinex history datas define
type HistoryItem struct {
	TimeStamp int64  `json:"timestamp"`
	Tid       uint64 `json:"tid"`
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	Exchange  string `json:"exchange"`
	Type      string `json:"type"`
}
