package entity

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
	"time"
)

func (this *Advertise) SetId(id int64) (result *Advertise) {
	this.Id = id
	return this
}
func (this *Advertise) GetId() (id int64) {
	return this.Id
}
func (this *Advertise) SetPrice(price math.BigDecimal) (result *Advertise) {
	this.Price = price
	return this
}
func (this *Advertise) GetPrice() (price math.BigDecimal) {
	return this.Price
}
func (this *Advertise) SetAdvertiseType(advertiseType constant.AdvertiseType) (result *Advertise) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *Advertise) GetAdvertiseType() (advertiseType constant.AdvertiseType) {
	return this.AdvertiseType
}
func (this *Advertise) SetMember(member Member) (result *Advertise) {
	this.Member = member
	return this
}
func (this *Advertise) GetMember() (member Member) {
	return this.Member
}
func (this *Advertise) SetCreateTime(createTime time.Time) (result *Advertise) {
	this.CreateTime = createTime
	return this
}
func (this *Advertise) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Advertise) SetUpdateTime(updateTime time.Time) (result *Advertise) {
	this.UpdateTime = updateTime
	return this
}
func (this *Advertise) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *Advertise) SetCoin(coin *OtcCoin) (result *Advertise) {
	this.Coin = coin
	return this
}
func (this *Advertise) GetCoin() (coin *OtcCoin) {
	return this.Coin
}
func (this *Advertise) SetCountry(country Country) (result *Advertise) {
	this.Country = country
	return this
}
func (this *Advertise) GetCountry() (country Country) {
	return this.Country
}
func (this *Advertise) SetMinLimit(minLimit math.BigDecimal) (result *Advertise) {
	this.MinLimit = minLimit
	return this
}
func (this *Advertise) GetMinLimit() (minLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *Advertise) SetMaxLimit(maxLimit math.BigDecimal) (result *Advertise) {
	this.MaxLimit = maxLimit
	return this
}
func (this *Advertise) GetMaxLimit() (maxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *Advertise) SetRemark(remark string) (result *Advertise) {
	this.Remark = remark
	return this
}
func (this *Advertise) GetRemark() (remark string) {
	return this.Remark
}
func (this *Advertise) SetTimeLimit(timeLimit int64) (result *Advertise) {
	this.TimeLimit = timeLimit
	return this
}
func (this *Advertise) GetTimeLimit() (timeLimit int64) {
	return this.TimeLimit
}
func (this *Advertise) SetPremiseRate(premiseRate math.BigDecimal) (result *Advertise) {
	this.PremiseRate = premiseRate
	return this
}
func (this *Advertise) GetPremiseRate() (premiseRate math.BigDecimal) {
	return this.PremiseRate
}
func (this *Advertise) SetLevel(level constant.AdvertiseLevel) (result *Advertise) {
	this.Level = level
	return this
}
func (this *Advertise) GetLevel() (level constant.AdvertiseLevel) {
	return this.Level
}
func (this *Advertise) SetPayMode(payMode string) (result *Advertise) {
	this.PayMode = payMode
	return this
}
func (this *Advertise) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *Advertise) SetVersion(version int64) (result *Advertise) {
	this.Version = version
	return this
}
func (this *Advertise) GetVersion() (version int64) {
	return this.Version
}
func (this *Advertise) SetStatus(status constant.AdvertiseControlStatus) (result *Advertise) {
	this.Status = status
	return this
}
func (this *Advertise) GetStatus() (status constant.AdvertiseControlStatus) {
	return this.Status
}
func (this *Advertise) SetPriceType(priceType constant.PriceType) (result *Advertise) {
	this.PriceType = priceType
	return this
}
func (this *Advertise) GetPriceType() (priceType constant.PriceType) {
	return this.PriceType
}
func (this *Advertise) SetNumber(number math.BigDecimal) (result *Advertise) {
	this.Number = number
	return this
}
func (this *Advertise) GetNumber() (number math.BigDecimal) {
	return this.Number
}
func (this *Advertise) SetDealAmount(dealAmount math.BigDecimal) (result *Advertise) {
	this.DealAmount = dealAmount
	return this
}
func (this *Advertise) GetDealAmount() (dealAmount math.BigDecimal) {
	return this.DealAmount
}
func (this *Advertise) SetRemainAmount(remainAmount math.BigDecimal) (result *Advertise) {
	this.RemainAmount = remainAmount
	return this
}
func (this *Advertise) GetRemainAmount() (remainAmount math.BigDecimal) {
	return this.RemainAmount
}
func (this *Advertise) SetAuto(auto constant.BooleanEnum) (result *Advertise) {
	this.Auto = auto
	return this
}
func (this *Advertise) GetAuto() (auto constant.BooleanEnum) {
	return this.Auto
}
func (this *Advertise) SetMarketPrice(marketPrice math.BigDecimal) (result *Advertise) {
	this.MarketPrice = marketPrice
	return this
}
func (this *Advertise) GetMarketPrice() (marketPrice math.BigDecimal) {
	return this.MarketPrice
}
func (this *Advertise) SetAutoword(autoword string) (result *Advertise) {
	this.Autoword = autoword
	return this
}
func (this *Advertise) GetAutoword() (autoword string) {
	return this.Autoword
}
func (this *Advertise) SetLimitMoney(limitMoney string) (result *Advertise) {
	this.LimitMoney = limitMoney
	return this
}
func (this *Advertise) GetLimitMoney() (limitMoney string) {
	return this.LimitMoney
}
func (this *Advertise) SetUsername(username string) (result *Advertise) {
	this.Username = username
	return this
}
func (this *Advertise) GetUsername() (username string) {
	return this.Username
}
func (this *Advertise) SetCoinUnit(coinUnit string) (result *Advertise) {
	this.CoinUnit = coinUnit
	return this
}
func (this *Advertise) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}

type Advertise struct {
	Id            int64
	Price         math.BigDecimal
	AdvertiseType constant.AdvertiseType
	Member        Member
	CreateTime    time.Time
	UpdateTime    time.Time
	Coin          *OtcCoin
	Country       Country
	MinLimit      math.BigDecimal
	MaxLimit      math.BigDecimal
	Remark        string
	TimeLimit     int64
	PremiseRate   math.BigDecimal
	Level         constant.AdvertiseLevel
	PayMode       string
	Version       int64
	Status        constant.AdvertiseControlStatus
	PriceType     constant.PriceType
	Number        math.BigDecimal
	DealAmount    math.BigDecimal
	RemainAmount  math.BigDecimal
	Auto          constant.BooleanEnum
	MarketPrice   math.BigDecimal
	Autoword      string
	LimitMoney    string
	Username      string
	CoinUnit      string
}
