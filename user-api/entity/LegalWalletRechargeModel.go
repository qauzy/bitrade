package entity

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
)

func (this *LegalWalletRechargeModel) SetAmount(amount math.BigDecimal) (result *LegalWalletRechargeModel) {
	this.Amount = amount
	return this
}
func (this *LegalWalletRechargeModel) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletRechargeModel) SetPaymentInstrument(paymentInstrument string) (result *LegalWalletRechargeModel) {
	this.PaymentInstrument = paymentInstrument
	return this
}
func (this *LegalWalletRechargeModel) GetPaymentInstrument() (paymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletRechargeModel) SetUnit(unit string) (result *LegalWalletRechargeModel) {
	this.Unit = unit
	return this
}
func (this *LegalWalletRechargeModel) GetUnit() (unit string) {
	return this.Unit
}
func (this *LegalWalletRechargeModel) SetPayMode(payMode constant.PayMode) (result *LegalWalletRechargeModel) {
	this.PayMode = payMode
	return this
}
func (this *LegalWalletRechargeModel) GetPayMode() (payMode constant.PayMode) {
	return this.PayMode
}
func (this *LegalWalletRechargeModel) SetRemark(remark string) (result *LegalWalletRechargeModel) {
	this.Remark = remark
	return this
}
func (this *LegalWalletRechargeModel) GetRemark() (remark string) {
	return this.Remark
}

type LegalWalletRechargeModel struct {
	Amount            math.BigDecimal
	PaymentInstrument string
	Unit              string
	PayMode           constant.PayMode
	Remark            string
}
