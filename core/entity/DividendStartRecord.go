package entity

import (
	"github.com/qauzy/math"
	"time"
)

func (this *DividendStartRecord) SetId(id int64) (result *DividendStartRecord) {
	this.Id = id
	return this
}
func (this *DividendStartRecord) GetId() (id int64) {
	return this.Id
}
func (this *DividendStartRecord) SetStart(start int64) (result *DividendStartRecord) {
	this.Start = start
	return this
}
func (this *DividendStartRecord) GetStart() (start int64) {
	return this.Start
}
func (this *DividendStartRecord) SetStartDate(startDate time.Time) (result *DividendStartRecord) {
	this.StartDate = startDate
	return this
}
func (this *DividendStartRecord) GetStartDate() (startDate time.Time) {
	return this.StartDate
}
func (this *DividendStartRecord) SetEnd(end int64) (result *DividendStartRecord) {
	this.End = end
	return this
}
func (this *DividendStartRecord) GetEnd() (end int64) {
	return this.End
}
func (this *DividendStartRecord) SetEndDate(endDate time.Time) (result *DividendStartRecord) {
	this.EndDate = endDate
	return this
}
func (this *DividendStartRecord) GetEndDate() (endDate time.Time) {
	return this.EndDate
}
func (this *DividendStartRecord) SetDate(date time.Time) (result *DividendStartRecord) {
	this.Date = date
	return this
}
func (this *DividendStartRecord) GetDate() (date time.Time) {
	return this.Date
}
func (this *DividendStartRecord) SetUnit(unit string) (result *DividendStartRecord) {
	this.Unit = unit
	return this
}
func (this *DividendStartRecord) GetUnit() (unit string) {
	return this.Unit
}
func (this *DividendStartRecord) SetRate(rate math.BigDecimal) (result *DividendStartRecord) {
	this.Rate = rate
	return this
}
func (this *DividendStartRecord) GetRate() (rate math.BigDecimal) {
	return this.Rate
}
func (this *DividendStartRecord) SetAmount(amount math.BigDecimal) (result *DividendStartRecord) {
	this.Amount = amount
	return this
}
func (this *DividendStartRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *DividendStartRecord) SetAdmin(admin Admin) (result *DividendStartRecord) {
	this.Admin = admin
	return this
}
func (this *DividendStartRecord) GetAdmin() (admin Admin) {
	return this.Admin
}

type DividendStartRecord struct {
	Id        int64
	Start     int64
	StartDate time.Time
	End       int64
	EndDate   time.Time
	Date      time.Time
	Unit      string
	Rate      math.BigDecimal
	Amount    math.BigDecimal
	Admin     Admin
}
