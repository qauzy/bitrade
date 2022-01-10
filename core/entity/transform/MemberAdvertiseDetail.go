package transform

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PriceType"
	"bitrade/core/entity"
	"github.com/qauzy/math"
)

func (this *MemberAdvertiseDetail) SetId(id int64) (result *MemberAdvertiseDetail) {
	this.Id = id
	return this
}
func (this *MemberAdvertiseDetail) GetId() (id int64) {
	return this.Id
}
func (this *MemberAdvertiseDetail) SetCoinId(coinId int64) (result *MemberAdvertiseDetail) {
	this.CoinId = coinId
	return this
}
func (this *MemberAdvertiseDetail) GetCoinId() (coinId int64) {
	return this.CoinId
}
func (this *MemberAdvertiseDetail) SetCoinName(coinName string) (result *MemberAdvertiseDetail) {
	this.CoinName = coinName
	return this
}
func (this *MemberAdvertiseDetail) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *MemberAdvertiseDetail) SetCoinNameCn(coinNameCn string) (result *MemberAdvertiseDetail) {
	this.CoinNameCn = coinNameCn
	return this
}
func (this *MemberAdvertiseDetail) GetCoinNameCn() (coinNameCn string) {
	return this.CoinNameCn
}
func (this *MemberAdvertiseDetail) SetCoinUnit(coinUnit string) (result *MemberAdvertiseDetail) {
	this.CoinUnit = coinUnit
	return this
}
func (this *MemberAdvertiseDetail) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *MemberAdvertiseDetail) SetCountry(country *entity.Country) (result *MemberAdvertiseDetail) {
	this.Country = country
	return this
}
func (this *MemberAdvertiseDetail) GetCountry() (country *entity.Country) {
	return this.Country
}
func (this *MemberAdvertiseDetail) SetPriceType(priceType PriceType.PriceType) (result *MemberAdvertiseDetail) {
	this.PriceType = priceType
	return this
}
func (this *MemberAdvertiseDetail) GetPriceType() (priceType PriceType.PriceType) {
	return this.PriceType
}
func (this *MemberAdvertiseDetail) SetPrice(price math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.Price = price
	return this
}
func (this *MemberAdvertiseDetail) GetPrice() (price math.BigDecimal) {
	return this.Price
}
func (this *MemberAdvertiseDetail) SetAdvertiseType(advertiseType AdvertiseType.AdvertiseType) (result *MemberAdvertiseDetail) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *MemberAdvertiseDetail) GetAdvertiseType() (advertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *MemberAdvertiseDetail) SetMinLimit(minLimit math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.MinLimit = minLimit
	return this
}
func (this *MemberAdvertiseDetail) GetMinLimit() (minLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *MemberAdvertiseDetail) SetMaxLimit(maxLimit math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.MaxLimit = maxLimit
	return this
}
func (this *MemberAdvertiseDetail) GetMaxLimit() (maxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *MemberAdvertiseDetail) SetRemark(remark string) (result *MemberAdvertiseDetail) {
	this.Remark = remark
	return this
}
func (this *MemberAdvertiseDetail) GetRemark() (remark string) {
	return this.Remark
}
func (this *MemberAdvertiseDetail) SetTimeLimit(timeLimit int64) (result *MemberAdvertiseDetail) {
	this.TimeLimit = timeLimit
	return this
}
func (this *MemberAdvertiseDetail) GetTimeLimit() (timeLimit int64) {
	return this.TimeLimit
}
func (this *MemberAdvertiseDetail) SetPremiseRate(premiseRate math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.PremiseRate = premiseRate
	return this
}
func (this *MemberAdvertiseDetail) GetPremiseRate() (premiseRate math.BigDecimal) {
	return this.PremiseRate
}
func (this *MemberAdvertiseDetail) SetPayMode(payMode string) (result *MemberAdvertiseDetail) {
	this.PayMode = payMode
	return this
}
func (this *MemberAdvertiseDetail) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *MemberAdvertiseDetail) SetStatus(status AdvertiseControlStatus.AdvertiseControlStatus) (result *MemberAdvertiseDetail) {
	this.Status = status
	return this
}
func (this *MemberAdvertiseDetail) GetStatus() (status AdvertiseControlStatus.AdvertiseControlStatus) {
	return this.Status
}
func (this *MemberAdvertiseDetail) SetNumber(number math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.Number = number
	return this
}
func (this *MemberAdvertiseDetail) GetNumber() (number math.BigDecimal) {
	return this.Number
}
func (this *MemberAdvertiseDetail) SetMarketPrice(marketPrice math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.MarketPrice = marketPrice
	return this
}
func (this *MemberAdvertiseDetail) GetMarketPrice() (marketPrice math.BigDecimal) {
	return this.MarketPrice
}
func (this *MemberAdvertiseDetail) SetAuto(auto BooleanEnum.BooleanEnum) (result *MemberAdvertiseDetail) {
	this.Auto = auto
	return this
}
func (this *MemberAdvertiseDetail) GetAuto() (auto BooleanEnum.BooleanEnum) {
	return this.Auto
}
func (this *MemberAdvertiseDetail) SetAutoword(autoword string) (result *MemberAdvertiseDetail) {
	this.Autoword = autoword
	return this
}
func (this *MemberAdvertiseDetail) GetAutoword() (autoword string) {
	return this.Autoword
}
func ToMemberAdvertiseDetail(advertise *entity.Advertise) (result *MemberAdvertiseDetail) {
	return new(MemberAdvertiseDetail).SetId(advertise.GetId()).SetAdvertiseType(advertise.GetAdvertiseType()).SetCoinId(advertise.GetCoin().GetId()).SetCoinName(advertise.GetCoin().GetName()).SetCoinNameCn(advertise.GetCoin().GetNameCn()).SetCoinUnit(advertise.GetCoin().GetUnit()).SetCountry(advertise.GetCountry()).SetAuto(advertise.GetAuto()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetNumber(advertise.GetNumber()).SetPayMode(advertise.GetPayMode()).SetPremiseRate(advertise.GetPremiseRate()).SetPrice(advertise.GetPrice()).SetPriceType(advertise.GetPriceType()).SetRemark(advertise.GetRemark()).SetStatus(advertise.GetStatus()).SetTimeLimit(advertise.GetTimeLimit()).SetAutoword(advertise.GetAutoword())
}

type MemberAdvertiseDetail struct {
	Id            int64
	CoinId        int64
	CoinName      string
	CoinNameCn    string
	CoinUnit      string
	Country       *entity.Country
	PriceType     PriceType.PriceType
	Price         math.BigDecimal
	AdvertiseType AdvertiseType.AdvertiseType
	MinLimit      math.BigDecimal
	MaxLimit      math.BigDecimal
	Remark        string
	TimeLimit     int64
	PremiseRate   math.BigDecimal
	PayMode       string
	Status        AdvertiseControlStatus.AdvertiseControlStatus
	Number        math.BigDecimal
	MarketPrice   math.BigDecimal
	Auto          BooleanEnum.BooleanEnum
	Autoword      string
}
