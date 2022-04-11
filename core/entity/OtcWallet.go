package entity

import "github.com/qauzy/math"

func (this *OtcWallet) SetId(Id int64) (result *OtcWallet) {
	this.Id = Id
	return this
}
func (this *OtcWallet) GetId() (Id int64) {
	return this.Id
}
func (this *OtcWallet) SetBalance(Balance math.BigDecimal) (result *OtcWallet) {
	this.Balance = Balance
	return this
}
func (this *OtcWallet) GetBalance() (Balance math.BigDecimal) {
	return this.Balance
}
func (this *OtcWallet) SetFrozenBalance(FrozenBalance math.BigDecimal) (result *OtcWallet) {
	this.FrozenBalance = FrozenBalance
	return this
}
func (this *OtcWallet) GetFrozenBalance() (FrozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *OtcWallet) SetReleaseBalance(ReleaseBalance math.BigDecimal) (result *OtcWallet) {
	this.ReleaseBalance = ReleaseBalance
	return this
}
func (this *OtcWallet) GetReleaseBalance() (ReleaseBalance math.BigDecimal) {
	return this.ReleaseBalance
}
func (this *OtcWallet) SetIsLock(IsLock int) (result *OtcWallet) {
	this.IsLock = IsLock
	return this
}
func (this *OtcWallet) GetIsLock() (IsLock int) {
	return this.IsLock
}
func (this *OtcWallet) SetMemberId(MemberId int64) (result *OtcWallet) {
	this.MemberId = MemberId
	return this
}
func (this *OtcWallet) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *OtcWallet) SetVersion(Version int) (result *OtcWallet) {
	this.Version = Version
	return this
}
func (this *OtcWallet) GetVersion() (Version int) {
	return this.Version
}
func (this *OtcWallet) SetCoin(Coin *Coin) (result *OtcWallet) {
	this.Coin = Coin
	return this
}
func (this *OtcWallet) GetCoin() (Coin *Coin) {
	return this.Coin
}
func NewOtcWallet(id int64, balance math.BigDecimal, frozenBalance math.BigDecimal, releaseBalance math.BigDecimal, isLock int, memberId int64, version int, coin *Coin) (ret *OtcWallet) {
	ret = new(OtcWallet)
	ret.Id = id
	ret.Balance = balance
	ret.FrozenBalance = frozenBalance
	ret.ReleaseBalance = releaseBalance
	ret.IsLock = isLock
	ret.MemberId = memberId
	ret.Version = version
	ret.Coin = coin
	return
}

type OtcWallet struct {
	Id             int64           `gorm:"column:id" json:"id"`
	Balance        math.BigDecimal `gorm:"column:balance" json:"balance"`
	FrozenBalance  math.BigDecimal `gorm:"column:frozen_balance" json:"frozenBalance"`
	ReleaseBalance math.BigDecimal `gorm:"column:release_balance" json:"releaseBalance"`
	IsLock         int             `gorm:"column:is_lock" json:"isLock"`
	MemberId       int64           `gorm:"column:member_id" json:"memberId"`
	Version        int             `gorm:"column:version" json:"version"`
	Coin           *Coin           `gorm:"column:coin" json:"coin"`
}
