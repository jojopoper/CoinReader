package cap

import (
	"fmt"
)

// CoinCapacity : capacity of coin
// only for float
type CoinCapacity struct {
	Name      string
	Decimal   int
	Min       float64
	decFormat string
}

// Set : set coin name and decimal and order min value
func (ths *CoinCapacity) Set(n string, d int, m float64) {
	ths.Name = n
	ths.Decimal = d
	ths.decFormat = fmt.Sprintf("%%.%df", ths.Decimal)
	ths.Min = m
}

// Format : return decimal string format
func (ths *CoinCapacity) Format() string {
	return ths.decFormat
}
