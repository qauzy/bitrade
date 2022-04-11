package entity

import "github.com/qauzy/chocolate/xtime"

func (this *MemberLog) SetDate(Date xtime.Xtime) (result *MemberLog) {
	this.Date = Date
	return this
}
func (this *MemberLog) GetDate() (Date xtime.Xtime) {
	return this.Date
}
func (this *MemberLog) SetYear(Year int) (result *MemberLog) {
	this.Year = Year
	return this
}
func (this *MemberLog) GetYear() (Year int) {
	return this.Year
}
func (this *MemberLog) SetMonth(Month int) (result *MemberLog) {
	this.Month = Month
	return this
}
func (this *MemberLog) GetMonth() (Month int) {
	return this.Month
}
func (this *MemberLog) SetDay(Day int) (result *MemberLog) {
	this.Day = Day
	return this
}
func (this *MemberLog) GetDay() (Day int) {
	return this.Day
}
func (this *MemberLog) SetRegistrationNum(RegistrationNum int) (result *MemberLog) {
	this.RegistrationNum = RegistrationNum
	return this
}
func (this *MemberLog) GetRegistrationNum() (RegistrationNum int) {
	return this.RegistrationNum
}
func (this *MemberLog) SetApplicationNum(ApplicationNum int) (result *MemberLog) {
	this.ApplicationNum = ApplicationNum
	return this
}
func (this *MemberLog) GetApplicationNum() (ApplicationNum int) {
	return this.ApplicationNum
}
func (this *MemberLog) SetBussinessNum(BussinessNum int) (result *MemberLog) {
	this.BussinessNum = BussinessNum
	return this
}
func (this *MemberLog) GetBussinessNum() (BussinessNum int) {
	return this.BussinessNum
}
func NewMemberLog(date xtime.Xtime, year int, month int, day int, registrationNum int, applicationNum int, bussinessNum int) (ret *MemberLog) {
	ret = new(MemberLog)
	ret.Date = date
	ret.Year = year
	ret.Month = month
	ret.Day = day
	ret.RegistrationNum = registrationNum
	ret.ApplicationNum = applicationNum
	ret.BussinessNum = bussinessNum
	return
}

type MemberLog struct {
	Date            xtime.Xtime `gorm:"column:date" json:"date"`
	Year            int         `gorm:"column:year" json:"year"`
	Month           int         `gorm:"column:month" json:"month"`
	Day             int         `gorm:"column:day" json:"day"`
	RegistrationNum int         `gorm:"column:registration_num" json:"registrationNum"`
	ApplicationNum  int         `gorm:"column:application_num" json:"applicationNum"`
	BussinessNum    int         `gorm:"column:bussiness_num" json:"bussinessNum"`
}
