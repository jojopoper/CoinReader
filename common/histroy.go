package common

import (
	"time"
)

// History Trade histroy data struct
type History struct {
	DateTime time.Time
	Type     string
	Price    float64
	Amount   float64
	Total    float64
}
