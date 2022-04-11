package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *ExchangeTurnoverStatistics) SetDate(Date xtime.Xtime) (result *ExchangeTurnoverStatistics) {
	this.Date = Date
	return this
}
func (this *ExchangeTurnoverStatistics) GetDate() (Date xtime.Xtime) {
	return this.Date
}
func (this *ExchangeTurnoverStatistics) SetYear(Year int) (result *ExchangeTurnoverStatistics) {
	this.Year = Year
	return this
}
func (this *ExchangeTurnoverStatistics) GetYear() (Year int) {
	return this.Year
}
func (this *ExchangeTurnoverStatistics) SetMonth(Month int) (result *ExchangeTurnoverStatistics) {
	this.Month = Month
	return this
}
func (this *ExchangeTurnoverStatistics) GetMonth() (Month int) {
	return this.Month
}
func (this *ExchangeTurnoverStatistics) SetDay(Day int) (result *ExchangeTurnoverStatistics) {
	this.Day = Day
	return this
}
func (this *ExchangeTurnoverStatistics) GetDay() (Day int) {
	return this.Day
}
func (this *ExchangeTurnoverStatistics) SetBaseSymbol(BaseSymbol string) (result *ExchangeTurnoverStatistics) {
	this.BaseSymbol = BaseSymbol
	return this
}
func (this *ExchangeTurnoverStatistics) GetBaseSymbol() (BaseSymbol string) {
	return this.BaseSymbol
}
func (this *ExchangeTurnoverStatistics) SetCoinSymbol(CoinSymbol string) (result *ExchangeTurnoverStatistics) {
	this.CoinSymbol = CoinSymbol
	return this
}
func (this *ExchangeTurnoverStatistics) GetCoinSymbol() (CoinSymbol string) {
	return this.CoinSymbol
}
func (this *ExchangeTurnoverStatistics) SetAmount(Amount math.BigDecimal) (result *ExchangeTurnoverStatistics) {
	this.Amount = Amount
	return this
}
func (this *ExchangeTurnoverStatistics) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *ExchangeTurnoverStatistics) SetMoney(Money math.BigDecimal) (result *ExchangeTurnoverStatistics) {
	this.Money = Money
	return this
}
func (this *ExchangeTurnoverStatistics) GetMoney() (Money math.BigDecimal) {
	return this.Money
}
func NewExchangeTurnoverStatistics(date xtime.Xtime, year int, month int, day int, baseSymbol string, coinSymbol string, amount math.BigDecimal, money math.BigDecimal) (ret *ExchangeTurnoverStatistics) {
	ret = new(ExchangeTurnoverStatistics)
	ret.Date = date
	ret.Year = year
	ret.Month = month
	ret.Day = day
	ret.BaseSymbol = baseSymbol
	ret.CoinSymbol = coinSymbol
	ret.Amount = amount
	ret.Money = money
	return
}

type ExchangeTurnoverStatistics struct {
	Date       xtime.Xtime     `gorm:"column:date" json:"date"`
	Year       int             `gorm:"column:year" json:"year"`
	Month      int             `gorm:"column:month" json:"month"`
	Day        int             `gorm:"column:day" json:"day"`
	BaseSymbol string          `gorm:"column:base_symbol" json:"baseSymbol"`
	CoinSymbol string          `gorm:"column:coin_symbol" json:"coinSymbol"`
	Amount     math.BigDecimal `gorm:"column:amount" json:"amount"`
	Money      math.BigDecimal `gorm:"column:money" json:"money"`
}
