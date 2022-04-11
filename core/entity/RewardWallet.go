package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/math"
)

func (this *RewardWallet) SetId(Id int64) (result *RewardWallet) {
	this.Id = Id
	return this
}
func (this *RewardWallet) GetId() (Id int64) {
	return this.Id
}
func (this *RewardWallet) SetMemberId(MemberId int64) (result *RewardWallet) {
	this.MemberId = MemberId
	return this
}
func (this *RewardWallet) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *RewardWallet) SetCoinUnit(CoinUnit string) (result *RewardWallet) {
	this.CoinUnit = CoinUnit
	return this
}
func (this *RewardWallet) GetCoinUnit() (CoinUnit string) {
	return this.CoinUnit
}
func (this *RewardWallet) SetBalance(Balance math.BigDecimal) (result *RewardWallet) {
	this.Balance = Balance
	return this
}
func (this *RewardWallet) GetBalance() (Balance math.BigDecimal) {
	return this.Balance
}
func (this *RewardWallet) SetVersion(Version int) (result *RewardWallet) {
	this.Version = Version
	return this
}
func (this *RewardWallet) GetVersion() (Version int) {
	return this.Version
}
func (this *RewardWallet) SetFrozenBalance(FrozenBalance math.BigDecimal) (result *RewardWallet) {
	this.FrozenBalance = FrozenBalance
	return this
}
func (this *RewardWallet) GetFrozenBalance() (FrozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *RewardWallet) SetIsLock(IsLock BooleanEnum.BooleanEnum) (result *RewardWallet) {
	this.IsLock = IsLock
	return this
}
func (this *RewardWallet) GetIsLock() (IsLock BooleanEnum.BooleanEnum) {
	return this.IsLock
}
func NewRewardWallet(id int64, memberId int64, coinUnit string, balance math.BigDecimal, version int, frozenBalance math.BigDecimal, isLock BooleanEnum.BooleanEnum) (ret *RewardWallet) {
	ret = new(RewardWallet)
	ret.Id = id
	ret.MemberId = memberId
	ret.CoinUnit = coinUnit
	ret.Balance = balance
	ret.Version = version
	ret.FrozenBalance = frozenBalance
	ret.IsLock = isLock
	return
}

type RewardWallet struct {
	Id            int64                   `gorm:"column:id" json:"id"`
	MemberId      int64                   `gorm:"column:member_id" json:"memberId"`
	CoinUnit      string                  `gorm:"column:coin_unit" json:"coinUnit"`
	Balance       math.BigDecimal         `gorm:"column:balance" json:"balance"`
	Version       int                     `gorm:"column:version" json:"version"`
	FrozenBalance math.BigDecimal         `gorm:"column:frozen_balance" json:"frozenBalance"`
	IsLock        BooleanEnum.BooleanEnum `gorm:"column:is_lock" json:"isLock"`
}
