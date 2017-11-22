package poloniex

// PHistory : poloniex history result json data define
type PHistory struct {
	GlobalTradeID uint64 `json:"globalTradeID"`
	TradeID       uint64 `json:"tradeID"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	Amount        string `json:"amount"`
	Total         string `json:"total"`
}

// POrderList : ploniexe order book result json data define
type POrderList struct {
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
	IsFrozen string          `json:"isFrozen"`
	Seq      uint64          `json:"seq"`
}
