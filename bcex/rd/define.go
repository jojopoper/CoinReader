package rd

// BcexRespHeader : response header define
type BcexRespHeader struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// BcexOrderList : orders list for bcex.ca define
type BcexOrderList struct {
	Bids     [][]float64 `json:"bids"`
	Asks     [][]float64 `json:"asks"`
	DateTime int64       `json:"date"`
}

// BcexOrderResp : order response define
type BcexOrderResp struct {
	BcexRespHeader
	Datas *BcexOrderList `json:"data"`
}

// HistoryItem : history item define
type HistoryItem struct {
	Type     string  `json:"type"`
	DateTime string  `json:"date"`
	Price    float64 `json:"price"`
	Amount   string  `json:"amount"`
	Tid      string  `json:"tid"`
}

// BcexHistoryResp : history response define
type BcexHistoryResp struct {
	BcexRespHeader
	Datas []*HistoryItem `json:"data"`
}
