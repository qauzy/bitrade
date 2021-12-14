package entity

func (this *LegalWalletRecharge) SetId(id int64) {
	this.Id = id
}
func (this *LegalWalletRecharge) GetId() (id int64) {
	return this.Id
}
func (this *LegalWalletRecharge) SetMember(member Member) {
	this.Member = member
}
func (this *LegalWalletRecharge) GetMember() (member Member) {
	return this.Member
}
func (this *LegalWalletRecharge) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *LegalWalletRecharge) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *LegalWalletRecharge) SetPaymentInstrument(paymentInstrument string) {
	this.PaymentInstrument = paymentInstrument
}
func (this *LegalWalletRecharge) GetPaymentInstrument() (paymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletRecharge) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *LegalWalletRecharge) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *LegalWalletRecharge) SetState(state LegalWalletState) {
	this.State = state
}
func (this *LegalWalletRecharge) GetState() (state LegalWalletState) {
	return this.State
}
func (this *LegalWalletRecharge) SetPayMode(payMode PayMode) {
	this.PayMode = payMode
}
func (this *LegalWalletRecharge) GetPayMode() (payMode PayMode) {
	return this.PayMode
}
func (this *LegalWalletRecharge) SetRemark(remark string) {
	this.Remark = remark
}
func (this *LegalWalletRecharge) GetRemark() (remark string) {
	return this.Remark
}
func (this *LegalWalletRecharge) SetCreationTime(creationTime Date) {
	this.CreationTime = creationTime
}
func (this *LegalWalletRecharge) GetCreationTime() (creationTime Date) {
	return this.CreationTime
}
func (this *LegalWalletRecharge) SetDealTime(dealTime Date) {
	this.DealTime = dealTime
}
func (this *LegalWalletRecharge) GetDealTime() (dealTime Date) {
	return this.DealTime
}
func (this *LegalWalletRecharge) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *LegalWalletRecharge) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *LegalWalletRecharge) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *LegalWalletRecharge) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}

type LegalWalletRecharge struct {
	Id                int64
	Member            Member
	Amount            BigDecimal
	PaymentInstrument string
	Coin              Coin
	State             LegalWalletState
	PayMode           PayMode
	Remark            string
	CreationTime      Date
	DealTime          Date
	Admin             Admin
	UpdateTime        Date
}
