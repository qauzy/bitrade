package entity

import "github.com/qauzy/math"

func (this *AssetExchangeCoin) SetId(Id int64) (result *AssetExchangeCoin) {
	this.Id = Id
	return this
}
func (this *AssetExchangeCoin) GetId() (Id int64) {
	return this.Id
}
func (this *AssetExchangeCoin) SetFromUnit(FromUnit string) (result *AssetExchangeCoin) {
	this.FromUnit = FromUnit
	return this
}
func (this *AssetExchangeCoin) GetFromUnit() (FromUnit string) {
	return this.FromUnit
}
func (this *AssetExchangeCoin) SetToUnit(ToUnit string) (result *AssetExchangeCoin) {
	this.ToUnit = ToUnit
	return this
}
func (this *AssetExchangeCoin) GetToUnit() (ToUnit string) {
	return this.ToUnit
}
func (this *AssetExchangeCoin) SetExchangeRate(ExchangeRate math.BigDecimal) (result *AssetExchangeCoin) {
	this.ExchangeRate = ExchangeRate
	return this
}
func (this *AssetExchangeCoin) GetExchangeRate() (ExchangeRate math.BigDecimal) {
	return this.ExchangeRate
}

type AssetExchangeCoin struct {
	Id           int64           `gorm:"column:id" json:"id"`
	FromUnit     string          `gorm:"column:from_unit" json:"fromUnit"`
	ToUnit       string          `gorm:"column:to_unit" json:"toUnit"`
	ExchangeRate math.BigDecimal `gorm:"column:exchange_rate" json:"exchangeRate"`
}
