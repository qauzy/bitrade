package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *DividendStartRecord) SetId(Id int64) (result *DividendStartRecord) {
	this.Id = Id
	return this
}
func (this *DividendStartRecord) GetId() (Id int64) {
	return this.Id
}
func (this *DividendStartRecord) SetStart(Start int64) (result *DividendStartRecord) {
	this.Start = Start
	return this
}
func (this *DividendStartRecord) GetStart() (Start int64) {
	return this.Start
}
func (this *DividendStartRecord) SetStartDate(StartDate xtime.Xtime) (result *DividendStartRecord) {
	this.StartDate = StartDate
	return this
}
func (this *DividendStartRecord) GetStartDate() (StartDate xtime.Xtime) {
	return this.StartDate
}
func (this *DividendStartRecord) SetEnd(End int64) (result *DividendStartRecord) {
	this.End = End
	return this
}
func (this *DividendStartRecord) GetEnd() (End int64) {
	return this.End
}
func (this *DividendStartRecord) SetEndDate(EndDate xtime.Xtime) (result *DividendStartRecord) {
	this.EndDate = EndDate
	return this
}
func (this *DividendStartRecord) GetEndDate() (EndDate xtime.Xtime) {
	return this.EndDate
}
func (this *DividendStartRecord) SetDate(Date xtime.Xtime) (result *DividendStartRecord) {
	this.Date = Date
	return this
}
func (this *DividendStartRecord) GetDate() (Date xtime.Xtime) {
	return this.Date
}
func (this *DividendStartRecord) SetUnit(Unit string) (result *DividendStartRecord) {
	this.Unit = Unit
	return this
}
func (this *DividendStartRecord) GetUnit() (Unit string) {
	return this.Unit
}
func (this *DividendStartRecord) SetRate(Rate math.BigDecimal) (result *DividendStartRecord) {
	this.Rate = Rate
	return this
}
func (this *DividendStartRecord) GetRate() (Rate math.BigDecimal) {
	return this.Rate
}
func (this *DividendStartRecord) SetAmount(Amount math.BigDecimal) (result *DividendStartRecord) {
	this.Amount = Amount
	return this
}
func (this *DividendStartRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *DividendStartRecord) SetAdmin(Admin *Admin) (result *DividendStartRecord) {
	this.Admin = Admin
	return this
}
func (this *DividendStartRecord) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func NewDividendStartRecord(id int64, start int64, startDate xtime.Xtime, end int64, endDate xtime.Xtime, date xtime.Xtime, unit string, rate math.BigDecimal, amount math.BigDecimal, admin *Admin) (ret *DividendStartRecord) {
	ret = new(DividendStartRecord)
	ret.Id = id
	ret.Start = start
	ret.StartDate = startDate
	ret.End = end
	ret.EndDate = endDate
	ret.Date = date
	ret.Unit = unit
	ret.Rate = rate
	ret.Amount = amount
	ret.Admin = admin
	return
}

type DividendStartRecord struct {
	Id        int64           `gorm:"column:id" json:"id"`
	Start     int64           `gorm:"column:start" json:"start"`
	StartDate xtime.Xtime     `gorm:"column:start_date" json:"startDate"`
	End       int64           `gorm:"column:end" json:"end"`
	EndDate   xtime.Xtime     `gorm:"column:end_date" json:"endDate"`
	Date      xtime.Xtime     `gorm:"column:date" json:"date"`
	Unit      string          `gorm:"column:unit" json:"unit"`
	Rate      math.BigDecimal `gorm:"column:rate" json:"rate"`
	Amount    math.BigDecimal `gorm:"column:amount" json:"amount"`
	Admin     *Admin          `gorm:"column:admin" json:"admin"`
}
