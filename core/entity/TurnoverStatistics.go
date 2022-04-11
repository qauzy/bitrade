package entity

import (
	"bitrade/core/constant/TransactionTypeEnum"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *TurnoverStatistics) SetDate(Date xtime.Xtime) (result *TurnoverStatistics) {
	this.Date = Date
	return this
}
func (this *TurnoverStatistics) GetDate() (Date xtime.Xtime) {
	return this.Date
}
func (this *TurnoverStatistics) SetYear(Year int) (result *TurnoverStatistics) {
	this.Year = Year
	return this
}
func (this *TurnoverStatistics) GetYear() (Year int) {
	return this.Year
}
func (this *TurnoverStatistics) SetMonth(Month int) (result *TurnoverStatistics) {
	this.Month = Month
	return this
}
func (this *TurnoverStatistics) GetMonth() (Month int) {
	return this.Month
}
func (this *TurnoverStatistics) SetDay(Day int) (result *TurnoverStatistics) {
	this.Day = Day
	return this
}
func (this *TurnoverStatistics) GetDay() (Day int) {
	return this.Day
}
func (this *TurnoverStatistics) SetType(Type TransactionTypeEnum.TransactionTypeEnum) (result *TurnoverStatistics) {
	this.Type = Type
	return this
}
func (this *TurnoverStatistics) GetType() (Type TransactionTypeEnum.TransactionTypeEnum) {
	return this.Type
}
func (this *TurnoverStatistics) SetUnit(Unit string) (result *TurnoverStatistics) {
	this.Unit = Unit
	return this
}
func (this *TurnoverStatistics) GetUnit() (Unit string) {
	return this.Unit
}
func (this *TurnoverStatistics) SetAmount(Amount math.BigDecimal) (result *TurnoverStatistics) {
	this.Amount = Amount
	return this
}
func (this *TurnoverStatistics) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *TurnoverStatistics) SetFee(Fee math.BigDecimal) (result *TurnoverStatistics) {
	this.Fee = Fee
	return this
}
func (this *TurnoverStatistics) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func NewTurnoverStatistics(date xtime.Xtime, year int, month int, day int, oType TransactionTypeEnum.TransactionTypeEnum, unit string, amount math.BigDecimal, fee math.BigDecimal) (ret *TurnoverStatistics) {
	ret = new(TurnoverStatistics)
	ret.Date = date
	ret.Year = year
	ret.Month = month
	ret.Day = day
	ret.Type = oType
	ret.Unit = unit
	ret.Amount = amount
	ret.Fee = fee
	return
}

type TurnoverStatistics struct {
	Date   xtime.Xtime                             `gorm:"column:date" json:"date"`
	Year   int                                     `gorm:"column:year" json:"year"`
	Month  int                                     `gorm:"column:month" json:"month"`
	Day    int                                     `gorm:"column:day" json:"day"`
	Type   TransactionTypeEnum.TransactionTypeEnum `gorm:"column:type" json:"type"`
	Unit   string                                  `gorm:"column:unit" json:"unit"`
	Amount math.BigDecimal                         `gorm:"column:amount" json:"amount"`
	Fee    math.BigDecimal                         `gorm:"column:fee" json:"fee"`
}
