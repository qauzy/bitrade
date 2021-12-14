package entity

func (this *LegalWalletWithdraw) SetId(id int64) {
	this.Id = id
}
func (this *LegalWalletWithdraw) GetId() (id int64) {
	return this.Id
}
func (this *LegalWalletWithdraw) SetMember(member Member) {
	this.Member = member
}
func (this *LegalWalletWithdraw) GetMember() (member Member) {
	return this.Member
}
func (this *LegalWalletWithdraw) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *LegalWalletWithdraw) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *LegalWalletWithdraw) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *LegalWalletWithdraw) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *LegalWalletWithdraw) SetDealTime(dealTime Date) {
	this.DealTime = dealTime
}
func (this *LegalWalletWithdraw) GetDealTime() (dealTime Date) {
	return this.DealTime
}
func (this *LegalWalletWithdraw) SetStatus(status WithdrawStatus) {
	this.Status = status
}
func (this *LegalWalletWithdraw) GetStatus() (status WithdrawStatus) {
	return this.Status
}
func (this *LegalWalletWithdraw) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *LegalWalletWithdraw) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *LegalWalletWithdraw) SetRemark(remark string) {
	this.Remark = remark
}
func (this *LegalWalletWithdraw) GetRemark() (remark string) {
	return this.Remark
}
func (this *LegalWalletWithdraw) SetPayMode(payMode PayMode) {
	this.PayMode = payMode
}
func (this *LegalWalletWithdraw) GetPayMode() (payMode PayMode) {
	return this.PayMode
}
func (this *LegalWalletWithdraw) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *LegalWalletWithdraw) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *LegalWalletWithdraw) SetPaymentInstrument(paymentInstrument string) {
	this.PaymentInstrument = paymentInstrument
}
func (this *LegalWalletWithdraw) GetPaymentInstrument() (paymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletWithdraw) SetRemitTime(remitTime Date) {
	this.RemitTime = remitTime
}
func (this *LegalWalletWithdraw) GetRemitTime() (remitTime Date) {
	return this.RemitTime
}
func (this *LegalWalletWithdraw) SetAccount(account string) {
	this.Account = account
}
func (this *LegalWalletWithdraw) GetAccount() (account string) {
	return this.Account
}

type LegalWalletWithdraw struct {
	Id                int64
	Member            Member
	Amount            BigDecimal
	CreateTime        Date
	DealTime          Date
	Status            WithdrawStatus
	Admin             Admin
	Remark            string
	PayMode           PayMode
	Coin              Coin
	PaymentInstrument string
	RemitTime         Date
	Account           string
}
