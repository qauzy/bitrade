package vo

import (
	"github.com/qauzy/math"
	"time"
)

func (this *ExchangeTurnoverStatisticsVO) SetDate(date time.Time) (result *ExchangeTurnoverStatisticsVO) {
	this.Date = date
	return this
}
func (this *ExchangeTurnoverStatisticsVO) GetDate() (date time.Time) {
	return this.Date
}
func (this *ExchangeTurnoverStatisticsVO) SetBaseSymbol(baseSymbol string) (result *ExchangeTurnoverStatisticsVO) {
	this.BaseSymbol = baseSymbol
	return this
}
func (this *ExchangeTurnoverStatisticsVO) GetBaseSymbol() (baseSymbol string) {
	return this.BaseSymbol
}
func (this *ExchangeTurnoverStatisticsVO) SetCoinSymbol(coinSymbol string) (result *ExchangeTurnoverStatisticsVO) {
	this.CoinSymbol = coinSymbol
	return this
}
func (this *ExchangeTurnoverStatisticsVO) GetCoinSymbol() (coinSymbol string) {
	return this.CoinSymbol
}
func (this *ExchangeTurnoverStatisticsVO) SetAmount(amount math.BigDecimal) (result *ExchangeTurnoverStatisticsVO) {
	this.Amount = amount
	return this
}
func (this *ExchangeTurnoverStatisticsVO) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *ExchangeTurnoverStatisticsVO) SetMoney(money math.BigDecimal) (result *ExchangeTurnoverStatisticsVO) {
	this.Money = money
	return this
}
func (this *ExchangeTurnoverStatisticsVO) GetMoney() (money math.BigDecimal) {
	return this.Money
}

type ExchangeTurnoverStatisticsVO struct {
	Date       time.Time
	BaseSymbol string
	CoinSymbol string
	Amount     math.BigDecimal
	Money      math.BigDecimal
}
