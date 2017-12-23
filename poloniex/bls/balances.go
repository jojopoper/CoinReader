package bls

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jojopoper/CoinReader/common/ab"
	"github.com/jojopoper/CoinReader/common/tool"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// Balances : poloniex account balances define
type Balances struct {
	ab.AsetBalances
	tool.CryptoUtils
}

// Init : init balance struct
func (ths *Balances) Init(id string, ks []string, v ...interface{}) {
	ths.AsetBalances.Init(id, ks, v...)
	ths.baseInit()
	ths.initBalanceParams()
}

func (ths *Balances) baseInit() {
	ths.Address = "poloniex.com"
	ths.BalanceAddr = fmt.Sprintf("https://%s/tradingApi", ths.Address)
	ths.SetUpdateFunc(ths.update)
	ths.SetReturnBalancesFunc(ths.getBalance, ths.getLockBalance)
}

func (ths *Balances) initBalanceParams() {
	ths.RequestClt = new(rhttp.CHttp)
	ths.RequestClt.SetDecodeFunc(ths.decodeBalances)
	if err := ths.SetClient(ths.RequestClt); err != nil {
		_L.Error("Poloniex : Init requestCtl set http client has error \n%+v", err)
	}
}

func (ths *Balances) clearValues() {
	ths.Values = make(map[string]float64)
}

func (ths *Balances) update() bool {
	tm := time.Now().UnixNano()
	datas := fmt.Sprintf("command=returnCompleteBalances&nonce=%d", tm)
	h512result, _ := ths.GetHmacSHA512(ths.Keys[1], datas)
	ths.clearValues()
	ret, err := ths.RequestClt.ClientPostFormWithHeader(ths.BalanceAddr, rhttp.ReturnCustomType,
		datas, map[string]string{
			"Key":  ths.Keys[0],
			"Sign": h512result,
		})
	if err == nil {
		ths.Values = ret.(map[string]float64)
		return true
	}

	_L.Error("Poloniex : Client get(balance) has error :\n%+v", err)
	ths.initBalanceParams()
	return false
}

func (ths *Balances) getBalance(c string) float64 {
	coin := strings.ToUpper(c)
	v, ok := ths.Values[coin+"_balance"]
	if ok {
		return v
	}
	return -1.0
}

func (ths *Balances) getLockBalance(c string) float64 {
	coin := strings.ToUpper(c)
	v, ok := ths.Values[coin+"_balance_lock"]
	if ok {
		return v
	}
	return -1.0
}

func (ths *Balances) decodeBalances(b []byte) (interface{}, error) {
	hasErr := make(map[string]interface{})
	err := json.Unmarshal(b, &hasErr)
	if err != nil {
		return nil, err
	}
	if hasErr["error"] != nil {
		return nil, fmt.Errorf("%s", hasErr["error"].(string))
	}
	bals := make(map[string]*BalanceItem)
	err = json.Unmarshal(b, &bals)
	if err != nil {
		_L.Error("Poloniex : decodeBalances has error :\n%+v", err)
		return nil, err
	}
	ret := make(map[string]float64)
	for k, v := range bals {
		ret[fmt.Sprintf("%s_balance", k)], _ = strconv.ParseFloat(v.Available, 64)
		ret[fmt.Sprintf("%s_balance_lock", k)], _ = strconv.ParseFloat(v.OnOrders, 64)
	}
	return ret, err
}
