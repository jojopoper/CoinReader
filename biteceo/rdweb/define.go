package rd

// ResponseData : bite.ceo order list datas define
type ResponseData struct {
	Asks [][]string      `json:"sell"`
	Bids [][]string      `json:"buy"`
	Hist [][]interface{} `json:"log"`
}
