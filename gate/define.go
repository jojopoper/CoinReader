package gate

// ResultData : gate.io result datas define
type ResultData struct {
	Result      bool        `json:"result"`
	AskRate0    float64     `json:"ask_rate0"`
	BidRate0    float64     `json:"bid_rate0"`
	MarketRates bool        `json:"market_rates"`
	RecentRates interface{} `json:"recent_rates"`
	AskList     [][]string  `json:"ask_list"`
	BidList     [][]string  `json:"bid_list"`
	TradeList   [][]string  `json:"trade_list"`
}
