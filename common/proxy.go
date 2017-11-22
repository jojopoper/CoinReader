package common

// ProxyDef : proxy define
type ProxyDef struct {
	use     bool
	Address string // proxy server address. If not set, proxy is not be used.
	Port    string // proxy server port. If not set, 8181 is default.
}

// GetInitProxy : get init proxy object
func GetInitProxy(v ...interface{}) *ProxyDef {
	ret := &ProxyDef{}

	vLen := len(v)
	if vLen >= 1 {
		switch v[0].(type) {
		case string:
			ret.Address = v[0].(string)
		default:
			panic("The first parameter(Address) must be of type string!")
		}
	}
	if vLen >= 2 {
		switch v[1].(type) {
		case string:
			ret.Port = v[1].(string)
		default:
			panic("The second parameter(Port) must be of type string!")
		}
	}
	if len(ret.Address) > 0 {
		ret.use = true
		if len(ret.Port) == 0 {
			ret.Port = "8181"
		}
	} else {
		ret.use = false
	}
	return ret
}

// UseProxy : Return use or not use a proxy configuration.
func (ths *ProxyDef) UseProxy() bool {
	return ths.use
}
