package tool

import (
	"reflect"
	"sync"
)

// ReverseSlice : reverse a slice
type ReverseSlice struct {
	sync.Mutex
}

// R : reverse
func (ths *ReverseSlice) R(s interface{}) {
	ths.Lock()
	defer ths.Unlock()
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
