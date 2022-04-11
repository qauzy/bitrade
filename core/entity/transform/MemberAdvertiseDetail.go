package transform

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PriceType"
	"bitrade/core/entity"
	"github.com/qauzy/math"
)

func (this *MemberAdvertiseDetail) SetId(Id int64) (result *MemberAdvertiseDetail) {
	this.Id = Id
	return this
}
func (this *MemberAdvertiseDetail) GetId() (Id int64) {
	return this.Id
}
func (this *MemberAdvertiseDetail) SetCoinId(CoinId int64) (result *MemberAdvertiseDetail) {
	this.CoinId = CoinId
	return this
}
func (this *MemberAdvertiseDetail) GetCoinId() (CoinId int64) {
	return this.CoinId
}
func (this *MemberAdvertiseDetail) SetCoinName(CoinName string) (result *MemberAdvertiseDetail) {
	this.CoinName = CoinName
	return this
}
func (this *MemberAdvertiseDetail) GetCoinName() (CoinName string) {
	return this.CoinName
}
func (this *MemberAdvertiseDetail) SetCoinNameCn(CoinNameCn string) (result *MemberAdvertiseDetail) {
	this.CoinNameCn = CoinNameCn
	return this
}
func (this *MemberAdvertiseDetail) GetCoinNameCn() (CoinNameCn string) {
	return this.CoinNameCn
}
func (this *MemberAdvertiseDetail) SetCoinUnit(CoinUnit string) (result *MemberAdvertiseDetail) {
	this.CoinUnit = CoinUnit
	return this
}
func (this *MemberAdvertiseDetail) GetCoinUnit() (CoinUnit string) {
	return this.CoinUnit
}
func (this *MemberAdvertiseDetail) SetCountry(Country *entity.Country) (result *MemberAdvertiseDetail) {
	this.Country = Country
	return this
}
func (this *MemberAdvertiseDetail) GetCountry() (Country *entity.Country) {
	return this.Country
}
func (this *MemberAdvertiseDetail) SetPriceType(PriceType PriceType.PriceType) (result *MemberAdvertiseDetail) {
	this.PriceType = PriceType
	return this
}
func (this *MemberAdvertiseDetail) GetPriceType() (PriceType PriceType.PriceType) {
	return this.PriceType
}
func (this *MemberAdvertiseDetail) SetPrice(Price math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.Price = Price
	return this
}
func (this *MemberAdvertiseDetail) GetPrice() (Price math.BigDecimal) {
	return this.Price
}
func (this *MemberAdvertiseDetail) SetAdvertiseType(AdvertiseType AdvertiseType.AdvertiseType) (result *MemberAdvertiseDetail) {
	this.AdvertiseType = AdvertiseType
	return this
}
func (this *MemberAdvertiseDetail) GetAdvertiseType() (AdvertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *MemberAdvertiseDetail) SetMinLimit(MinLimit math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.MinLimit = MinLimit
	return this
}
func (this *MemberAdvertiseDetail) GetMinLimit() (MinLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *MemberAdvertiseDetail) SetMaxLimit(MaxLimit math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.MaxLimit = MaxLimit
	return this
}
func (this *MemberAdvertiseDetail) GetMaxLimit() (MaxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *MemberAdvertiseDetail) SetRemark(Remark string) (result *MemberAdvertiseDetail) {
	this.Remark = Remark
	return this
}
func (this *MemberAdvertiseDetail) GetRemark() (Remark string) {
	return this.Remark
}
func (this *MemberAdvertiseDetail) SetTimeLimit(TimeLimit int) (result *MemberAdvertiseDetail) {
	this.TimeLimit = TimeLimit
	return this
}
func (this *MemberAdvertiseDetail) GetTimeLimit() (TimeLimit int) {
	return this.TimeLimit
}
func (this *MemberAdvertiseDetail) SetPremiseRate(PremiseRate math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.PremiseRate = PremiseRate
	return this
}
func (this *MemberAdvertiseDetail) GetPremiseRate() (PremiseRate math.BigDecimal) {
	return this.PremiseRate
}
func (this *MemberAdvertiseDetail) SetPayMode(PayMode string) (result *MemberAdvertiseDetail) {
	this.PayMode = PayMode
	return this
}
func (this *MemberAdvertiseDetail) GetPayMode() (PayMode string) {
	return this.PayMode
}
func (this *MemberAdvertiseDetail) SetStatus(Status AdvertiseControlStatus.AdvertiseControlStatus) (result *MemberAdvertiseDetail) {
	this.Status = Status
	return this
}
func (this *MemberAdvertiseDetail) GetStatus() (Status AdvertiseControlStatus.AdvertiseControlStatus) {
	return this.Status
}
func (this *MemberAdvertiseDetail) SetNumber(Number math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.Number = Number
	return this
}
func (this *MemberAdvertiseDetail) GetNumber() (Number math.BigDecimal) {
	return this.Number
}
func (this *MemberAdvertiseDetail) SetMarketPrice(MarketPrice math.BigDecimal) (result *MemberAdvertiseDetail) {
	this.MarketPrice = MarketPrice
	return this
}
func (this *MemberAdvertiseDetail) GetMarketPrice() (MarketPrice math.BigDecimal) {
	return this.MarketPrice
}
func (this *MemberAdvertiseDetail) SetAuto(Auto BooleanEnum.BooleanEnum) (result *MemberAdvertiseDetail) {
	this.Auto = Auto
	return this
}
func (this *MemberAdvertiseDetail) GetAuto() (Auto BooleanEnum.BooleanEnum) {
	return this.Auto
}
func (this *MemberAdvertiseDetail) SetAutoword(Autoword string) (result *MemberAdvertiseDetail) {
	this.Autoword = Autoword
	return this
}
func (this *MemberAdvertiseDetail) GetAutoword() (Autoword string) {
	return this.Autoword
}
func ToMemberAdvertiseDetail(advertise *entity.Advertise) (result *MemberAdvertiseDetail) {
	return new(MemberAdvertiseDetail).SetId(advertise.GetId()).SetAdvertiseType(advertise.GetAdvertiseType()).SetCoinId(advertise.GetCoin().GetId()).SetCoinName(advertise.GetCoin().GetName()).SetCoinNameCn(advertise.GetCoin().GetNameCn()).SetCoinUnit(advertise.GetCoin().GetUnit()).SetCountry(advertise.GetCountry()).SetAuto(advertise.GetAuto()).SetMaxLimit(advertise.GetMaxLimit()).SetMinLimit(advertise.GetMinLimit()).SetNumber(advertise.GetNumber()).SetPayMode(advertise.GetPayMode()).SetPremiseRate(advertise.GetPremiseRate()).SetPrice(advertise.GetPrice()).SetPriceType(advertise.GetPriceType()).SetRemark(advertise.GetRemark()).SetStatus(advertise.GetStatus()).SetTimeLimit(advertise.GetTimeLimit()).SetAutoword(advertise.GetAutoword())
}
func NewMemberAdvertiseDetail(id int64, coinId int64, coinName string, coinNameCn string, coinUnit string, country *entity.Country, priceType PriceType.PriceType, price math.BigDecimal, advertiseType AdvertiseType.AdvertiseType, minLimit math.BigDecimal, maxLimit math.BigDecimal, remark string, timeLimit int, premiseRate math.BigDecimal, payMode string, status AdvertiseControlStatus.AdvertiseControlStatus, number math.BigDecimal, marketPrice math.BigDecimal, auto BooleanEnum.BooleanEnum, autoword string) (ret *MemberAdvertiseDetail) {
	ret = new(MemberAdvertiseDetail)
	ret.Id = id
	ret.CoinId = coinId
	ret.CoinName = coinName
	ret.CoinNameCn = coinNameCn
	ret.CoinUnit = coinUnit
	ret.Country = country
	ret.PriceType = priceType
	ret.Price = price
	ret.AdvertiseType = advertiseType
	ret.MinLimit = minLimit
	ret.MaxLimit = maxLimit
	ret.Remark = remark
	ret.TimeLimit = timeLimit
	ret.PremiseRate = premiseRate
	ret.PayMode = payMode
	ret.Status = status
	ret.Number = number
	ret.MarketPrice = marketPrice
	ret.Auto = auto
	ret.Autoword = autoword
	return
}

