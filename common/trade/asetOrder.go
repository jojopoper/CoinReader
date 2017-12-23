package trade

import (
	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/CoinReader/common/nt"
	"github.com/jojopoper/rhttp"
)

// AsetOrder : aset order submit or cancel or list define
type AsetOrder struct {
	common.BaseUserInfo
	nt.NetworkClient
	SubmitOrderAddr string
	CancelOrderAddr string
	OrderListAddr   string
	SubmitCtl       *rhttp.CHttp
	CancelCtl       *rhttp.CHttp
	ListCtl         *rhttp.CHttp
	submit          SubmitOrderFunc
	cancel          CancelOrderFunc
	orderList       GetOrderListFunc
	Result          *OrderReslut
}

// Init : init information
func (ths *AsetOrder) Init(id string, ks []string, v ...interface{}) {
	ths.BaseUserInfo.Init(id, ks...)
	ths.Proxy = nt.GetInitProxy(v...)
	ths.Result = new(OrderReslut)
}

// SetFuncs : set sumbit / cancel / orderlist function
func (ths *AsetOrder) SetFuncs(s SubmitOrderFunc, c CancelOrderFunc, l GetOrderListFunc) {
	ths.submit = s
	ths.cancel = c
	ths.orderList = l
}

// SubmitOrder : submit order
func (ths *AsetOrder) SubmitOrder(o IOrderInfo, price, amount float64) *OrderReslut {
	if ths.submit == nil {
		panic("Have to set sumbit order function")
	}
	return ths.submit(o, price, amount)
}

// CancelOrder : cancel order
func (ths *AsetOrder) CancelOrder(o IOrderBase) *OrderReslut {
	if ths.cancel == nil {
		panic("Have to set cancel order function")
	}
	return ths.cancel(o)
}

// OrderList : return open order list
func (ths *AsetOrder) OrderList(mk, cn string) []*OrderListItem {
	if ths.orderList == nil {
		panic("Have to set order list function")
	}
	return ths.orderList(mk, cn)
}
