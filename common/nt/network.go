package nt

import (
	"github.com/jojopoper/rhttp"
)

// NetworkClient : net work client function define
type NetworkClient struct {
	Proxy *ProxyDef
}

// SetClient : set http client to rhttp
func (ths *NetworkClient) SetClient(c *rhttp.CHttp) error {
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
