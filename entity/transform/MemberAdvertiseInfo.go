package transform

func (this *MemberAdvertiseInfo) SetBuy(buy []ScanAdvertise) {
	this.Buy = buy
}
func (this *MemberAdvertiseInfo) GetBuy() (buy []ScanAdvertise) {
	return this.Buy
}
func (this *MemberAdvertiseInfo) SetSell(sell []ScanAdvertise) {
	this.Sell = sell
}
func (this *MemberAdvertiseInfo) GetSell() (sell []ScanAdvertise) {
	return this.Sell
}
func (this *MemberAdvertiseInfo) SetUsername(username string) {
	this.Username = username
}
func (this *MemberAdvertiseInfo) GetUsername() (username string) {
	return this.Username
}
func (this *MemberAdvertiseInfo) SetAvatar(avatar string) {
	this.Avatar = avatar
}
func (this *MemberAdvertiseInfo) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *MemberAdvertiseInfo) SetRealVerified(realVerified BooleanEnum) {
	this.RealVerified = realVerified
}
func (this *MemberAdvertiseInfo) GetRealVerified() (realVerified BooleanEnum) {
	return this.RealVerified
}
func (this *MemberAdvertiseInfo) SetEmailVerified(emailVerified BooleanEnum) {
	this.EmailVerified = emailVerified
}
func (this *MemberAdvertiseInfo) GetEmailVerified() (emailVerified BooleanEnum) {
	return this.EmailVerified
}
func (this *MemberAdvertiseInfo) SetPhoneVerified(phoneVerified BooleanEnum) {
	this.PhoneVerified = phoneVerified
}
func (this *MemberAdvertiseInfo) GetPhoneVerified() (phoneVerified BooleanEnum) {
	return this.PhoneVerified
}
func (this *MemberAdvertiseInfo) SetTransactions(transactions int) {
	this.Transactions = transactions
}
func (this *MemberAdvertiseInfo) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *MemberAdvertiseInfo) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberAdvertiseInfo) GetCreateTime() (createTime Date) {
	return this.CreateTime
}

type MemberAdvertiseInfo struct {
	Buy           []ScanAdvertise
	Sell          []ScanAdvertise
	Username      string
	Avatar        string
	RealVerified  BooleanEnum
	EmailVerified BooleanEnum
	PhoneVerified BooleanEnum
	Transactions  int
	CreateTime    Date
}
