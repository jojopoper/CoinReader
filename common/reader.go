package common

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jojopoper/rhttp"
)

// ReadFunc : readout function define
type ReadFunc func() bool

// ReaderInterface : 读取接口定义
type ReaderInterface interface {
	ReadAll() bool
	ReadHistorys() bool
	ReadOrders() bool
	PrintOrders(int) string
	PrintHistorys(int) string
	GetResultDatas() *Results
}

// ReaderDef : reader base struct
type ReaderDef struct {
	Address      string
	Monetary     string
	Coin         string
	Datas        *Results
	Proxy        *ProxyDef
	readHistorys ReadFunc
	readOrders   ReadFunc
	orderLock    *sync.Mutex
	historyLock  *sync.Mutex
	OrderAddr    string
	HistoryAddr  string
}

// Init init parameters
// m is MonetaryName string
// c is CoinName string
// v is optional parameters, Set the parameter order described below:
// rof is read order function point
// rhf is read history functon point
// v[0] -- proxyAddress string
// v[1] -- proxyPort string
func (ths *ReaderDef) Init(m, c string, rof, rhf ReadFunc, v ...interface{}) {
	ths.Monetary = m
	ths.Coin = c
	ths.readHistorys = rhf
	ths.readOrders = rof
	ths.orderLock = new(sync.Mutex)
	ths.historyLock = new(sync.Mutex)

	ths.Datas = GetInitResults()
	ths.Proxy = GetInitProxy(v...)
}

// ReadAll : readout order and history datas from website
func (ths *ReaderDef) ReadAll() bool {
	return ths.readHistorys() && ths.readOrders()
}

// ReadHistorys : readout history datas from website
func (ths *ReaderDef) ReadHistorys() bool {
	ths.historyLock.Lock()
	defer ths.historyLock.Unlock()
	if ths.readHistorys == nil {
		panic("You have to set readHistory function in basic struct")
	}
	return ths.readHistorys()
}

// ReadOrders : readout order datas from website
func (ths *ReaderDef) ReadOrders() bool {
	ths.orderLock.Lock()
	defer ths.orderLock.Unlock()
	if ths.readOrders == nil {
		panic("You have to set readOrders function in basic struct")
	}
	return ths.readOrders()
}

// GetResultDatas : get website result datas
func (ths *ReaderDef) GetResultDatas() *Results {
	ths.orderLock.Lock()
	ths.historyLock.Lock()
	defer ths.orderLock.Unlock()
	defer ths.historyLock.Unlock()
	return ths.Datas
}

// PrintOrders : print order datas to string
func (ths *ReaderDef) PrintOrders(length int) string {
	ths.orderLock.Lock()
	defer ths.orderLock.Unlock()
	if ths.Datas.Orders == nil {
		return "> No datas!!\r\n"
	}

	buyList, _ := ths.Datas.Orders[OrderBuyStringKey]
	sellList, _ := ths.Datas.Orders[OrderSellStringKey]

	relLenBuy := len(buyList)
	relLenSell := len(sellList)

	if length != -1 {
		if length < relLenBuy {
			relLenBuy = length
		}
		if length < relLenSell {
			relLenSell = length
		}
	}

	ret := fmt.Sprintf("\r\n>  %s / %s Open orders (Records length = %d)\r\n",
		strings.ToUpper(ths.Monetary), strings.ToUpper(ths.Coin), length)
	//> Price          Amount       Total             Price          Amount       Total
	//> 0.00001071    868.80058877    0.00930485              0.00001074      15933.88623733  0.17112994
	ret += ">      ************ Buy ************                         ************ Sell ************ \r\n"
	ret += "> Price         Amount          Total                   Price           Amount          Total\r\n"
	indexBuy := 0
	indexSell := 0
	for indexBuy < relLenBuy || indexSell < relLenSell {
		if (indexBuy < relLenBuy) && (indexSell < relLenSell) {
			bItm := buyList[indexBuy]
			sItm := sellList[indexSell]
			ret += fmt.Sprintf("> %.8f\t%.8f\t%.8f\t\t%.8f\t%.8f\t%.8f\r\n",
				bItm.Price, bItm.Amount, bItm.Total,
				sItm.Price, sItm.Amount, sItm.Total)
		} else if (indexBuy >= relLenBuy) && (indexSell < relLenSell) {
			sItm := sellList[indexSell]
			//                  > 0.00001071    8.80058877    0.00930485              0.00001074      15933.88623733  0.17112994
			ret += fmt.Sprintf("> -         \t-         \t-         \t\t%.8f\t%.8f\t%.8f\r\n",
				sItm.Price, sItm.Amount, sItm.Total)
		} else if (indexBuy < relLenBuy) && (indexSell >= relLenSell) {
			bItm := buyList[indexBuy]
			//                  > 0.00001071    8.80058877    0.00930485              0.00001074      15933.88623733  0.17112994
			ret += fmt.Sprintf("> %.8f\t%.8f\t%.8f\t\t-         \t-         \t-\r\n",
				bItm.Price, bItm.Amount, bItm.Total)
		} else {
			break
		}
		indexBuy++
		indexSell++
	}
	return ret
}

// PrintHistorys : print histroy datas to string
func (ths *ReaderDef) PrintHistorys(length int) string {
	ths.historyLock.Lock()
	defer ths.historyLock.Unlock()
	if ths.Datas.Historys == nil {
		return "> No datas!!\r\n"
	}

	relLen := len(ths.Datas.Historys)
	if length != -1 && length < relLen {
		relLen = length
	}

	ret := fmt.Sprintf("\r\n>  %s / %s Trade history datas (Records length = %d)\r\n",
		strings.ToUpper(ths.Monetary), strings.ToUpper(ths.Coin), relLen)
	//> 2016-06-02 09:58:21   buy     0.00001069      187.09073900    0.00199999
	ret += "> DateTime              Type    Price           Amount          Total\r\n"

	for index, his := range ths.Datas.Historys {
		if index < relLen {
			ret += fmt.Sprintf("> %s\t%s\t%.8f\t%.8f\t%.8f\r\n",
				his.DateTime.Format("2006-01-02 15:04:05"),
				his.Type, his.Price, his.Amount, his.Total)
		}
	}
	return ret
}

// SetClient : set http client to rhttp
func (ths *ReaderDef) SetClient(c *rhttp.CHttp) error {
	if ths.Proxy.UseProxy() {
		client, err := c.GetProxyClient(30, ths.Proxy.Address, ths.Proxy.Port)
		if err != nil {
			return err
		}
		c.SetClient(client)
	} else {
		c.SetClient(c.GetClient(30))
	}
	return nil
}
