package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
)

func (this *OtcCoin) SetId(id int64) (result *OtcCoin) {
	this.Id = id
	return this
}
func (this *OtcCoin) GetId() (id int64) {
	return this.Id
}
func (this *OtcCoin) SetName(name string) (result *OtcCoin) {
	this.Name = name
	return this
}
func (this *OtcCoin) GetName() (name string) {
	return this.Name
}
func (this *OtcCoin) SetNameCn(nameCn string) (result *OtcCoin) {
	this.NameCn = nameCn
	return this
}
func (this *OtcCoin) GetNameCn() (nameCn string) {
	return this.NameCn
}
func (this *OtcCoin) SetUnit(unit string) (result *OtcCoin) {
	this.Unit = unit
	return this
}
func (this *OtcCoin) GetUnit() (unit string) {
	return this.Unit
}
func (this *OtcCoin) SetStatus(status CommonStatus.CommonStatus) (result *OtcCoin) {
	this.Status = status
	return this
}
func (this *OtcCoin) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *OtcCoin) SetJyRate(jyRate math.BigDecimal) (result *OtcCoin) {
	this.JyRate = jyRate
	return this
}
func (this *OtcCoin) GetJyRate() (jyRate math.BigDecimal) {
	return this.JyRate
}
func (this *OtcCoin) SetSellMinAmount(sellMinAmount math.BigDecimal) (result *OtcCoin) {
	this.SellMinAmount = sellMinAmount
	return this
}
func (this *OtcCoin) GetSellMinAmount() (sellMinAmount math.BigDecimal) {
	return this.SellMinAmount
}
func (this *OtcCoin) SetBuyMinAmount(buyMinAmount math.BigDecimal) (result *OtcCoin) {
	this.BuyMinAmount = buyMinAmount
	return this
}
func (this *OtcCoin) GetBuyMinAmount() (buyMinAmount math.BigDecimal) {
	return this.BuyMinAmount
}
func (this *OtcCoin) SetSort(sort int) (result *OtcCoin) {
	this.Sort = sort
	return this
}
func (this *OtcCoin) GetSort() (sort int) {
	return this.Sort
}
func (this *OtcCoin) SetIsPlatformCoin(isPlatformCoin BooleanEnum.BooleanEnum) (result *OtcCoin) {
	this.IsPlatformCoin = isPlatformCoin
	return this
}
func (this *OtcCoin) GetIsPlatformCoin() (isPlatformCoin BooleanEnum.BooleanEnum) {
	return this.IsPlatformCoin
}
func (this *OtcCoin) SetCoinScale(coinScale int64) (result *OtcCoin) {
	this.CoinScale = coinScale
	return this
}
func (this *OtcCoin) GetCoinScale() (coinScale int64) {
	return this.CoinScale
}
func (this *OtcCoin) SetMaxTradingTime(maxTradingTime int64) (result *OtcCoin) {
	this.MaxTradingTime = maxTradingTime
	return this
}
func (this *OtcCoin) GetMaxTradingTime() (maxTradingTime int64) {
	return this.MaxTradingTime
}
func (this *OtcCoin) SetMaxVolume(maxVolume int64) (result *OtcCoin) {
	this.MaxVolume = maxVolume
	return this
}
func (this *OtcCoin) GetMaxVolume() (maxVolume int64) {
	return this.MaxVolume
}

type OtcCoin struct {
	Id             int64
	Name           string
	NameCn         string
	Unit           string
	Status         CommonStatus.CommonStatus
	JyRate         math.BigDecimal
	SellMinAmount  math.BigDecimal
	BuyMinAmount   math.BigDecimal
	Sort           int
	IsPlatformCoin BooleanEnum.BooleanEnum
	CoinScale      int64
	MaxTradingTime int64
	MaxVolume      int64
}
