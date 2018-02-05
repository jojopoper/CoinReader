package rd1

// ResponseItem : bite.ceo order item datas define
type ResponseItem struct {
	Asks [][]string `json:"s"`
	Bids [][]string `json:"b"`
}

// ResponseData : bite.ceo order list datas define
type ResponseData struct {
	Depth *ResponseItem `json:"depth"`
}
