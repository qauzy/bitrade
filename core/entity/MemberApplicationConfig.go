package entity

func (this *MemberApplicationConfig) SetId(id int64) (result *MemberApplicationConfig) {
	this.Id = id
	return this
}
func (this *MemberApplicationConfig) GetId() (id int64) {
	return this.Id
}
func (this *MemberApplicationConfig) SetWithdrawCoinOn(withdrawCoinOn constant.BooleanEnum) (result *MemberApplicationConfig) {
	this.WithdrawCoinOn = withdrawCoinOn
	return this
}
func (this *MemberApplicationConfig) GetWithdrawCoinOn() (withdrawCoinOn constant.BooleanEnum) {
	return this.WithdrawCoinOn
}
func (this *MemberApplicationConfig) SetRechargeCoinOn(rechargeCoinOn constant.BooleanEnum) (result *MemberApplicationConfig) {
	this.RechargeCoinOn = rechargeCoinOn
	return this
}
func (this *MemberApplicationConfig) GetRechargeCoinOn() (rechargeCoinOn constant.BooleanEnum) {
	return this.RechargeCoinOn
}
func (this *MemberApplicationConfig) SetPromotionOn(promotionOn constant.BooleanEnum) (result *MemberApplicationConfig) {
	this.PromotionOn = promotionOn
	return this
}
func (this *MemberApplicationConfig) GetPromotionOn() (promotionOn constant.BooleanEnum) {
	return this.PromotionOn
}
func (this *MemberApplicationConfig) SetTransactionOn(transactionOn constant.BooleanEnum) (result *MemberApplicationConfig) {
	this.TransactionOn = transactionOn
	return this
}
func (this *MemberApplicationConfig) GetTransactionOn() (transactionOn constant.BooleanEnum) {
	return this.TransactionOn
}

type MemberApplicationConfig struct {
	Id             int64
	WithdrawCoinOn constant.BooleanEnum
	RechargeCoinOn constant.BooleanEnum
	PromotionOn    constant.BooleanEnum
	TransactionOn  constant.BooleanEnum
}
