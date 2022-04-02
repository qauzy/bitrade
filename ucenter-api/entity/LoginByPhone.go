package entity

func (this *LoginByPhone) SetPhone(Phone string) (result *LoginByPhone) {
	this.Phone = Phone
	return this
}
func (this *LoginByPhone) GetPhone() (Phone string) {
	return this.Phone
}
func (this *LoginByPhone) SetPassword(Password string) (result *LoginByPhone) {
	this.Password = Password
	return this
}
func (this *LoginByPhone) GetPassword() (Password string) {
	return this.Password
}
func (this *LoginByPhone) SetUsername(Username string) (result *LoginByPhone) {
	this.Username = Username
	return this
}
func (this *LoginByPhone) GetUsername() (Username string) {
	return this.Username
}
func (this *LoginByPhone) SetCountry(Country string) (result *LoginByPhone) {
	this.Country = Country
	return this
}
func (this *LoginByPhone) GetCountry() (Country string) {
	return this.Country
}
func (this *LoginByPhone) SetCode(Code string) (result *LoginByPhone) {
	this.Code = Code
	return this
}
func (this *LoginByPhone) GetCode() (Code string) {
	return this.Code
}
func (this *LoginByPhone) SetTicket(Ticket string) (result *LoginByPhone) {
	this.Ticket = Ticket
	return this
}
func (this *LoginByPhone) GetTicket() (Ticket string) {
	return this.Ticket
}
func (this *LoginByPhone) SetRandStr(RandStr string) (result *LoginByPhone) {
	this.RandStr = RandStr
	return this
}
func (this *LoginByPhone) GetRandStr() (RandStr string) {
	return this.RandStr
}
func (this *LoginByPhone) SetPromotion(Promotion string) (result *LoginByPhone) {
	this.Promotion = Promotion
	return this
}
func (this *LoginByPhone) GetPromotion() (Promotion string) {
	return this.Promotion
}

type LoginByPhone struct {
	Phone     string `gorm:"column:phone" json:"phone"`
	Password  string `gorm:"column:password" json:"password"`
	Username  string `gorm:"column:username" json:"username"`
	Country   string `gorm:"column:country" json:"country"`
	Code      string `gorm:"column:code" json:"code"`
	Ticket    string `gorm:"column:ticket" json:"ticket"`
	RandStr   string `gorm:"column:rand_str" json:"randStr"`
	Promotion string `gorm:"column:promotion" json:"promotion"`
}
