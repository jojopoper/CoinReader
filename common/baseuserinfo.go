package common

// BaseUserInfo : account base information define
type BaseUserInfo struct {
	UserID  string
	Keys    []string
	Address string
}

// Init : init base user infor
func (ths *BaseUserInfo) Init(id string, ks ...string) {
	ths.UserID = id
	ths.Keys = ks
}
