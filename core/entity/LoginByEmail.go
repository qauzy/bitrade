package entity

func (this *LoginByEmail) SetEmail(email string) (result *LoginByEmail) {
	this.Email = email
	return this
}
func (this *LoginByEmail) GetEmail() (email string) {
	return this.Email
}
func (this *LoginByEmail) SetPassword(password string) (result *LoginByEmail) {
	this.Password = password
	return this
}
func (this *LoginByEmail) GetPassword() (password string) {
	return this.Password
}
func (this *LoginByEmail) SetUsername(username string) (result *LoginByEmail) {
	this.Username = username
	return this
}
func (this *LoginByEmail) GetUsername() (username string) {
	return this.Username
}
func (this *LoginByEmail) SetCountry(country string) (result *LoginByEmail) {
	this.Country = country
	return this
}
func (this *LoginByEmail) GetCountry() (country string) {
	return this.Country
}
func (this *LoginByEmail) SetTicket(ticket string) (result *LoginByEmail) {
	this.Ticket = ticket
	return this
}
func (this *LoginByEmail) GetTicket() (ticket string) {
	return this.Ticket
}
func (this *LoginByEmail) SetRandStr(randStr string) (result *LoginByEmail) {
	this.RandStr = randStr
	return this
}
func (this *LoginByEmail) GetRandStr() (randStr string) {
	return this.RandStr
}
func (this *LoginByEmail) SetPromotion(promotion string) (result *LoginByEmail) {
	this.Promotion = promotion
	return this
}
func (this *LoginByEmail) GetPromotion() (promotion string) {
	return this.Promotion
}

type LoginByEmail struct {
	Email     string
	Password  string
	Username  string
	Country   string
	Ticket    string
	RandStr   string
	Promotion string
}
