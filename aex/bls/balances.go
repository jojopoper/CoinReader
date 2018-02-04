package bls

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jojopoper/CoinReader/aex"
	"github.com/jojopoper/CoinReader/common/ab"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Balances : aex account balances define
type Balances struct {
	ab.AsetBalances
	aex.AexUserInfo
	version string
}

// Init : init balance struct
func (ths *Balances) Init(id string, ks []string, v ...interface{}) {
	ths.AsetBalances.Init(id, ks, v...)
	ths.baseInit()
	ths.initBalanceParams()
}

func (ths *Balances) baseInit() {
	ths.Address = "aex.com"
	ths.version = "Update time 2017-02-17"
	ths.AexUserInfo.Init(&ths.BaseUserInfo)
	ths.BalanceAddr = fmt.Sprintf("https://api.%s/getMyBalance.php", ths.Address)
	ths.SetUpdateFunc(ths.update)
	ths.SetReturnBalancesFunc(ths.getBalance, ths.getLockBalance)
}

func (ths *Balances) initBalanceParams() {
	ths.RequestClt = new(rhttp.CHttp)
	ths.RequestClt.SetDecodeFunc(ths.decodeBalances)
	if err := ths.SetClient(ths.RequestClt); err != nil {
		_L.Error("Aex : Init requestCtl set http client has error \n%+v", err)
	}
}

func (ths *Balances) clearValues() {
	ths.Values = make(map[string]float64)
}

func (ths *Balances) update() bool {
	ret, err := ths.RequestClt.ClientPostForm(ths.BalanceAddr,
		rhttp.ReturnCustomType, ths.GetReqParams())
	ths.clearValues()
	if err == nil {
		ths.Values = ret.(map[string]float64)
		return true
	}

	_L.Error("Aex : Client get(balance) has error :\n%+v", err)
	ths.initBalanceParams()
	return false
}

func (ths *Balances) getBalance(c string) float64 {
	coin := strings.ToLower(c)
	v, ok := ths.Values[coin+"_balance"]
	if ok {
		return v
	}
	return -1.0
}

func (ths *Balances) getLockBalance(c string) float64 {
	coin := strings.ToLower(c)
	v, ok := ths.Values[coin+"_balance_lock"]
	if ok {
		return v
	}
	return -1.0
}

func (ths *Balances) decodeBalances(b []byte) (interface{}, error) {
	bals := make(map[string]string)
	err := json.Unmarshal(b, &bals)
	if err != nil {
		_L.Trace("%s", string(b))
		_L.Error("Aex : decodeBalances has error :\n%+v", err)
		return nil, err
	}
	ret := make(map[string]float64)
	for k, v := range bals {
		ret[k], _ = strconv.ParseFloat(v, 64)
	}
	return ret, err
}
