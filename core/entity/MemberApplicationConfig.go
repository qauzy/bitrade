package entity

func (this *MemberApplicationConfig) SetId(id int64) (result *MemberApplicationConfig) {
	this.Id = id
	return this
}
func (this *MemberApplicationConfig) GetId() (id int64) {
	return this.Id
}
func (this *MemberApplicationConfig) SetWithdrawCoinOn(withdrawCoinOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.WithdrawCoinOn = withdrawCoinOn
	return this
}
func (this *MemberApplicationConfig) GetWithdrawCoinOn() (withdrawCoinOn BooleanEnum.BooleanEnum) {
	return this.WithdrawCoinOn
}
func (this *MemberApplicationConfig) SetRechargeCoinOn(rechargeCoinOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.RechargeCoinOn = rechargeCoinOn
	return this
}
func (this *MemberApplicationConfig) GetRechargeCoinOn() (rechargeCoinOn BooleanEnum.BooleanEnum) {
	return this.RechargeCoinOn
}
func (this *MemberApplicationConfig) SetPromotionOn(promotionOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.PromotionOn = promotionOn
	return this
}
func (this *MemberApplicationConfig) GetPromotionOn() (promotionOn BooleanEnum.BooleanEnum) {
	return this.PromotionOn
}
func (this *MemberApplicationConfig) SetTransactionOn(transactionOn BooleanEnum.BooleanEnum) (result *MemberApplicationConfig) {
	this.TransactionOn = transactionOn
	return this
}
func (this *MemberApplicationConfig) GetTransactionOn() (transactionOn BooleanEnum.BooleanEnum) {
	return this.TransactionOn
}

type MemberApplicationConfig struct {
	Id             int64
	WithdrawCoinOn BooleanEnum.BooleanEnum
	RechargeCoinOn BooleanEnum.BooleanEnum
	PromotionOn    BooleanEnum.BooleanEnum
	TransactionOn  BooleanEnum.BooleanEnum
}
