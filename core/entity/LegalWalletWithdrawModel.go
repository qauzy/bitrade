package entity

import (
	"bitrade/core/constant/PayMode"
	"github.com/qauzy/math"
)

func (this *LegalWalletWithdrawModel) SetAmount(amount math.BigDecimal) (result *LegalWalletWithdrawModel) {
	this.Amount = amount
	return this
}
func (this *LegalWalletWithdrawModel) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletWithdrawModel) SetUnit(unit string) (result *LegalWalletWithdrawModel) {
	this.Unit = unit
	return this
}
func (this *LegalWalletWithdrawModel) GetUnit() (unit string) {
	return this.Unit
}
func (this *LegalWalletWithdrawModel) SetPayMode(payMode PayMode.PayMode) (result *LegalWalletWithdrawModel) {
	this.PayMode = payMode
	return this
}
func (this *LegalWalletWithdrawModel) GetPayMode() (payMode PayMode.PayMode) {
	return this.PayMode
}
func (this *LegalWalletWithdrawModel) SetRemark(remark string) (result *LegalWalletWithdrawModel) {
	this.Remark = remark
	return this
}
func (this *LegalWalletWithdrawModel) GetRemark() (remark string) {
	return this.Remark
}
func (this *LegalWalletWithdrawModel) SetAccount(account string) (result *LegalWalletWithdrawModel) {
	this.Account = account
	return this
}
func (this *LegalWalletWithdrawModel) GetAccount() (account string) {
	return this.Account
}

type LegalWalletWithdrawModel struct {
	Amount  math.BigDecimal
	Unit    string
	PayMode PayMode.PayMode
	Remark  string
	Account string
}
