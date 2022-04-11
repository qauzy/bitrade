package entity

import "bitrade/core/constant/BooleanEnum"

func (this *MemberApplicationConfig) SetId(Id int64) (result *MemberApplicationConfig) {
	this.Id = Id
	return this
}
func (this *MemberApplicationConfig) GetId() (Id int64) {
	return this.Id
}
func (this *MemberApplicationConfig) SetWithdrawCoinOn(WithdrawCoinOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.WithdrawCoinOn = WithdrawCoinOn
	return this
}
func (this *MemberApplicationConfig) GetWithdrawCoinOn() (WithdrawCoinOn BooleanEnum.BooleanEnum) {
	return this.WithdrawCoinOn
}
func (this *MemberApplicationConfig) SetRechargeCoinOn(RechargeCoinOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.RechargeCoinOn = RechargeCoinOn
	return this
}
func (this *MemberApplicationConfig) GetRechargeCoinOn() (RechargeCoinOn BooleanEnum.BooleanEnum) {
	return this.RechargeCoinOn
}
func (this *MemberApplicationConfig) SetPromotionOn(PromotionOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.PromotionOn = PromotionOn
	return this
}
func (this *MemberApplicationConfig) GetPromotionOn() (PromotionOn BooleanEnum.BooleanEnum) {
	return this.PromotionOn
}
func (this *MemberApplicationConfig) SetTransactionOn(TransactionOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.TransactionOn = TransactionOn
	return this
}
func (this *MemberApplicationConfig) GetTransactionOn() (TransactionOn BooleanEnum.BooleanEnum) {
	return this.TransactionOn
}
func NewMemberApplicationConfig(id int64, withdrawCoinOn BooleanEnum.BooleanEnum, rechargeCoinOn BooleanEnum.BooleanEnum, promotionOn BooleanEnum.BooleanEnum, transactionOn BooleanEnum.BooleanEnum) (ret *MemberApplicationConfig) {
	ret = new(MemberApplicationConfig)
	ret.Id = id
	ret.WithdrawCoinOn = withdrawCoinOn
	ret.RechargeCoinOn = rechargeCoinOn
	ret.PromotionOn = promotionOn
	ret.TransactionOn = transactionOn
	return
}

type MemberApplicationConfig struct {
	Id             int64                   `gorm:"column:id" json:"id"`
	WithdrawCoinOn BooleanEnum.BooleanEnum `gorm:"column:withdraw_coin_on" json:"withdrawCoinOn"`
	RechargeCoinOn BooleanEnum.BooleanEnum `gorm:"column:recharge_coin_on" json:"rechargeCoinOn"`
	PromotionOn    BooleanEnum.BooleanEnum `gorm:"column:promotion_on" json:"promotionOn"`
	TransactionOn  BooleanEnum.BooleanEnum `gorm:"column:transaction_on" json:"transactionOn"`
}
