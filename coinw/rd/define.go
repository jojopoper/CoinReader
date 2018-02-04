package rd

// CoinWRespHeader : response header define
type CoinWRespHeader struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

// OrderItem : order item define
type OrderItem struct {
	ID     int    `json:"id"`
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

// CoinWOrderList : orders list for coinw.com define
type CoinWOrderList struct {
	Bids []*OrderItem `json:"bids"`
	Asks []*OrderItem `json:"asks"`
}

// CoinWOrderResp : order response define
type CoinWOrderResp struct {
	CoinWRespHeader
	Datas *CoinWOrderList `json:"data"`
}

// HistoryItem : history item define
type HistoryItem struct {
	OrderItem
	Time   string `json:"time"`
	EnType string `json:"en_type"`
	Type   string `json:"type"`
}

// CoinWHistoryResp : history response define
type CoinWHistoryResp struct {
	CoinWRespHeader
	Datas []*HistoryItem `json:"data"`
}
