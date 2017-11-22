package common

// Results : result datas
type Results struct {
	Orders   map[string][]*OrderBook
	Historys []*History
}

// GetInitResults : get init result object
func GetInitResults() *Results {
	ret := &Results{
		Orders:   make(map[string][]*OrderBook),
		Historys: make([]*History, 0),
	}
	ret.ClearHistorys()
	ret.ClearOrderBook()
	return ret
}

// AddOrder : add order book obejct
// key is "OrderBuyStringKey" or "OrderSellStringKey"
func (ths *Results) AddOrder(key string, o *OrderBook) int {
	_, ok := ths.Orders[key]
	if ok {
		ths.Orders[key] = append(ths.Orders[key], o)
		return len(ths.Orders[key])
	}
	return -1
}

// ClearOrderBook : clear order book context
func (ths *Results) ClearOrderBook() {
	ths.Orders[OrderBuyStringKey] = make([]*OrderBook, 0)
	ths.Orders[OrderSellStringKey] = make([]*OrderBook, 0)
}

// AddHistory : add a history object
func (ths *Results) AddHistory(h *History) {
	ths.Historys = append(ths.Historys, h)
}

// ClearHistorys : clear history context
func (ths *Results) ClearHistorys() {
	ths.Historys = make([]*History, 0)
}
