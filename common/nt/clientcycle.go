package nt

import (
	"net/http"
	"sync"

	_r "github.com/jojopoper/rhttp"
)

// ClientCycle : http client cycle list define
type ClientCycle struct {
	sync.Mutex
	size   int
	first  *ClientItem
	cursor *ClientItem
}

// Make : make number of http clients cycle
func (ths *ClientCycle) Make(cnt int) {
	chttp := new(_r.CHttp)
	ths.first = new(ClientItem)
	ths.cursor = ths.first
	curr := ths.first
	for i := 0; i < cnt; i++ {
		// curr.Index = i // for debug
		if i == cnt-1 {
			curr.Set(chttp.GetClient(30), ths.first)
		} else {
			curr.Set(chttp.GetClient(30), new(ClientItem))
			curr = curr.Next()
		}
	}
}

// MakeProxy : make number of http proxy client cycle
func (ths *ClientCycle) MakeProxy(cnt int, ip, port string) error {
	chttp := new(_r.CHttp)
	ths.first = new(ClientItem)
	ths.cursor = ths.first
	curr := ths.first
	for i := 0; i < cnt; i++ {
		tmpClient, err := chttp.GetProxyClient(30, ip, port)
		if err != nil {
			return err
		}
		if i == cnt-1 {
			curr.Set(tmpClient, ths.first)
		} else {
			curr = curr.Set(tmpClient, new(ClientItem)).Next()
		}
	}
	return nil
}

// Next : get next client point
func (ths *ClientCycle) Next() *ClientItem {
	ths.Lock()
	defer ths.Unlock()
	ths.cursor = ths.cursor.Next()
	return ths.cursor
}

// First : get first client point
func (ths *ClientCycle) First() *ClientItem {
	return ths.first
}

// IsFirst : is first ClientItem
func (ths *ClientCycle) IsFirst(c *ClientItem) bool {
	return c == ths.first
}

// Size : http client cycle list size
func (ths *ClientCycle) Size() int {
	return ths.size
}

// ClientItem : http client list cycle item define
type ClientItem struct {
	c    *http.Client
	next *ClientItem
	// Index int // for debug
}

// Get : get current http client
func (ths *ClientItem) Get() *http.Client {
	return ths.c
}

// Next : get next http client
func (ths *ClientItem) Next() *ClientItem {
	return ths.next
}

// Set : set http client
func (ths *ClientItem) Set(c *http.Client, n *ClientItem) *ClientItem {
	ths.c = c
	ths.next = n
	return ths
}
