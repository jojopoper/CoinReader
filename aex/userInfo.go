package aex

import (
	"fmt"
	"sync"
	"time"

	"github.com/jojopoper/CoinReader/common"
	"github.com/jojopoper/CoinReader/common/tool"
	"github.com/jojopoper/go-models/checker"
)

// AexUserInfo : aex user information define
type AexUserInfo struct {
	sync.Mutex
	checker.CheckBase
	tool.CryptoUtils
	info      *common.BaseUserInfo
	tm        int64
	md5Result string
}

// Init : init struct information
func (ths *AexUserInfo) Init(info *common.BaseUserInfo) {
	ths.info = info
	ths.CheckBase.Init(5000)
	ths.SetExeFunc(ths.execute)
	ths.execute()
	ths.Start()
}

// GetReqParams : return http request params
func (ths *AexUserInfo) GetReqParams() string {
	ths.Lock()
	defer ths.Unlock()
	return fmt.Sprintf("key=%s&skey=%s&time=%d&md5=%s", ths.info.Keys[0], ths.info.Keys[1], ths.tm, ths.md5Result)
}

func (ths *AexUserInfo) execute() {
	ths.Lock()
	defer ths.Unlock()
	ths.tm = time.Now().Unix()
	// format is : key_用戶ID_skey_time
	ths.md5Result, _ = ths.GetMd5(fmt.Sprintf("%s_%s_%s_%d", ths.info.Keys[0], ths.info.UserID, ths.info.Keys[1], ths.tm))
}
