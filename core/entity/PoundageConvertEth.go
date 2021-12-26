package entity

import (
	"github.com/qauzy/math"
	"time"
)

func (this *PoundageConvertEth) SetId(id int64) (result *PoundageConvertEth) {
	this.Id = id
	return this
}
func (this *PoundageConvertEth) GetId() (id int64) {
	return this.Id
}
func (this *PoundageConvertEth) SetExchangeOrderId(exchangeOrderId string) (result *PoundageConvertEth) {
	this.ExchangeOrderId = exchangeOrderId
	return this
}
func (this *PoundageConvertEth) GetExchangeOrderId() (exchangeOrderId string) {
	return this.ExchangeOrderId
}
func (this *PoundageConvertEth) SetMemberId(memberId int64) (result *PoundageConvertEth) {
	this.MemberId = memberId
	return this
}
func (this *PoundageConvertEth) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *PoundageConvertEth) SetMineAmount(mineAmount math.BigDecimal) (result *PoundageConvertEth) {
	this.MineAmount = mineAmount
	return this
}
func (this *PoundageConvertEth) GetMineAmount() (mineAmount math.BigDecimal) {
	return this.MineAmount
}
func (this *PoundageConvertEth) SetPoundageAmount(poundageAmount math.BigDecimal) (result *PoundageConvertEth) {
	this.PoundageAmount = poundageAmount
	return this
}
func (this *PoundageConvertEth) GetPoundageAmount() (poundageAmount math.BigDecimal) {
	return this.PoundageAmount
}
func (this *PoundageConvertEth) SetPoundageAmountEth(poundageAmountEth math.BigDecimal) (result *PoundageConvertEth) {
	this.PoundageAmountEth = poundageAmountEth
	return this
}
func (this *PoundageConvertEth) GetPoundageAmountEth() (poundageAmountEth math.BigDecimal) {
	return this.PoundageAmountEth
}
func (this *PoundageConvertEth) SetCoinId(coinId string) (result *PoundageConvertEth) {
	this.CoinId = coinId
	return this
}
func (this *PoundageConvertEth) GetCoinId() (coinId string) {
	return this.CoinId
}
func (this *PoundageConvertEth) SetTransactionTime(transactionTime time.Time) (result *PoundageConvertEth) {
	this.TransactionTime = transactionTime
	return this
}
func (this *PoundageConvertEth) GetTransactionTime() (transactionTime time.Time) {
	return this.TransactionTime
}
func (this *PoundageConvertEth) SetSymbol(symbol string) (result *PoundageConvertEth) {
	this.Symbol = symbol
	return this
}
func (this *PoundageConvertEth) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *PoundageConvertEth) SetType(oType string) (result *PoundageConvertEth) {
	this.Type = oType
	return this
}
func (this *PoundageConvertEth) GetType() (oType string) {
	return this.Type
}
func (this *PoundageConvertEth) SetDirection(direction string) (result *PoundageConvertEth) {
	this.Direction = direction
	return this
}
func (this *PoundageConvertEth) GetDirection() (direction string) {
	return this.Direction
}
func (this *PoundageConvertEth) SetUsdtRate(usdtRate string) (result *PoundageConvertEth) {
	this.UsdtRate = usdtRate
	return this
}
func (this *PoundageConvertEth) GetUsdtRate() (usdtRate string) {
	return this.UsdtRate
}
func (this *PoundageConvertEth) SetEthUsdtRate(ethUsdtRate string) (result *PoundageConvertEth) {
	this.EthUsdtRate = ethUsdtRate
	return this
}
func (this *PoundageConvertEth) GetEthUsdtRate() (ethUsdtRate string) {
	return this.EthUsdtRate
}

type PoundageConvertEth struct {
	Id                int64
	ExchangeOrderId   string
	MemberId          int64
	MineAmount        math.BigDecimal
	PoundageAmount    math.BigDecimal
	PoundageAmountEth math.BigDecimal
	CoinId            string
	TransactionTime   time.Time
	Symbol            string
	Type              string
	Direction         string
	UsdtRate          string
	EthUsdtRate       string
}
