package entity

func (this *LoginByEmail) SetEmail(Email string) (result *LoginByEmail) {
	this.Email = Email
	return this
}
func (this *LoginByEmail) GetEmail() (Email string) {
	return this.Email
}
func (this *LoginByEmail) SetPassword(Password string) (result *LoginByEmail) {
	this.Password = Password
	return this
}
func (this *LoginByEmail) GetPassword() (Password string) {
	return this.Password
}
func (this *LoginByEmail) SetUsername(Username string) (result *LoginByEmail) {
	this.Username = Username
	return this
}
func (this *LoginByEmail) GetUsername() (Username string) {
	return this.Username
}
func (this *LoginByEmail) SetCountry(Country string) (result *LoginByEmail) {
	this.Country = Country
	return this
}
func (this *LoginByEmail) GetCountry() (Country string) {
	return this.Country
}
func (this *LoginByEmail) SetTicket(Ticket string) (result *LoginByEmail) {
	this.Ticket = Ticket
	return this
}
func (this *LoginByEmail) GetTicket() (Ticket string) {
	return this.Ticket
}
func (this *LoginByEmail) SetRandStr(RandStr string) (result *LoginByEmail) {
	this.RandStr = RandStr
	return this
}
func (this *LoginByEmail) GetRandStr() (RandStr string) {
	return this.RandStr
}
func (this *LoginByEmail) SetPromotion(Promotion string) (result *LoginByEmail) {
	this.Promotion = Promotion
	return this
}
func (this *LoginByEmail) GetPromotion() (Promotion string) {
	return this.Promotion
}

type LoginByEmail struct {
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	Username  string `gorm:"column:username" json:"username"`
	Country   string `gorm:"column:country" json:"country"`
	Ticket    string `gorm:"column:ticket" json:"ticket"`
	RandStr   string `gorm:"column:rand_str" json:"randStr"`
	Promotion string `gorm:"column:promotion" json:"promotion"`
}
