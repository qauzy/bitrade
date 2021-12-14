package entity

func (this *MemberLog) SetDate(date Date) {
	this.Date = date
}
func (this *MemberLog) GetDate() (date Date) {
	return this.Date
}
func (this *MemberLog) SetYear(year int) {
	this.Year = year
}
func (this *MemberLog) GetYear() (year int) {
	return this.Year
}
func (this *MemberLog) SetMonth(month int) {
	this.Month = month
}
func (this *MemberLog) GetMonth() (month int) {
	return this.Month
}
func (this *MemberLog) SetDay(day int) {
	this.Day = day
}
func (this *MemberLog) GetDay() (day int) {
	return this.Day
}
func (this *MemberLog) SetRegistrationNum(registrationNum int) {
	this.RegistrationNum = registrationNum
}
func (this *MemberLog) GetRegistrationNum() (registrationNum int) {
	return this.RegistrationNum
}
func (this *MemberLog) SetApplicationNum(applicationNum int) {
	this.ApplicationNum = applicationNum
}
func (this *MemberLog) GetApplicationNum() (applicationNum int) {
	return this.ApplicationNum
}
func (this *MemberLog) SetBussinessNum(bussinessNum int) {
	this.BussinessNum = bussinessNum
}
func (this *MemberLog) GetBussinessNum() (bussinessNum int) {
	return this.BussinessNum
}

type MemberLog struct {
	Date            Date
	Year            int
	Month           int
	Day             int
	RegistrationNum int
	ApplicationNum  int
	BussinessNum    int
}
