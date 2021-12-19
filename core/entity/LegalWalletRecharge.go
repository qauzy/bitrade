package entity

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
	"time"
)

func (this *LegalWalletRecharge) SetId(id int64) (result *LegalWalletRecharge) {
	this.Id = id
	return this
}
func (this *LegalWalletRecharge) GetId() (id int64) {
	return this.Id
}
func (this *LegalWalletRecharge) SetMember(member Member) (result *LegalWalletRecharge) {
	this.Member = member
	return this
}
func (this *LegalWalletRecharge) GetMember() (member Member) {
	return this.Member
}
func (this *LegalWalletRecharge) SetAmount(amount math.BigDecimal) (result *LegalWalletRecharge) {
	this.Amount = amount
	return this
}
func (this *LegalWalletRecharge) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletRecharge) SetPaymentInstrument(paymentInstrument string) (result *LegalWalletRecharge) {
	this.PaymentInstrument = paymentInstrument
	return this
}
func (this *LegalWalletRecharge) GetPaymentInstrument() (paymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletRecharge) SetCoin(coin Coin) (result *LegalWalletRecharge) {
	this.Coin = coin
	return this
}
func (this *LegalWalletRecharge) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *LegalWalletRecharge) SetState(state constant.LegalWalletState) (result *LegalWalletRecharge) {
	this.State = state
	return this
}
func (this *LegalWalletRecharge) GetState() (state constant.LegalWalletState) {
	return this.State
}
func (this *LegalWalletRecharge) SetPayMode(payMode constant.PayMode) (result *LegalWalletRecharge) {
	this.PayMode = payMode
	return this
}
func (this *LegalWalletRecharge) GetPayMode() (payMode constant.PayMode) {
	return this.PayMode
}
func (this *LegalWalletRecharge) SetRemark(remark string) (result *LegalWalletRecharge) {
	this.Remark = remark
	return this
}
func (this *LegalWalletRecharge) GetRemark() (remark string) {
	return this.Remark
}
func (this *LegalWalletRecharge) SetCreationTime(creationTime time.Time) (result *LegalWalletRecharge) {
	this.CreationTime = creationTime
	return this
}
func (this *LegalWalletRecharge) GetCreationTime() (creationTime time.Time) {
	return this.CreationTime
}
func (this *LegalWalletRecharge) SetDealTime(dealTime time.Time) (result *LegalWalletRecharge) {
	this.DealTime = dealTime
	return this
}
func (this *LegalWalletRecharge) GetDealTime() (dealTime time.Time) {
	return this.DealTime
}
func (this *LegalWalletRecharge) SetAdmin(admin Admin) (result *LegalWalletRecharge) {
	this.Admin = admin
	return this
}
func (this *LegalWalletRecharge) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *LegalWalletRecharge) SetUpdateTime(updateTime time.Time) (result *LegalWalletRecharge) {
	this.UpdateTime = updateTime
	return this
}
func (this *LegalWalletRecharge) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}

type LegalWalletRecharge struct {
	Id                int64
	Member            Member
	Amount            math.BigDecimal
	PaymentInstrument string
	Coin              Coin
	State             constant.LegalWalletState
	PayMode           constant.PayMode
	Remark            string
	CreationTime      time.Time
	DealTime          time.Time
	Admin             Admin
	UpdateTime        time.Time
}
