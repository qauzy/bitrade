package vo

import (
	"github.com/qauzy/math"
	"time"
)

func (this *TurnoverStatisticsVO) SetDate(date time.Time) (result *TurnoverStatisticsVO) {
	this.Date = date
	return this
}
func (this *TurnoverStatisticsVO) GetDate() (date time.Time) {
	return this.Date
}
func (this *TurnoverStatisticsVO) SetUnit(unit string) (result *TurnoverStatisticsVO) {
	this.Unit = unit
	return this
}
func (this *TurnoverStatisticsVO) GetUnit() (unit string) {
	return this.Unit
}
func (this *TurnoverStatisticsVO) SetAmount(amount math.BigDecimal) (result *TurnoverStatisticsVO) {
	this.Amount = amount
	return this
}
func (this *TurnoverStatisticsVO) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *TurnoverStatisticsVO) SetMoney(money math.BigDecimal) (result *TurnoverStatisticsVO) {
	this.Money = money
	return this
}
func (this *TurnoverStatisticsVO) GetMoney() (money math.BigDecimal) {
	return this.Money
}

type TurnoverStatisticsVO struct {
	Date   time.Time
	Unit   string
	Amount math.BigDecimal
	Money  math.BigDecimal
}
