package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *PoundageConvertEth) SetId(Id int64) (result *PoundageConvertEth) {
	this.Id = Id
	return this
}
func (this *PoundageConvertEth) GetId() (Id int64) {
	return this.Id
}
func (this *PoundageConvertEth) SetExchangeOrderId(ExchangeOrderId string) (result *PoundageConvertEth) {
	this.ExchangeOrderId = ExchangeOrderId
	return this
}
func (this *PoundageConvertEth) GetExchangeOrderId() (ExchangeOrderId string) {
	return this.ExchangeOrderId
}
func (this *PoundageConvertEth) SetMemberId(MemberId int64) (result *PoundageConvertEth) {
	this.MemberId = MemberId
	return this
}
func (this *PoundageConvertEth) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *PoundageConvertEth) SetMineAmount(MineAmount math.BigDecimal) (result *PoundageConvertEth) {
	this.MineAmount = MineAmount
	return this
}
func (this *PoundageConvertEth) GetMineAmount() (MineAmount math.BigDecimal) {
	return this.MineAmount
}
func (this *PoundageConvertEth) SetPoundageAmount(PoundageAmount math.BigDecimal) (result *PoundageConvertEth) {
	this.PoundageAmount = PoundageAmount
	return this
}
func (this *PoundageConvertEth) GetPoundageAmount() (PoundageAmount math.BigDecimal) {
	return this.PoundageAmount
}
func (this *PoundageConvertEth) SetPoundageAmountEth(PoundageAmountEth math.BigDecimal) (result *PoundageConvertEth) {
	this.PoundageAmountEth = PoundageAmountEth
	return this
}
func (this *PoundageConvertEth) GetPoundageAmountEth() (PoundageAmountEth math.BigDecimal) {
	return this.PoundageAmountEth
}
func (this *PoundageConvertEth) SetCoinId(CoinId string) (result *PoundageConvertEth) {
	this.CoinId = CoinId
	return this
}
func (this *PoundageConvertEth) GetCoinId() (CoinId string) {
	return this.CoinId
}
func (this *PoundageConvertEth) SetTransactionTime(TransactionTime xtime.Xtime) (result *PoundageConvertEth) {
	this.TransactionTime = TransactionTime
	return this
}
func (this *PoundageConvertEth) GetTransactionTime() (TransactionTime xtime.Xtime) {
	return this.TransactionTime
}
func (this *PoundageConvertEth) SetSymbol(Symbol string) (result *PoundageConvertEth) {
	this.Symbol = Symbol
	return this
}
func (this *PoundageConvertEth) GetSymbol() (Symbol string) {
	return this.Symbol
}
func (this *PoundageConvertEth) SetType(Type string) (result *PoundageConvertEth) {
	this.Type = Type
	return this
}
func (this *PoundageConvertEth) GetType() (Type string) {
	return this.Type
}
func (this *PoundageConvertEth) SetDirection(Direction string) (result *PoundageConvertEth) {
	this.Direction = Direction
	return this
}
func (this *PoundageConvertEth) GetDirection() (Direction string) {
	return this.Direction
}
func (this *PoundageConvertEth) SetUsdtRate(UsdtRate string) (result *PoundageConvertEth) {
	this.UsdtRate = UsdtRate
	return this
}
func (this *PoundageConvertEth) GetUsdtRate() (UsdtRate string) {
	return this.UsdtRate
}
func (this *PoundageConvertEth) SetEthUsdtRate(EthUsdtRate string) (result *PoundageConvertEth) {
	this.EthUsdtRate = EthUsdtRate
	return this
}
func (this *PoundageConvertEth) GetEthUsdtRate() (EthUsdtRate string) {
	return this.EthUsdtRate
}
func NewPoundageConvertEth(id int64, exchangeOrderId string, memberId int64, mineAmount math.BigDecimal, poundageAmount math.BigDecimal, poundageAmountEth math.BigDecimal, coinId string, transactionTime xtime.Xtime, symbol string, oType string, direction string, usdtRate string, ethUsdtRate string) (ret *PoundageConvertEth) {
	ret = new(PoundageConvertEth)
	ret.Id = id
	ret.ExchangeOrderId = exchangeOrderId
	ret.MemberId = memberId
	ret.MineAmount = mineAmount
	ret.PoundageAmount = poundageAmount
	ret.PoundageAmountEth = poundageAmountEth
	ret.CoinId = coinId
	ret.TransactionTime = transactionTime
	ret.Symbol = symbol
	ret.Type = oType
	ret.Direction = direction
	ret.UsdtRate = usdtRate
	ret.EthUsdtRate = ethUsdtRate
	return
}

type PoundageConvertEth struct {
	Id                int64           `gorm:"column:id" json:"id"`
	ExchangeOrderId   string          `gorm:"column:exchange_order_id" json:"exchangeOrderId"`
	MemberId          int64           `gorm:"column:member_id" json:"memberId"`
	MineAmount        math.BigDecimal `gorm:"column:mine_amount" json:"mineAmount"`
	PoundageAmount    math.BigDecimal `gorm:"column:poundage_amount" json:"poundageAmount"`
	PoundageAmountEth math.BigDecimal `gorm:"column:poundage_amount_eth" json:"poundageAmountEth"`
	CoinId            string          `gorm:"column:coin_id" json:"coinId"`
	TransactionTime   xtime.Xtime     `gorm:"column:transaction_time" json:"transactionTime"`
	Symbol            string          `gorm:"column:symbol" json:"symbol"`
	Type              string          `gorm:"column:type" json:"type"`
	Direction         string          `gorm:"column:direction" json:"direction"`
	UsdtRate          string          `gorm:"column:usdt_rate" json:"usdtRate"`
	EthUsdtRate       string          `gorm:"column:eth_usdt_rate" json:"ethUsdtRate"`
}
