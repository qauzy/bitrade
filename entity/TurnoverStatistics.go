
package entity

func (this *TurnoverStatistics) SetDate(date Date) {
	this.Date = date
}
func (this *TurnoverStatistics) GetDate() (date Date) {
	return this.Date
}
func (this *TurnoverStatistics) SetYear(year int) {
	this.Year = year
}
func (this *TurnoverStatistics) GetYear() (year int) {
	return this.Year
}
func (this *TurnoverStatistics) SetMonth(month int) {
	this.Month = month
}
func (this *TurnoverStatistics) GetMonth() (month int) {
	return this.Month
}
func (this *TurnoverStatistics) SetDay(day int) {
	this.Day = day
}
func (this *TurnoverStatistics) GetDay() (day int) {
	return this.Day
}
func (this *TurnoverStatistics) SetType(type TransactionTypeEnum) {
	this.Type = type
}
func (this *TurnoverStatistics) GetType() (type TransactionTypeEnum) {
	return this.Type
}
func (this *TurnoverStatistics) SetUnit(unit string) {
	this.Unit = unit
}
func (this *TurnoverStatistics) GetUnit() (unit string) {
	return this.Unit
}
func (this *TurnoverStatistics) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *TurnoverStatistics) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *TurnoverStatistics) SetFee(fee BigDecimal) {
	this.Fee = fee
}
func (this *TurnoverStatistics) GetFee() (fee BigDecimal) {
	return this.Fee
}

type TurnoverStatistics struct {
	Date   Date
	Year   int
	Month  int
	Day    int
	Type   TransactionTypeEnum
	Unit   string
	Amount BigDecimal
	Fee    BigDecimal
}

