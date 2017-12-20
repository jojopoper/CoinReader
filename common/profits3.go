package common

import (
	"fmt"
	"math"
	"sync"
)

// OrderElement : profit calculation element define
type OrderElement struct {
	OrderBook
	Label string
	Fee   float64
}

// PrintString : print data to string
func (ths *OrderElement) PrintString() string {
	ret := fmt.Sprintf("> Price\t\tAmount\t\tTotal\n")
	ret += fmt.Sprintf("> %.8f\t\t%.8f\t\t%.8f\n", ths.Price, ths.Amount, ths.Total)
	return ret
}

// GainResult : gain result define
type GainResult struct {
	Span    float64
	Get     float64
	Percent float64
	Profit  float64
	From    *OrderElement
	To      *OrderElement
	Agent   *OrderElement
}

// PrintString : print results to string
func (ths *GainResult) PrintString() string {
	ret := fmt.Sprintf("\n> From (%s)  Order (Fee = %.3f%%)\n", ths.From.Label, ths.From.Fee*100.0)
	ret += ths.From.PrintString() + "\n"
	ret += fmt.Sprintf("> To(%s)  Order (Fee = %.3f%%)\n", ths.To.Label, ths.To.Fee*100.0)
	ret += ths.To.PrintString() + "\n"
	ret += fmt.Sprintf("> Agent(%s) Order (Fee = %.3f%%)\n", ths.Agent.Label, ths.Agent.Fee*100.0)
	ret += ths.Agent.PrintString() + "\n"
	ret += fmt.Sprintf("> Span    = %.8f\n", ths.Span)
	ret += fmt.Sprintf("> Get     = %.8f\n", ths.Get)
	ret += fmt.Sprintf("> * Percent = %.8f%% *\n", ths.Percent)
	ret += fmt.Sprintf("> * Profit  = %.8f *\n", ths.Profit)
	return ret
}

// Profits3Element : 3 elements profits define
type Profits3Element struct {
	sync.Mutex
	from  *OrderElement
	to    *OrderElement
	agent *OrderElement
}

// NewElement : new OrderElement object
// fee : is exchange fee , exsample(0.25% fee=0.0025)
func (ths *Profits3Element) NewElement(s *OrderBook, fee float64, lab string) *OrderElement {
	return &OrderElement{
		OrderBook: *s,
		Label:     lab,
		Fee:       fee,
	}
}

// Reset : reset infor
// f is from; t is to; a is agent
func (ths *Profits3Element) reset(f, t, a *OrderElement) *Profits3Element {
	ths.from = f
	ths.to = t
	ths.agent = a
	return ths
}

// DcEcBcGain :  Dc - Digital currency Ec - exchange currency Bc - Baisc currency
// From a "digital currency" (using the "basic currency") to buy the money through
// the "exchange of money" to calculate the gain of the basic currency.
// For example, buy eth from BTC, sell eth to CNY, and then buy BTC with CNY.
func (ths *Profits3Element) DcEcBcGain(f, t, a *OrderElement) *GainResult {
	ths.Lock()
	defer ths.Unlock()
	ths.reset(f, t, a)

	ret := &GainResult{}
	// 第一步计算Dc挂单
	// 1.1 先得到Dc与Ec量的最小值 - 因为买入当前Dc后需要到Ec去卖出，所以需要计算最小能在Dc买入多少量
	minAmount := math.Min(ths.from.Amount, ths.to.Amount/(1.0-ths.from.Fee))
	// 1.2 计算需要花销多少Bc
	span := minAmount * ths.from.Price
	// 1.3 计算获得Dc的挂单数据
	ret.From = &OrderElement{
		OrderBook: OrderBook{
			Price:  ths.from.Price,
			Amount: minAmount,
			Total:  span,
		},
		Label: ths.from.Label,
		Fee:   f.Fee,
	}
	// 1.4 计算实际获得多少Dc（因为有交易手续费）
	getDcRealAmount := minAmount * (1.0 - ths.from.Fee)
	// 第二步 计算Ec挂单
	// 2.1 计算获得的Dc能拿到多少Ec（含交易手续费之后的值）
	getEcTotal := getDcRealAmount * ths.to.Price
	// 2.2 计算Ec总量是否超出Bc实际挂单总量
	if getEcTotal > ths.agent.Total {
		getEcTotal = ths.agent.Total
		getDcRealAmount = getEcTotal / ths.to.Price
		minAmount = getDcRealAmount / (1.0 - ths.from.Fee)
		span = minAmount * ths.from.Price
		ret.From.Amount = minAmount
		ret.From.Total = span
	}
	// 2.3 得到Ec挂单数据
	ret.To = &OrderElement{
		OrderBook: OrderBook{
			Price:  ths.to.Price,
			Amount: getDcRealAmount,
			Total:  getEcTotal,
		},
		Label: ths.to.Label,
		Fee:   ths.to.Fee,
	}
	// 2.4 计算实际获得Ec总量(除去手续费之后的总量)
	getEcRealTotal := getEcTotal * (1.0 - ths.to.Fee)
	// 第三步计算Bc挂单
	// 3.1 计算能够获得多少Bc（含交易手续费之后的值）
	result := getEcRealTotal / ths.agent.Price
	// 3.2 计算实际到手的Bc结果
	realResult := result * (1.0 - ths.agent.Fee)
	// 3.3 计算第二步花销Bc和最终获取Bc的差值
	profit := realResult - span
	// 3.3 以第二步花销为基础计算盈亏比例
	percent := profit / span * 100.0
	// 3.4 计算Bc挂单数据
	ret.Agent = &OrderElement{
		OrderBook: OrderBook{
			Price:  ths.agent.Price,
			Amount: result,
			Total:  getEcRealTotal,
		},
		Label: ths.agent.Label,
		Fee:   ths.agent.Fee,
	}
	ret.Span = span
	ret.Get = realResult
	ret.Percent = percent
	ret.Profit = profit
	return ret
}

