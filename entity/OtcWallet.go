package entity

var SerialVersionUID int64 = 1

func (this *OtcWallet) SetId(id int64) {
	this.Id = id
}
func (this *OtcWallet) GetId() (id int64) {
	return this.Id
}
func (this *OtcWallet) SetBalance(balance BigDecimal) {
	this.Balance = balance
}
func (this *OtcWallet) GetBalance() (balance BigDecimal) {
	return this.Balance
}
func (this *OtcWallet) SetFrozenBalance(frozenBalance BigDecimal) {
	this.FrozenBalance = frozenBalance
}
func (this *OtcWallet) GetFrozenBalance() (frozenBalance BigDecimal) {
	return this.FrozenBalance
}
func (this *OtcWallet) SetReleaseBalance(releaseBalance BigDecimal) {
	this.ReleaseBalance = releaseBalance
}
func (this *OtcWallet) GetReleaseBalance() (releaseBalance BigDecimal) {
	return this.ReleaseBalance
}
func (this *OtcWallet) SetIsLock(isLock int64) {
	this.IsLock = isLock
}
func (this *OtcWallet) GetIsLock() (isLock int64) {
	return this.IsLock
}
func (this *OtcWallet) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *OtcWallet) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *OtcWallet) SetVersion(version int64) {
	this.Version = version
}
func (this *OtcWallet) GetVersion() (version int64) {
	return this.Version
}
func (this *OtcWallet) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *OtcWallet) GetCoin() (coin Coin) {
	return this.Coin
}

type OtcWallet struct {
	Id             int64
	Balance        BigDecimal
	FrozenBalance  BigDecimal
	ReleaseBalance BigDecimal
	IsLock         int64
	MemberId       int64
	Version        int64
	Coin           Coin
}
