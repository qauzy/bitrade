package entity

import (
	"bitrade/core/constant/PayMode"
	"github.com/qauzy/math"
)

func (this *LegalWalletWithdrawModel) SetAmount(Amount *math.BigDecimal) (result *LegalWalletWithdrawModel) {
	this.Amount = Amount
	return this
}
func (this *LegalWalletWithdrawModel) GetAmount() (Amount *math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletWithdrawModel) SetUnit(Unit string) (result *LegalWalletWithdrawModel) {
	this.Unit = Unit
	return this
}
func (this *LegalWalletWithdrawModel) GetUnit() (Unit string) {
	return this.Unit
}
func (this *LegalWalletWithdrawModel) SetPayMode(PayMode *PayMode.PayMode) (result *LegalWalletWithdrawModel) {
	this.PayMode = PayMode
	return this
}
func (this *LegalWalletWithdrawModel) GetPayMode() (PayMode *PayMode.PayMode) {
	return this.PayMode
}
func (this *LegalWalletWithdrawModel) SetRemark(Remark string) (result *LegalWalletWithdrawModel) {
	this.Remark = Remark
	return this
}
func (this *LegalWalletWithdrawModel) GetRemark() (Remark string) {
	return this.Remark
}
func (this *LegalWalletWithdrawModel) SetAccount(Account string) (result *LegalWalletWithdrawModel) {
	this.Account = Account
	return this
}
func (this *LegalWalletWithdrawModel) GetAccount() (Account string) {
	return this.Account
}

type LegalWalletWithdrawModel struct {
	Amount  *math.BigDecimal `gorm:"column:amount" json:"amount"`
	Unit    string           `gorm:"column:unit" json:"unit"`
	PayMode *PayMode.PayMode `gorm:"column:pay_mode" json:"payMode"`
	Remark  string           `gorm:"column:remark" json:"remark"`
	Account string           `gorm:"column:account" json:"account"`
}
