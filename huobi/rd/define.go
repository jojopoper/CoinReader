package rd

// ResultTick : huobi order book result tick data define
type ResultTick struct {
	Bids    [][]float64 `json:"bids"`
	Asks    [][]float64 `json:"asks"`
	Ts      int64       `json:"ts"`
	Version int64       `json:"version"`
}

// OrderResult : huobi order book result datas define
type OrderResult struct {
	Status string      `json:"status"`
	Ch     string      `json:"ch"`
	Ts     int64       `json:"ts"`
	Tick   *ResultTick `json:"tick"`
}

// TradeItem : huobi trade item data define
type TradeItem struct {
	ID        uint64  `json:"id"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
	Ts        int64   `json:"ts"`
}

// TradeData : huobi trade data define
type TradeData struct {
	ID   uint64       `json:"id"`
	Ts   int64        `json:"ts"`
	Data []*TradeItem `json:"data"`
}

// HistoryResult : huobi history result datas define
type HistoryResult struct {
	Status string       `json:"status"`
	Ch     string       `json:"ch"`
	Ts     int64        `json:"ts"`
	Data   []*TradeData `json:"data"`
}
