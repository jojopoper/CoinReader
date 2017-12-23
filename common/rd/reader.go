package rd

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/nt"
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
	nt.NetworkClient
	Address      string
	Monetary     *cap.CoinCapacity
	Coin         *cap.CoinCapacity
	outFormat    string
	Datas        *Results
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
func (ths *ReaderDef) Init(m, c *cap.CoinCapacity, rof, rhf ReadFunc, v ...interface{}) {
	ths.Monetary = m
	ths.Coin = c
	ths.outFormat = fmt.Sprintf("%s\t%s\t%s", ths.Monetary.Format(), ths.Coin.Format(), ths.Monetary.Format())
	ths.orderLock = new(sync.Mutex)
	ths.historyLock = new(sync.Mutex)
	ths.SetReaderFunc(rof, rhf)

	ths.Datas = GetInitResults()
	ths.Proxy = nt.GetInitProxy(v...)
}

// ReadAll : readout order and history datas from website
func (ths *ReaderDef) ReadAll() bool {
	return ths.readHistorys() && ths.readOrders()
}

// SetReaderFunc : set read order and history function
func (ths *ReaderDef) SetReaderFunc(rof ReadFunc, rhf ReadFunc) {
	ths.readOrders = rof
	ths.readHistorys = rhf
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
		strings.ToUpper(ths.Monetary.Name), strings.ToUpper(ths.Coin.Name), length)
	//> Price          Amount       Total             Price          Amount       Total
	//> 0.00001071    868.80058877    0.00930485              0.00001074      15933.88623733  0.17112994
	ret += ">      ************ Buy ************                         ************ Sell ************ \r\n"
	ret += "> Price         Amount          Total                   Price           Amount          Total\r\n"
	indexBuy := 0
	indexSell := 0
	format := ""
	for indexBuy < relLenBuy || indexSell < relLenSell {
		if (indexBuy < relLenBuy) && (indexSell < relLenSell) {
			bItm := buyList[indexBuy]
			sItm := sellList[indexSell]
			format = fmt.Sprintf("> %s\t\t%s\n", ths.outFormat, ths.outFormat)
			ret += fmt.Sprintf(format, bItm.Price, bItm.Amount, bItm.Total,
				sItm.Price, sItm.Amount, sItm.Total)
		} else if (indexBuy >= relLenBuy) && (indexSell < relLenSell) {
			sItm := sellList[indexSell]
			//                  > 0.00001071    8.80058877    0.00930485              0.00001074      15933.88623733  0.17112994
			format = fmt.Sprintf("> -         \t-         \t-         \t\t%s\n", ths.outFormat)
			ret += fmt.Sprintf(format, sItm.Price, sItm.Amount, sItm.Total)
		} else if (indexBuy < relLenBuy) && (indexSell >= relLenSell) {
			bItm := buyList[indexBuy]
			//                  > 0.00001071    8.80058877    0.00930485              0.00001074      15933.88623733  0.17112994
			format = fmt.Sprintf("> %s\t\t-         \t-         \t-\n", ths.outFormat)
			ret += fmt.Sprintf(format, bItm.Price, bItm.Amount, bItm.Total)
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

	ret := fmt.Sprintf("\n>  %s / %s Trade history datas (Records length = %d)\r\n",
		strings.ToUpper(ths.Monetary.Name), strings.ToUpper(ths.Coin.Name), relLen)
	//> 2016-06-02 09:58:21   buy     0.00001069      187.09073900    0.00199999
	ret += "> DateTime              Type    Price           Amount          Total\r\n"
	format := fmt.Sprintf("> %%s\t%%s\t%s\n", ths.outFormat)
	for index, his := range ths.Datas.Historys {
		if index < relLen {
			ret += fmt.Sprintf(format, his.DateTime.Format("2006-01-02 15:04:05"),
				his.Type, his.Price, his.Amount, his.Total)
		}
	}
	return ret
}
