package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/math"
)

func (this *RewardWallet) SetId(id int64) (result *RewardWallet) {
	this.Id = id
	return this
}
func (this *RewardWallet) GetId() (id int64) {
	return this.Id
}
func (this *RewardWallet) SetMemberId(memberId int64) (result *RewardWallet) {
	this.MemberId = memberId
	return this
}
func (this *RewardWallet) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *RewardWallet) SetCoinUnit(coinUnit string) (result *RewardWallet) {
	this.CoinUnit = coinUnit
	return this
}
func (this *RewardWallet) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *RewardWallet) SetBalance(balance math.BigDecimal) (result *RewardWallet) {
	this.Balance = balance
	return this
}
func (this *RewardWallet) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *RewardWallet) SetVersion(version int) (result *RewardWallet) {
	this.Version = version
	return this
}
func (this *RewardWallet) GetVersion() (version int) {
	return this.Version
}
func (this *RewardWallet) SetFrozenBalance(frozenBalance math.BigDecimal) (result *RewardWallet) {
	this.FrozenBalance = frozenBalance
	return this
}
func (this *RewardWallet) GetFrozenBalance() (frozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *RewardWallet) SetIsLock(isLock BooleanEnum.BooleanEnum) (result *RewardWallet) {
	this.IsLock = isLock
	return this
}
func (this *RewardWallet) GetIsLock() (isLock BooleanEnum.BooleanEnum) {
	return this.IsLock
}

type RewardWallet struct {
	Id            int64
	MemberId      int64
	CoinUnit      string
	Balance       math.BigDecimal
	Version       int
	FrozenBalance math.BigDecimal
	IsLock        BooleanEnum.BooleanEnum
}
