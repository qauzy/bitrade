package entity

import "github.com/qauzy/math"

func (this *OtcWallet) SetId(id int64) (result *OtcWallet) {
	this.Id = id
	return this
}
func (this *OtcWallet) GetId() (id int64) {
	return this.Id
}
func (this *OtcWallet) SetBalance(balance math.BigDecimal) (result *OtcWallet) {
	this.Balance = balance
	return this
}
func (this *OtcWallet) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *OtcWallet) SetFrozenBalance(frozenBalance math.BigDecimal) (result *OtcWallet) {
	this.FrozenBalance = frozenBalance
	return this
}
func (this *OtcWallet) GetFrozenBalance() (frozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *OtcWallet) SetReleaseBalance(releaseBalance math.BigDecimal) (result *OtcWallet) {
	this.ReleaseBalance = releaseBalance
	return this
}
func (this *OtcWallet) GetReleaseBalance() (releaseBalance math.BigDecimal) {
	return this.ReleaseBalance
}
func (this *OtcWallet) SetIsLock(isLock int64) (result *OtcWallet) {
	this.IsLock = isLock
	return this
}
func (this *OtcWallet) GetIsLock() (isLock int64) {
	return this.IsLock
}
func (this *OtcWallet) SetMemberId(memberId int64) (result *OtcWallet) {
	this.MemberId = memberId
	return this
}
func (this *OtcWallet) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *OtcWallet) SetVersion(version int64) (result *OtcWallet) {
	this.Version = version
	return this
}
func (this *OtcWallet) GetVersion() (version int64) {
	return this.Version
}
func (this *OtcWallet) SetCoin(coin Coin) (result *OtcWallet) {
	this.Coin = coin
	return this
}
func (this *OtcWallet) GetCoin() (coin Coin) {
	return this.Coin
}

type OtcWallet struct {
	Id             int64
	Balance        math.BigDecimal
	FrozenBalance  math.BigDecimal
	ReleaseBalance math.BigDecimal
	IsLock         int64
	MemberId       int64
	Version        int64
	Coin           Coin
}
