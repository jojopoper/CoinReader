package bls

// BalanceItem : user account balance item define
type BalanceItem struct {
	Available string `json:"available"`
	OnOrders  string `json:"onOrders"`
	BtcValue  string `json:"btcValue"`
}
