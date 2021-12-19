package entity

func (this *TurnoverStatistics) SetDate(date time.Time) (result *TurnoverStatistics) {
	this.Date = date
	return this
}
func (this *TurnoverStatistics) GetDate() (date time.Time) {
	return this.Date
}
func (this *TurnoverStatistics) SetYear(year int) (result *TurnoverStatistics) {
	this.Year = year
	return this
}
func (this *TurnoverStatistics) GetYear() (year int) {
	return this.Year
}
func (this *TurnoverStatistics) SetMonth(month int) (result *TurnoverStatistics) {
	this.Month = month
	return this
}
func (this *TurnoverStatistics) GetMonth() (month int) {
	return this.Month
}
func (this *TurnoverStatistics) SetDay(day int) (result *TurnoverStatistics) {
	this.Day = day
	return this
}
func (this *TurnoverStatistics) GetDay() (day int) {
	return this.Day
}
func (this *TurnoverStatistics) SetType(oType constant.TransactionTypeEnum) (result *TurnoverStatistics) {
	this.Type = oType
	return this
}
func (this *TurnoverStatistics) GetType() (oType constant.TransactionTypeEnum) {
	return this.Type
}
func (this *TurnoverStatistics) SetUnit(unit string) (result *TurnoverStatistics) {
	this.Unit = unit
	return this
}
func (this *TurnoverStatistics) GetUnit() (unit string) {
	return this.Unit
}
func (this *TurnoverStatistics) SetAmount(amount math.BigDecimal) (result *TurnoverStatistics) {
	this.Amount = amount
	return this
}
func (this *TurnoverStatistics) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *TurnoverStatistics) SetFee(fee math.BigDecimal) (result *TurnoverStatistics) {
	this.Fee = fee
	return this
}
func (this *TurnoverStatistics) GetFee() (fee math.BigDecimal) {
	return this.Fee
}

type TurnoverStatistics struct {
	Date   time.Time
	Year   int
	Month  int
	Day    int
	Type   constant.TransactionTypeEnum
	Unit   string
	Amount math.BigDecimal
	Fee    math.BigDecimal
}