type MemberAdvertiseDetail struct {
	Id            int64                                         `gorm:"column:id" json:"id"`
	CoinId        int64                                         `gorm:"column:coin_id" json:"coinId"`
	CoinName      string                                        `gorm:"column:coin_name" json:"coinName"`
	CoinNameCn    string                                        `gorm:"column:coin_name_cn" json:"coinNameCn"`
	CoinUnit      string                                        `gorm:"column:coin_unit" json:"coinUnit"`
	Country       *entity.Country                               `gorm:"column:country" json:"country"`
	PriceType     PriceType.PriceType                           `gorm:"column:price_type" json:"priceType"`
	Price         math.BigDecimal                               `gorm:"column:price" json:"price"`
	AdvertiseType AdvertiseType.AdvertiseType                   `gorm:"column:advertise_type" json:"advertiseType"`
	MinLimit      math.BigDecimal                               `gorm:"column:min_limit" json:"minLimit"`
	MaxLimit      math.BigDecimal                               `gorm:"column:max_limit" json:"maxLimit"`
	Remark        string                                        `gorm:"column:remark" json:"remark"`
	TimeLimit     int                                           `gorm:"column:time_limit" json:"timeLimit"`
	PremiseRate   math.BigDecimal                               `gorm:"column:premise_rate" json:"premiseRate"`
	PayMode       string                                        `gorm:"column:pay_mode" json:"payMode"`
	Status        AdvertiseControlStatus.AdvertiseControlStatus `gorm:"column:status" json:"status"`
	Number        math.BigDecimal                               `gorm:"column:number" json:"number"`
	MarketPrice   math.BigDecimal                               `gorm:"column:market_price" json:"marketPrice"`
	Auto          BooleanEnum.BooleanEnum                       `gorm:"column:auto" json:"auto"`
	Autoword      string                                        `gorm:"column:autoword" json:"autoword"`
}
