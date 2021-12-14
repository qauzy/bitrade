package entity

func (this *ReleaseBalance) SetId(id int64) {
	this.Id = id
}
func (this *ReleaseBalance) GetId() (id int64) {
	return this.Id
}
func (this *ReleaseBalance) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *ReleaseBalance) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *ReleaseBalance) SetCoinName(coinName string) {
	this.CoinName = coinName
}
func (this *ReleaseBalance) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *ReleaseBalance) SetCoinUnit(coinUnit string) {
	this.CoinUnit = coinUnit
}
func (this *ReleaseBalance) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *ReleaseBalance) SetRegisterTime(registerTime Date) {
	this.RegisterTime = registerTime
}
func (this *ReleaseBalance) GetRegisterTime() (registerTime Date) {
	return this.RegisterTime
}
func (this *ReleaseBalance) SetUserName(userName string) {
	this.UserName = userName
}
func (this *ReleaseBalance) GetUserName() (userName string) {
	return this.UserName
}
func (this *ReleaseBalance) SetPhone(phone string) {
	this.Phone = phone
}
func (this *ReleaseBalance) GetPhone() (phone string) {
	return this.Phone
}
func (this *ReleaseBalance) SetEmail(email string) {
	this.Email = email
}
func (this *ReleaseBalance) GetEmail() (email string) {
	return this.Email
}
func (this *ReleaseBalance) SetReleaseBalance(releaseBalance BigDecimal) {
	this.ReleaseBalance = releaseBalance
}
func (this *ReleaseBalance) GetReleaseBalance() (releaseBalance BigDecimal) {
	return this.ReleaseBalance
}
func (this *ReleaseBalance) SetReleaseState(releaseState string) {
	this.ReleaseState = releaseState
}
func (this *ReleaseBalance) GetReleaseState() (releaseState string) {
	return this.ReleaseState
}
func (this *ReleaseBalance) SetReleaseTime(releaseTime Date) {
	this.ReleaseTime = releaseTime
}
func (this *ReleaseBalance) GetReleaseTime() (releaseTime Date) {
	return this.ReleaseTime
}

type ReleaseBalance struct {
	Id             int64
	MemberId       int64
	CoinName       string
	CoinUnit       string
	RegisterTime   Date
	UserName       string
	Phone          string
	Email          string
	ReleaseBalance BigDecimal
	ReleaseState   string
	ReleaseTime    Date
}
