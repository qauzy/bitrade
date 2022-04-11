package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/math"
)

func (this *MemberWallet) SetId(Id int64) (result *MemberWallet) {
	this.Id = Id
	return this
}
func (this *MemberWallet) GetId() (Id int64) {
	return this.Id
}
func (this *MemberWallet) SetMemberId(MemberId int64) (result *MemberWallet) {
	this.MemberId = MemberId
	return this
}
func (this *MemberWallet) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *MemberWallet) SetCoin(Coin *Coin) (result *MemberWallet) {
	this.Coin = Coin
	return this
}
func (this *MemberWallet) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *MemberWallet) SetBalance(Balance math.BigDecimal) (result *MemberWallet) {
	this.Balance = Balance
	return this
}
func (this *MemberWallet) GetBalance() (Balance math.BigDecimal) {
	return this.Balance
}
func (this *MemberWallet) SetFrozenBalance(FrozenBalance math.BigDecimal) (result *MemberWallet) {
	this.FrozenBalance = FrozenBalance
	return this
}
func (this *MemberWallet) GetFrozenBalance() (FrozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *MemberWallet) SetAddress(Address string) (result *MemberWallet) {
	this.Address = Address
	return this
}
func (this *MemberWallet) GetAddress() (Address string) {
	return this.Address
}
func (this *MemberWallet) SetToken(Token string) (result *MemberWallet) {
	this.Token = Token
	return this
}
func (this *MemberWallet) GetToken() (Token string) {
	return this.Token
}
func (this *MemberWallet) SetContract(Contract string) (result *MemberWallet) {
	this.Contract = Contract
	return this
}
func (this *MemberWallet) GetContract() (Contract string) {
	return this.Contract
}
func (this *MemberWallet) SetVersion(Version int) (result *MemberWallet) {
	this.Version = Version
	return this
}
func (this *MemberWallet) GetVersion() (Version int) {
	return this.Version
}
func (this *MemberWallet) SetIsLock(IsLock BooleanEnum.BooleanEnum) (result *MemberWallet) {
	this.IsLock = IsLock
	return this
}
func (this *MemberWallet) GetIsLock() (IsLock BooleanEnum.BooleanEnum) {
	return this.IsLock
}
func (this *MemberWallet) SetReleaseBalance(ReleaseBalance math.BigDecimal) (result *MemberWallet) {
	this.ReleaseBalance = ReleaseBalance
	return this
}
func (this *MemberWallet) GetReleaseBalance() (ReleaseBalance math.BigDecimal) {
	return this.ReleaseBalance
}
func NewMemberWallet(id int64, memberId int64, coin *Coin, balance math.BigDecimal, frozenBalance math.BigDecimal, address string, token string, contract string, version int, isLock BooleanEnum.BooleanEnum, releaseBalance math.BigDecimal) (ret *MemberWallet) {
	ret = new(MemberWallet)
	ret.Id = id
	ret.MemberId = memberId
	ret.Coin = coin
	ret.Balance = balance
	ret.FrozenBalance = frozenBalance
	ret.Address = address
	ret.Token = token
	ret.Contract = contract
	ret.Version = version
	ret.IsLock = isLock
	ret.ReleaseBalance = releaseBalance
	return
}

type MemberWallet struct {
	Id             int64                   `gorm:"column:id" json:"id"`
	MemberId       int64                   `gorm:"column:member_id" json:"memberId"`
	Coin           *Coin                   `gorm:"column:coin" json:"coin"`
	Balance        math.BigDecimal         `gorm:"column:balance" json:"balance"`
	FrozenBalance  math.BigDecimal         `gorm:"column:frozen_balance" json:"frozenBalance"`
	Address        string                  `gorm:"column:address" json:"address"`
	Token          string                  `gorm:"column:token" json:"token"`
	Contract       string                  `gorm:"column:contract" json:"contract"`
	Version        int                     `gorm:"column:version" json:"version"`
	IsLock         BooleanEnum.BooleanEnum `gorm:"column:is_lock" json:"isLock"`
	ReleaseBalance math.BigDecimal         `gorm:"column:release_balance" json:"releaseBalance"`
}
