package entity

func (this *MemberLog) SetDate(date time.Time) (result *MemberLog) {
	this.Date = date
	return this
}
func (this *MemberLog) GetDate() (date time.Time) {
	return this.Date
}
func (this *MemberLog) SetYear(year int) (result *MemberLog) {
	this.Year = year
	return this
}
func (this *MemberLog) GetYear() (year int) {
	return this.Year
}
func (this *MemberLog) SetMonth(month int) (result *MemberLog) {
	this.Month = month
	return this
}
func (this *MemberLog) GetMonth() (month int) {
	return this.Month
}
func (this *MemberLog) SetDay(day int) (result *MemberLog) {
	this.Day = day
	return this
}
func (this *MemberLog) GetDay() (day int) {
	return this.Day
}
func (this *MemberLog) SetRegistrationNum(registrationNum int) (result *MemberLog) {
	this.RegistrationNum = registrationNum
	return this
}
func (this *MemberLog) GetRegistrationNum() (registrationNum int) {
	return this.RegistrationNum
}
func (this *MemberLog) SetApplicationNum(applicationNum int) (result *MemberLog) {
	this.ApplicationNum = applicationNum
	return this
}
func (this *MemberLog) GetApplicationNum() (applicationNum int) {
	return this.ApplicationNum
}
func (this *MemberLog) SetBussinessNum(bussinessNum int) (result *MemberLog) {
	this.BussinessNum = bussinessNum
	return this
}
func (this *MemberLog) GetBussinessNum() (bussinessNum int) {
	return this.BussinessNum
}

type MemberLog struct {
	Date            time.Time
	Year            int
	Month           int
	Day             int
	RegistrationNum int
	ApplicationNum  int
	BussinessNum    int
}
