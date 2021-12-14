package entity

func (this *DividendStartRecord) SetId(id int64) {
	this.Id = id
}
func (this *DividendStartRecord) GetId() (id int64) {
	return this.Id
}
func (this *DividendStartRecord) SetStart(start int64) {
	this.Start = start
}
func (this *DividendStartRecord) GetStart() (start int64) {
	return this.Start
}
func (this *DividendStartRecord) SetStartDate(startDate Date) {
	this.StartDate = startDate
}
func (this *DividendStartRecord) GetStartDate() (startDate Date) {
	return this.StartDate
}
func (this *DividendStartRecord) SetEnd(end int64) {
	this.End = end
}
func (this *DividendStartRecord) GetEnd() (end int64) {
	return this.End
}
func (this *DividendStartRecord) SetEndDate(endDate Date) {
	this.EndDate = endDate
}
func (this *DividendStartRecord) GetEndDate() (endDate Date) {
	return this.EndDate
}
func (this *DividendStartRecord) SetDate(date Date) {
	this.Date = date
}
func (this *DividendStartRecord) GetDate() (date Date) {
	return this.Date
}
func (this *DividendStartRecord) SetUnit(unit string) {
	this.Unit = unit
}
func (this *DividendStartRecord) GetUnit() (unit string) {
	return this.Unit
}
func (this *DividendStartRecord) SetRate(rate BigDecimal) {
	this.Rate = rate
}
func (this *DividendStartRecord) GetRate() (rate BigDecimal) {
	return this.Rate
}
func (this *DividendStartRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *DividendStartRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *DividendStartRecord) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *DividendStartRecord) GetAdmin() (admin Admin) {
	return this.Admin
}

type DividendStartRecord struct {
	Id        int64
	Start     int64
	StartDate Date
	End       int64
	EndDate   Date
	Date      Date
	Unit      string
	Rate      BigDecimal
	Amount    BigDecimal
	Admin     Admin
}
