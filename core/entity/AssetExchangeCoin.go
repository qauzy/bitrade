package entity

import "github.com/qauzy/math"

func (this *AssetExchangeCoin) SetId(id int64) (result *AssetExchangeCoin) {
	this.Id = id
	return this
}
func (this *AssetExchangeCoin) GetId() (id int64) {
	return this.Id
}
func (this *AssetExchangeCoin) SetFromUnit(fromUnit string) (result *AssetExchangeCoin) {
	this.FromUnit = fromUnit
	return this
}
func (this *AssetExchangeCoin) GetFromUnit() (fromUnit string) {
	return this.FromUnit
}
func (this *AssetExchangeCoin) SetToUnit(toUnit string) (result *AssetExchangeCoin) {
	this.ToUnit = toUnit
	return this
}
func (this *AssetExchangeCoin) GetToUnit() (toUnit string) {
	return this.ToUnit
}
func (this *AssetExchangeCoin) SetExchangeRate(exchangeRate math.BigDecimal) (result *AssetExchangeCoin) {
	this.ExchangeRate = exchangeRate
	return this
}
func (this *AssetExchangeCoin) GetExchangeRate() (exchangeRate math.BigDecimal) {
	return this.ExchangeRate
}

type AssetExchangeCoin struct {
	Id           int64
	FromUnit     string
	ToUnit       string
	ExchangeRate math.BigDecimal
}
