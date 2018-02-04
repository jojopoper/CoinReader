package rd

import (
	"fmt"
	"strings"

	"github.com/jojopoper/CoinReader/common/cap"
	"github.com/jojopoper/CoinReader/common/rd"
	"github.com/jojopoper/rhttp"
	_L "github.com/jojopoper/xlog"
)

type SymbolType int

const (
	DefaultSymbol  SymbolType = 0
	BcSymbol       SymbolType = 2
	HCashSymbol    SymbolType = 5
	StoxSymbol     SymbolType = 6
	ZrxSymbol      SymbolType = 7
	CdtSymbol      SymbolType = 8
	TntSymbol      SymbolType = 9
	ManaSymbol     SymbolType = 10
	OmgSymbol      SymbolType = 11
	KyberSymbol    SymbolType = 12
	EthSymbol      SymbolType = 14
	AeSymbol       SymbolType = 15
	StreamrSymbol  SymbolType = 16
	HyperpaySymbol SymbolType = 17
	DatumSymbol    SymbolType = 18
	OnerootSymbol  SymbolType = 19
	DewSymbol      SymbolType = 20
	MaggieSymbol   SymbolType = 21
	StorjSymbol    SymbolType = 23
	StatusSymbol   SymbolType = 24
	DogeSymbol     SymbolType = 25
	WiccSymbol     SymbolType = 28
	EosSymbol      SymbolType = 29
	SdaSymbol      SymbolType = 30
	CoinsSymbol    SymbolType = 31
	BdgSymbol      SymbolType = 32
	BeechatSymbol  SymbolType = 33
	RctSymbol      SymbolType = 36
)

// Reader coinw.com reader struct
type Reader struct {
	rd.ReaderDef
	orderClt      *rhttp.CHttp
	orderParams   string
	historyClt    *rhttp.CHttp
	historyParams string
}

// Init init parameters
func (ths *Reader) Init(m, c *cap.CoinCapacity, v ...interface{}) {
	ths.ReaderDef.Init(m, c, ths.rdOrders, ths.rdHistory, v...)
	ths.baseInit()
	ths.initOrderParams()
	ths.initHistParams()
}

func (ths *Reader) baseInit() {
	ths.Monetary.Name = strings.ToLower(ths.Monetary.Name)
	ths.Coin.Name = strings.ToLower(ths.Coin.Name)
	ths.Address = "www.coinw.com"

	ths.OrderAddr = fmt.Sprintf("https://%s/appApi.html", ths.Address)
	ths.orderParams = fmt.Sprintf("action=depth&symbol=%d", ths.getSymbol(ths.Coin.Name))

	ths.HistoryAddr = fmt.Sprintf("https://%s/appApi.html", ths.Address)
	ths.historyParams = fmt.Sprintf("action=trades&symbol=%d", ths.getSymbol(ths.Coin.Name))
}

func (ths *Reader) getSymbol(cn string) SymbolType {
	switch cn {
	case "eth":
		return EthSymbol
	case "coins":
		return CoinsSymbol
	case "hcash", "hsr":
		return HCashSymbol
	case "sda":
		return SdaSymbol
	case "bdg":
		return BdgSymbol
	case "beechat", "chat":
		return BeechatSymbol
	case "status", "snt":
		return StatusSymbol
	case "Storj":
		return StorjSymbol
	case "rct":
		return RctSymbol
	case "wicc":
		return WiccSymbol
	case "eos":
		return EosSymbol
	case "doge":
		return DogeSymbol
	case "dew":
		return DewSymbol
	case "maggie", "mag":
		return MaggieSymbol
	case "oneroot", "rnt":
		return OnerootSymbol
	case "datum", "dat":
		return DatumSymbol
	case "ae":
		return AeSymbol
	case "hyperpay", "hpy":
		return HyperpaySymbol
	case "streamr", "data":
		return StreamrSymbol
	case "omg":
		return OmgSymbol
	case "kyber":
		return KyberSymbol
	case "zrx", "0x":
		return ZrxSymbol
	case "tnt":
		return TntSymbol
	case "stox", "stx":
		return StoxSymbol
	case "cdt":
		return CdtSymbol
	case "bc":
		return BcSymbol
	default:
		return DefaultSymbol
	}
}

func (ths *Reader) initOrderParams() {
	ths.orderClt = new(rhttp.CHttp)
	ths.orderClt.SetDecodeFunc(ths.decodeOrders)
	if err := ths.SetClient(ths.orderClt); err != nil {
		_L.Error("Coinw : InitOrderParams set http client has error \n%+v", err)
	}
}

func (ths *Reader) initHistParams() {
	ths.historyClt = new(rhttp.CHttp)
	ths.historyClt.SetDecodeFunc(ths.decodeHistory)
	if err := ths.SetClient(ths.historyClt); err != nil {
		_L.Error("Coinw : initHistParams set http client has error \n%+v", err)
	}
}
