package entity

func (this *MemberApplicationConfig) SetId(id int64) {
	this.Id = id
}
func (this *MemberApplicationConfig) GetId() (id int64) {
	return this.Id
}
func (this *MemberApplicationConfig) SetWithdrawCoinOn(withdrawCoinOn BooleanEnum) {
	this.WithdrawCoinOn = withdrawCoinOn
}
func (this *MemberApplicationConfig) GetWithdrawCoinOn() (withdrawCoinOn BooleanEnum) {
	return this.WithdrawCoinOn
}
func (this *MemberApplicationConfig) SetRechargeCoinOn(rechargeCoinOn BooleanEnum) {
	this.RechargeCoinOn = rechargeCoinOn
}
func (this *MemberApplicationConfig) GetRechargeCoinOn() (rechargeCoinOn BooleanEnum) {
	return this.RechargeCoinOn
}
func (this *MemberApplicationConfig) SetPromotionOn(promotionOn BooleanEnum) {
	this.PromotionOn = promotionOn
}
func (this *MemberApplicationConfig) GetPromotionOn() (promotionOn BooleanEnum) {
	return this.PromotionOn
}
func (this *MemberApplicationConfig) SetTransactionOn(transactionOn BooleanEnum) {
	this.TransactionOn = transactionOn
}
func (this *MemberApplicationConfig) GetTransactionOn() (transactionOn BooleanEnum) {
	return this.TransactionOn
}

type MemberApplicationConfig struct {
	Id             int64
	WithdrawCoinOn BooleanEnum
	RechargeCoinOn BooleanEnum
	PromotionOn    BooleanEnum
	TransactionOn  BooleanEnum
}
