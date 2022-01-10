package entity

import (
	"bitrade/core/constant/PayMode"
	"bitrade/core/constant/WithdrawStatus"
	"github.com/qauzy/math"
	"time"
)

func (this *LegalWalletWithdraw) SetId(id int64) (result *LegalWalletWithdraw) {
	this.Id = id
	return this
}
func (this *LegalWalletWithdraw) GetId() (id int64) {
	return this.Id
}
func (this *LegalWalletWithdraw) SetMember(member *Member) (result *LegalWalletWithdraw) {
	this.Member = member
	return this
}
func (this *LegalWalletWithdraw) GetMember() (member *Member) {
	return this.Member
}
func (this *LegalWalletWithdraw) SetAmount(amount math.BigDecimal) (result *LegalWalletWithdraw) {
	this.Amount = amount
	return this
}
func (this *LegalWalletWithdraw) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletWithdraw) SetCreateTime(createTime time.Time) (result *LegalWalletWithdraw) {
	this.CreateTime = createTime
	return this
}
func (this *LegalWalletWithdraw) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *LegalWalletWithdraw) SetDealTime(dealTime time.Time) (result *LegalWalletWithdraw) {
	this.DealTime = dealTime
	return this
}
func (this *LegalWalletWithdraw) GetDealTime() (dealTime time.Time) {
	return this.DealTime
}
func (this *LegalWalletWithdraw) SetStatus(status WithdrawStatus.WithdrawStatus) (result *LegalWalletWithdraw) {
	this.Status = status
	return this
}
func (this *LegalWalletWithdraw) GetStatus() (status WithdrawStatus.WithdrawStatus) {
	return this.Status
}
func (this *LegalWalletWithdraw) SetAdmin(admin *Admin) (result *LegalWalletWithdraw) {
	this.Admin = admin
	return this
}
func (this *LegalWalletWithdraw) GetAdmin() (admin *Admin) {
	return this.Admin
}
func (this *LegalWalletWithdraw) SetRemark(remark string) (result *LegalWalletWithdraw) {
	this.Remark = remark
	return this
}
func (this *LegalWalletWithdraw) GetRemark() (remark string) {
	return this.Remark
}
func (this *LegalWalletWithdraw) SetPayMode(payMode PayMode.PayMode) (result *LegalWalletWithdraw) {
	this.PayMode = payMode
	return this
}
func (this *LegalWalletWithdraw) GetPayMode() (payMode PayMode.PayMode) {
	return this.PayMode
}
func (this *LegalWalletWithdraw) SetCoin(coin *Coin) (result *LegalWalletWithdraw) {
	this.Coin = coin
	return this
}
func (this *LegalWalletWithdraw) GetCoin() (coin *Coin) {
	return this.Coin
}
func (this *LegalWalletWithdraw) SetPaymentInstrument(paymentInstrument string) (result *LegalWalletWithdraw) {
	this.PaymentInstrument = paymentInstrument
	return this
}
func (this *LegalWalletWithdraw) GetPaymentInstrument() (paymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletWithdraw) SetRemitTime(remitTime time.Time) (result *LegalWalletWithdraw) {
	this.RemitTime = remitTime
	return this
}
func (this *LegalWalletWithdraw) GetRemitTime() (remitTime time.Time) {
	return this.RemitTime
}
func (this *LegalWalletWithdraw) SetAccount(account string) (result *LegalWalletWithdraw) {
	this.Account = account
	return this
}
func (this *LegalWalletWithdraw) GetAccount() (account string) {
	return this.Account
}

type LegalWalletWithdraw struct {
	Id                int64
	Member            *Member
	Amount            math.BigDecimal
	CreateTime        time.Time
	DealTime          time.Time
	Status            WithdrawStatus.WithdrawStatus
	Admin             *Admin
	Remark            string
	PayMode           PayMode.PayMode
	Coin              *Coin
	PaymentInstrument string
	RemitTime         time.Time
	Account           string
}
