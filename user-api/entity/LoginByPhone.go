package entity

func (this *LoginByPhone) SetPhone(phone string) (result *LoginByPhone) {
	this.Phone = phone
	return this
}
func (this *LoginByPhone) GetPhone() (phone string) {
	return this.Phone
}
func (this *LoginByPhone) SetPassword(password string) (result *LoginByPhone) {
	this.Password = password
	return this
}
func (this *LoginByPhone) GetPassword() (password string) {
	return this.Password
}
func (this *LoginByPhone) SetUsername(username string) (result *LoginByPhone) {
	this.Username = username
	return this
}
func (this *LoginByPhone) GetUsername() (username string) {
	return this.Username
}
func (this *LoginByPhone) SetCountry(country string) (result *LoginByPhone) {
	this.Country = country
	return this
}
func (this *LoginByPhone) GetCountry() (country string) {
	return this.Country
}
func (this *LoginByPhone) SetCode(code string) (result *LoginByPhone) {
	this.Code = code
	return this
}
func (this *LoginByPhone) GetCode() (code string) {
	return this.Code
}
func (this *LoginByPhone) SetTicket(ticket string) (result *LoginByPhone) {
	this.Ticket = ticket
	return this
}
func (this *LoginByPhone) GetTicket() (ticket string) {
	return this.Ticket
}
func (this *LoginByPhone) SetRandStr(randStr string) (result *LoginByPhone) {
	this.RandStr = randStr
	return this
}
func (this *LoginByPhone) GetRandStr() (randStr string) {
	return this.RandStr
}
func (this *LoginByPhone) SetPromotion(promotion string) (result *LoginByPhone) {
	this.Promotion = promotion
	return this
}
func (this *LoginByPhone) GetPromotion() (promotion string) {
	return this.Promotion
}

type LoginByPhone struct {
	Phone     string
	Password  string
	Username  string
	Country   string
	Code      string
	Ticket    string
	RandStr   string
	Promotion string
}