// BcEcDcGain :  Bc - Baisc currency Ec - exchange currency Dc - Digital currency
// From a "digital currency" (using the "basic currency") to buy the money through
// the "exchange of money" to calculate the gain of the basic currency.
// For example, sell BTC to CNY, buy bts from cny, sell bts to BTC.
func (ths *Profits3Element) BcEcDcGain(f, t, a *OrderElement) *GainResult {
	ths.Lock()
	defer ths.Unlock()
	ths.reset(f, t, a)

	ret := &GainResult{}
	// 第一步 计算Bc挂单
	// 1.1 先得到Bc与Ec总量的最小值 - 因为卖出当前Bc后需要到Ec去买入，所以需要计算最小能在Bc买入多少量
	minTotal := math.Min(ths.from.Total, ths.to.Total/(1.0-ths.from.Fee))
	// 1.2 计算最小的挂单量
	minAmount := minTotal / ths.from.Price
	// 1.3 计算获得Bc的挂单数据
	ret.From = &OrderElement{
		OrderBook: OrderBook{
			Price:  ths.from.Price,
			Amount: minAmount,
			Total:  minTotal,
		},
		Label: ths.from.Label,
		Fee:   f.Fee,
	}
	// 1.4 计算实际获得多少Bc（因为有交易手续费）
	getRealTotal := minTotal * (1.0 - ths.from.Fee)
	// 第二步 计算Ec挂单
	// 2.1 计算获得的Bc能拿到多少Ec（未刨除交易手续费之后的值）
	getEcAmount := getRealTotal / ths.to.Price
	// 2.2 计算Ec总量是否超出Dc实际挂单总量
	if getEcAmount > ths.agent.Amount {
		getEcAmount = ths.agent.Amount
		getRealTotal = getEcAmount * ths.to.Price
		minTotal = getRealTotal / (1.0 - ths.from.Fee)
		minAmount = minTotal / ths.from.Price
		ret.From.Amount = minAmount
		ret.From.Total = minTotal
	}
	// 2.3 得到Ec挂单数据
	ret.To = &OrderElement{
		OrderBook: OrderBook{
			Price:  ths.to.Price,
			Amount: getEcAmount,
			Total:  getRealTotal,
		},
		Label: ths.to.Label,
		Fee:   ths.to.Fee,
	}
	// 2.4 计算获得的实际的Ec数量
	getEcRealAmount := getEcAmount * (1.0 - ths.to.Fee)
	// 第三步 计算Dc挂单
	// 3.1 计算能够获得多少Bc（含交易手续费之后的值）
	result := getEcRealAmount * ths.agent.Price
	// 3.2 计算实际获得多少Bc
	realResult := result * (1.0 - ths.agent.Fee)
	// 3.3 计算第二步花销Bc和最终获取Bc的差值
	profit := realResult - minAmount
	// 3.4 以第二步花销为基础计算盈亏比例
	percent := profit / minAmount * 100.0
	// 3.5 计算Bc挂单数据
	ret.Agent = &OrderElement{
		OrderBook: OrderBook{
			Price:  ths.agent.Price,
			Amount: getEcRealAmount,
			Total:  result,
		},
		Label: ths.agent.Label,
		Fee:   ths.agent.Fee,
	}
	ret.Span = minAmount
	ret.Get = realResult
	ret.Percent = percent
	ret.Profit = profit
	return ret
}
