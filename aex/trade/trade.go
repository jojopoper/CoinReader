package trade

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jojopoper/CoinReader/aex"
	"github.com/jojopoper/CoinReader/common/trade"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

// TradeOrder : submit order and cancel order
type TradeOrder struct {
	trade.AsetOrder
	aex.AexUserInfo
	version string
}

// Init : init trade information
func (ths *TradeOrder) Init(id string, ks []string, v ...interface{}) {
	ths.AsetOrder.Init(id, ks, v...)
	ths.baseInit()
	ths.initSubmitParams()
	ths.initCancelParams()
	ths.initListOrderParams()
}

func (ths *TradeOrder) baseInit() {
	ths.Address = "aex.com"
	ths.version = "Update time 2017-02-17"
	ths.AexUserInfo.Init(&ths.BaseUserInfo)
	ths.SubmitOrderAddr = fmt.Sprintf("https://api.%s/submitOrder.php", ths.Address)
	ths.CancelOrderAddr = fmt.Sprintf("https://api.%s/cancelOrder.php", ths.Address)
	ths.OrderListAddr = fmt.Sprintf("https://api.%s/getOrderList.php", ths.Address)
	ths.SetFuncs(ths.submitOrder, ths.cancelOrder, ths.queryOrders)
}

func (ths *TradeOrder) initSubmitParams() {
	ths.SubmitCtl = new(rhttp.CHttp)
	// ths.SubmitCtl.SetDecodeFunc(ths.decodeSubmitOrder)
	if err := ths.SetClient(ths.SubmitCtl); err != nil {
		_L.Error("Aex : Init submit order Ctl set http client has error \n%+v", err)
	}
}

func (ths *TradeOrder) initCancelParams() {
	ths.CancelCtl = new(rhttp.CHttp)
	// ths.CancelCtl.SetDecodeFunc(ths.decodeCancelOrder)
	if err := ths.SetClient(ths.CancelCtl); err != nil {
		_L.Error("Aex : Init cancel order Ctl set http client has error \n%+v", err)
	}
}

func (ths *TradeOrder) initListOrderParams() {
	ths.ListCtl = new(rhttp.CHttp)
	ths.ListCtl.SetDecodeFunc(ths.decodeListOrder)
	if err := ths.SetClient(ths.ListCtl); err != nil {
		_L.Error("Aex : Init list order Ctl set http client has error \n%+v", err)
	}
}

func (ths *TradeOrder) submitOrder(o trade.IOrderInfo, p, a float64) *trade.OrderReslut {
	datas := fmt.Sprintf("%s&%s", ths.GetReqParams(), o.GetReqParams(p, a))
	// _L.Debug("Post data = %s", datas)
	ret, err := ths.SubmitCtl.ClientPostForm(ths.SubmitOrderAddr,
		rhttp.ReturnString, datas)
	if err == nil {
		// succ|1786
		// succ
		// overBalance
		tmp := ret.(string)
		rets := strings.Split(tmp, "|")
		ths.Result.IsSuccess = rets[0] == "succ"
		if ths.Result.IsSuccess && len(rets) == 2 {
			ths.Result.ID = rets[1]
		} else {
			ths.Result.ID = ""
		}
		ths.Result.Reason = tmp
	} else {
		_L.Error("Aex : Client submitOrder has error :\n%+v", err)
		ths.initSubmitParams()
		ths.Result.IsSuccess = false
		ths.Result.Reason = "Net error"
		ths.Result.ID = ""
	}

	return ths.Result
}

func (ths *TradeOrder) cancelOrder(o trade.IOrderBase) *trade.OrderReslut {
	datas := fmt.Sprintf("%s&%s", ths.GetReqParams(), o.GetReqParams(0.0, 0.0))
	// _L.Debug("Post data = %s", datas)
	ret, err := ths.CancelCtl.ClientPostForm(ths.CancelOrderAddr,
		rhttp.ReturnString, datas)
	ths.Result.ID = o.GetTradeID()
	if err == nil {
		// succ
		// no_record
		tmp := ret.(string)
		ths.Result.IsSuccess = tmp == "succ"
		ths.Result.Reason = tmp
	} else {
		_L.Error("Aex : Client cancelOrder has error :\n%+v", err)
		ths.initCancelParams()
		ths.Result.IsSuccess = false
		ths.Result.Reason = "Net error"
	}

	return ths.Result
}

func (ths *TradeOrder) queryOrders(mk, cn string) []*trade.OrderListItem {
	datas := fmt.Sprintf("%s&mk_type=%s&coinname=%s", ths.GetReqParams(), mk, cn)
	// _L.Debug("Post data = %s", datas)
	ret, err := ths.ListCtl.ClientPostForm(ths.OrderListAddr,
		rhttp.ReturnCustomType, datas)
	if err == nil {
		if ret != nil {
			return ret.([]*trade.OrderListItem)
		}
		return make([]*trade.OrderListItem, 0)
	}
	_L.Error("Aex : Client queryOrders has error :\n%+v", err)
	ths.initListOrderParams()
	return nil
}

// OrderOrigData : order item orig data
type OrderOrigData struct {
	ID       string `json:"id"`
	CoinName string `json:"coinname"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
	DateTime string `json:"time"`
}

func (ths *TradeOrder) decodeListOrder(b []byte) (interface{}, error) {
	// no_order
	// _L.Warning("%s", string(b))
	tmp := string(b)
	retDatas := make([]*trade.OrderListItem, 0)
	if tmp == "no_order" {
		return retDatas, nil
	}
	orgDatas := make([]*OrderOrigData, 0)
	err := json.Unmarshal(b, &orgDatas)
	if err == nil {
		for _, v := range orgDatas {
			itm := new(trade.OrderListItem)
			itm.CoinName = v.CoinName
			itm.TradeID = v.ID
			if v.Type == "2" {
				itm.Type = trade.TradeSell
			} else if v.Type == "1" {
				itm.Type = trade.TradeBuy
			}
			itm.Amount, _ = strconv.ParseFloat(v.Amount, 64)
			itm.Price, _ = strconv.ParseFloat(v.Price, 64)
			itm.DateTime = v.DateTime
			retDatas = append(retDatas, itm)
		}
	}
	return retDatas, err
}
