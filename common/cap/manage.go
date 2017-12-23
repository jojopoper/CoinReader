package cap

// Management : manage capacity define
type Management struct {
	Items map[string]*CoinCapacity
}

func (ths *Management) set(itm *CoinCapacity) {
	ths.Items[itm.Name] = itm
}

func (ths *Management) get(n string) *CoinCapacity {
	return ths.Items[n]
}

// AddDetail : detail information add capacity object
func (ths *Management) AddDetail(n string, d int, m float64) {
	itm := &CoinCapacity{}
	itm.Set(n, d, m)
	ths.set(itm)
}

// AddObject : object add capacity
func (ths *Management) AddObject(itm *CoinCapacity) {
	ths.set(itm)
}

// Format : return capacity format
func (ths *Management) Format(n string) string {
	return ths.get(n).Format()
}

// LowThreshold : return v is upper than min value of capacity
func (ths *Management) LowThreshold(n string, v float64) bool {
	return ths.get(n).Min < v
}

// GetMin : min value of capacity
func (ths *Management) GetMin(n string) float64 {
	return ths.get(n).Min
}

// GetObject : get coin capacity
func (ths *Management) GetObject(n string) *CoinCapacity {
	return ths.get(n)
}
