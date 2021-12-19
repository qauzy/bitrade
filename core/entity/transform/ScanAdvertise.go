package transform

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
	"time"
)

func (this *ScanAdvertise) SetMemberName(memberName string) (result *ScanAdvertise) {
	this.MemberName = memberName
	return this
}
func (this *ScanAdvertise) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *ScanAdvertise) SetAvatar(avatar string) (result *ScanAdvertise) {
	this.Avatar = avatar
	return this
}
func (this *ScanAdvertise) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *ScanAdvertise) SetAdvertiseId(advertiseId int64) (result *ScanAdvertise) {
	this.AdvertiseId = advertiseId
	return this
}
func (this *ScanAdvertise) GetAdvertiseId() (advertiseId int64) {
	return this.AdvertiseId
}
func (this *ScanAdvertise) SetTransactions(transactions int) (result *ScanAdvertise) {
	this.Transactions = transactions
	return this
}
func (this *ScanAdvertise) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *ScanAdvertise) SetPrice(price math.BigDecimal) (result *ScanAdvertise) {
	this.Price = price
	return this
}
func (this *ScanAdvertise) GetPrice() (price math.BigDecimal) {
	return this.Price
}
func (this *ScanAdvertise) SetMinLimit(minLimit math.BigDecimal) (result *ScanAdvertise) {
	this.MinLimit = minLimit
	return this
}
func (this *ScanAdvertise) GetMinLimit() (minLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *ScanAdvertise) SetMaxLimit(maxLimit math.BigDecimal) (result *ScanAdvertise) {
	this.MaxLimit = maxLimit
	return this
}
func (this *ScanAdvertise) GetMaxLimit() (maxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *ScanAdvertise) SetRemainAmount(remainAmount math.BigDecimal) (result *ScanAdvertise) {
	this.RemainAmount = remainAmount
	return this
}
func (this *ScanAdvertise) GetRemainAmount() (remainAmount math.BigDecimal) {
	return this.RemainAmount
}
func (this *ScanAdvertise) SetCreateTime(createTime time.Time) (result *ScanAdvertise) {
	this.CreateTime = createTime
	return this
}
func (this *ScanAdvertise) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *ScanAdvertise) SetPayMode(payMode string) (result *ScanAdvertise) {
	this.PayMode = payMode
	return this
}
func (this *ScanAdvertise) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *ScanAdvertise) SetCoinId(coinId int64) (result *ScanAdvertise) {
	this.CoinId = coinId
	return this
}
func (this *ScanAdvertise) GetCoinId() (coinId int64) {
	return this.CoinId
}
func (this *ScanAdvertise) SetUnit(unit string) (result *ScanAdvertise) {
	this.Unit = unit
	return this
}
func (this *ScanAdvertise) GetUnit() (unit string) {
	return this.Unit
}
func (this *ScanAdvertise) SetCoinName(coinName string) (result *ScanAdvertise) {
	this.CoinName = coinName
	return this
}
func (this *ScanAdvertise) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *ScanAdvertise) SetCoinNameCn(coinNameCn string) (result *ScanAdvertise) {
	this.CoinNameCn = coinNameCn
	return this
}
func (this *ScanAdvertise) GetCoinNameCn() (coinNameCn string) {
	return this.CoinNameCn
}
func (this *ScanAdvertise) SetLevel(level int) (result *ScanAdvertise) {
	this.Level = level
	return this
}
func (this *ScanAdvertise) GetLevel() (level int) {
	return this.Level
}
func (this *ScanAdvertise) SetAdvertiseType(advertiseType constant.AdvertiseType) (result *ScanAdvertise) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *ScanAdvertise) GetAdvertiseType() (advertiseType constant.AdvertiseType) {
	return this.AdvertiseType
}
func (this *ScanAdvertise) SetAdvType(advType int) (result *ScanAdvertise) {
	this.AdvType = advType
	return this
}
func (this *ScanAdvertise) GetAdvType() (advType int) {
	return this.AdvType
}
func (this *ScanAdvertise) SetPremiseRate(premiseRate math.BigDecimal) (result *ScanAdvertise) {
	this.PremiseRate = premiseRate
	return this
}
func (this *ScanAdvertise) GetPremiseRate() (premiseRate math.BigDecimal) {
	return this.PremiseRate
}

type ScanAdvertise struct {
	MemberName    string
	Avatar        string
	AdvertiseId   int64
	Transactions  int
	Price         math.BigDecimal
	MinLimit      math.BigDecimal
	MaxLimit      math.BigDecimal
	RemainAmount  math.BigDecimal
	CreateTime    time.Time
	PayMode       string
	CoinId        int64
	Unit          string
	CoinName      string
	CoinNameCn    string
	Level         int
	AdvertiseType constant.AdvertiseType
	AdvType       int
	PremiseRate   math.BigDecimal
}
