package entity

func (this *ReleaseBalance) SetId(id int64) (result *ReleaseBalance) {
	this.Id = id
	return this
}
func (this *ReleaseBalance) GetId() (id int64) {
	return this.Id
}
func (this *ReleaseBalance) SetMemberId(memberId int64) (result *ReleaseBalance) {
	this.MemberId = memberId
	return this
}
func (this *ReleaseBalance) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *ReleaseBalance) SetCoinName(coinName string) (result *ReleaseBalance) {
	this.CoinName = coinName
	return this
}
func (this *ReleaseBalance) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *ReleaseBalance) SetCoinUnit(coinUnit string) (result *ReleaseBalance) {
	this.CoinUnit = coinUnit
	return this
}
func (this *ReleaseBalance) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *ReleaseBalance) SetRegisterTime(registerTime time.Time) (result *ReleaseBalance) {
	this.RegisterTime = registerTime
	return this
}
func (this *ReleaseBalance) GetRegisterTime() (registerTime time.Time) {
	return this.RegisterTime
}
func (this *ReleaseBalance) SetUserName(userName string) (result *ReleaseBalance) {
	this.UserName = userName
	return this
}
func (this *ReleaseBalance) GetUserName() (userName string) {
	return this.UserName
}
func (this *ReleaseBalance) SetPhone(phone string) (result *ReleaseBalance) {
	this.Phone = phone
	return this
}
func (this *ReleaseBalance) GetPhone() (phone string) {
	return this.Phone
}
func (this *ReleaseBalance) SetEmail(email string) (result *ReleaseBalance) {
	this.Email = email
	return this
}
func (this *ReleaseBalance) GetEmail() (email string) {
	return this.Email
}
func (this *ReleaseBalance) SetReleaseBalance(releaseBalance math.BigDecimal) (result *ReleaseBalance) {
	this.ReleaseBalance = releaseBalance
	return this
}
func (this *ReleaseBalance) GetReleaseBalance() (releaseBalance math.BigDecimal) {
	return this.ReleaseBalance
}
func (this *ReleaseBalance) SetReleaseState(releaseState string) (result *ReleaseBalance) {
	this.ReleaseState = releaseState
	return this
}
func (this *ReleaseBalance) GetReleaseState() (releaseState string) {
	return this.ReleaseState
}
func (this *ReleaseBalance) SetReleaseTime(releaseTime time.Time) (result *ReleaseBalance) {
	this.ReleaseTime = releaseTime
	return this
}
func (this *ReleaseBalance) GetReleaseTime() (releaseTime time.Time) {
	return this.ReleaseTime
}

type ReleaseBalance struct {
	Id             int64
	MemberId       int64
	CoinName       string
	CoinUnit       string
	RegisterTime   time.Time
	UserName       string
	Phone          string
	Email          string
	ReleaseBalance math.BigDecimal
	ReleaseState   string
	ReleaseTime    time.Time
}
