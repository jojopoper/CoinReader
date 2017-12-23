package ab

import (
	"sync"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/CoinReader/common/nt"
	"github.com/jojopoper/rhttp"
)

// IAsetBalance : aset balance interface define
type IAsetBalance interface {
	Update() bool
	GetBalance(string) float64
	GetLockBalance(string) float64
}

// UpdateFunc : update balance function define
type UpdateFunc func() bool

// ReturnBalance : return balance value function define
type ReturnBalance func(string) float64

// AsetBalances : account balance define
type AsetBalances struct {
	common.BaseUserInfo
	nt.NetworkClient
	sync.Mutex
	BalanceAddr        string
	RequestClt         *rhttp.CHttp
	Values             map[string]float64
	updateFunc         UpdateFunc
	retBalanceFunc     ReturnBalance
	retLockBalanceFunc ReturnBalance
}

// Init : init balance struct
func (ths *AsetBalances) Init(id string, ks []string, v ...interface{}) {
	ths.BaseUserInfo.Init(id, ks...)
	ths.Proxy = nt.GetInitProxy(v...)
	ths.Values = make(map[string]float64)
}

// SetUpdateFunc : set update balance function
func (ths *AsetBalances) SetUpdateFunc(f UpdateFunc) {
	ths.updateFunc = f
}

// SetReturnBalancesFunc : set return balance and locked balance function
func (ths *AsetBalances) SetReturnBalancesFunc(f ReturnBalance, lf ReturnBalance) {
	ths.retBalanceFunc = f
	ths.retLockBalanceFunc = lf
}

// Update : update user account balances
func (ths *AsetBalances) Update() bool {
	if ths.updateFunc == nil {
		panic("Have to set update balance function")
	}
	ths.Lock()
	defer ths.Unlock()
	return ths.updateFunc()
}

// GetBalance : return balance can use to trade by cn(coin name)
func (ths *AsetBalances) GetBalance(cn string) float64 {
	if ths.retBalanceFunc == nil {
		panic("Have to set return balance function")
	}
	ths.Lock()
	defer ths.Unlock()
	return ths.retBalanceFunc(cn)
}

// GetLockBalance : return locked balance by cn(coin name)
func (ths *AsetBalances) GetLockBalance(cn string) float64 {
	if ths.retLockBalanceFunc == nil {
		panic("Have to set return balance function")
	}
	ths.Lock()
	defer ths.Unlock()
	return ths.retLockBalanceFunc(cn)
}
