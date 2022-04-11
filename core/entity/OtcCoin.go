package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
)

func (this *OtcCoin) SetId(Id int64) (result *OtcCoin) {
	this.Id = Id
	return this
}
func (this *OtcCoin) GetId() (Id int64) {
	return this.Id
}
func (this *OtcCoin) SetName(Name string) (result *OtcCoin) {
	this.Name = Name
	return this
}
func (this *OtcCoin) GetName() (Name string) {
	return this.Name
}
func (this *OtcCoin) SetNameCn(NameCn string) (result *OtcCoin) {
	this.NameCn = NameCn
	return this
}
func (this *OtcCoin) GetNameCn() (NameCn string) {
	return this.NameCn
}
func (this *OtcCoin) SetUnit(Unit string) (result *OtcCoin) {
	this.Unit = Unit
	return this
}
func (this *OtcCoin) GetUnit() (Unit string) {
	return this.Unit
}
func (this *OtcCoin) SetStatus(Status CommonStatus.CommonStatus) (result *OtcCoin) {
	this.Status = Status
	return this
}
func (this *OtcCoin) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *OtcCoin) SetJyRate(JyRate math.BigDecimal) (result *OtcCoin) {
	this.JyRate = JyRate
	return this
}
func (this *OtcCoin) GetJyRate() (JyRate math.BigDecimal) {
	return this.JyRate
}
func (this *OtcCoin) SetSellMinAmount(SellMinAmount math.BigDecimal) (result *OtcCoin) {
	this.SellMinAmount = SellMinAmount
	return this
}
func (this *OtcCoin) GetSellMinAmount() (SellMinAmount math.BigDecimal) {
	return this.SellMinAmount
}
func (this *OtcCoin) SetBuyMinAmount(BuyMinAmount math.BigDecimal) (result *OtcCoin) {
	this.BuyMinAmount = BuyMinAmount
	return this
}
func (this *OtcCoin) GetBuyMinAmount() (BuyMinAmount math.BigDecimal) {
	return this.BuyMinAmount
}
func (this *OtcCoin) SetSort(Sort int) (result *OtcCoin) {
	this.Sort = Sort
	return this
}
func (this *OtcCoin) GetSort() (Sort int) {
	return this.Sort
}
func (this *OtcCoin) SetIsPlatformCoin(IsPlatformCoin BooleanEnum.BooleanEnum) (result *OtcCoin) {
	this.IsPlatformCoin = IsPlatformCoin
	return this
}
func (this *OtcCoin) GetIsPlatformCoin() (IsPlatformCoin BooleanEnum.BooleanEnum) {
	return this.IsPlatformCoin
}
func (this *OtcCoin) SetCoinScale(CoinScale int) (result *OtcCoin) {
	this.CoinScale = CoinScale
	return this
}
func (this *OtcCoin) GetCoinScale() (CoinScale int) {
	return this.CoinScale
}
func (this *OtcCoin) SetMaxTradingTime(MaxTradingTime int) (result *OtcCoin) {
	this.MaxTradingTime = MaxTradingTime
	return this
}
func (this *OtcCoin) GetMaxTradingTime() (MaxTradingTime int) {
	return this.MaxTradingTime
}
func (this *OtcCoin) SetMaxVolume(MaxVolume int) (result *OtcCoin) {
	this.MaxVolume = MaxVolume
	return this
}
func (this *OtcCoin) GetMaxVolume() (MaxVolume int) {
	return this.MaxVolume
}
func NewOtcCoin(id int64, name string, nameCn string, unit string, status CommonStatus.CommonStatus, jyRate math.BigDecimal, sellMinAmount math.BigDecimal, buyMinAmount math.BigDecimal, sort int, isPlatformCoin BooleanEnum.BooleanEnum, coinScale int, maxTradingTime int, maxVolume int) (ret *OtcCoin) {
	ret = new(OtcCoin)
	ret.Id = id
	ret.Name = name
	ret.NameCn = nameCn
	ret.Unit = unit
	ret.Status = status
	ret.JyRate = jyRate
	ret.SellMinAmount = sellMinAmount
	ret.BuyMinAmount = buyMinAmount
	ret.Sort = sort
	ret.IsPlatformCoin = isPlatformCoin
	ret.CoinScale = coinScale
	ret.MaxTradingTime = maxTradingTime
	ret.MaxVolume = maxVolume
	return
}

type OtcCoin struct {
	Id             int64                     `gorm:"column:id" json:"id"`
	Name           string                    `gorm:"column:name" json:"name"`
	NameCn         string                    `gorm:"column:name_cn" json:"nameCn"`
	Unit           string                    `gorm:"column:unit" json:"unit"`
	Status         CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
	JyRate         math.BigDecimal           `gorm:"column:jy_rate" json:"jyRate"`
	SellMinAmount  math.BigDecimal           `gorm:"column:sell_min_amount" json:"sellMinAmount"`
	BuyMinAmount   math.BigDecimal           `gorm:"column:buy_min_amount" json:"buyMinAmount"`
	Sort           int                       `gorm:"column:sort" json:"sort"`
	IsPlatformCoin BooleanEnum.BooleanEnum   `gorm:"column:is_platform_coin" json:"isPlatformCoin"`
	CoinScale      int                       `gorm:"column:coin_scale" json:"coinScale"`
	MaxTradingTime int                       `gorm:"column:max_trading_time" json:"maxTradingTime"`
	MaxVolume      int                       `gorm:"column:max_volume" json:"maxVolume"`
}
