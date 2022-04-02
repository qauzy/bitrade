package entity

import (
	"bitrade/core/constant/PayMode"
	"github.com/qauzy/math"
)

func (this *LegalWalletRechargeModel) SetAmount(Amount *math.BigDecimal) (result *LegalWalletRechargeModel) {
	this.Amount = Amount
	return this
}
func (this *LegalWalletRechargeModel) GetAmount() (Amount *math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletRechargeModel) SetPaymentInstrument(PaymentInstrument string) (result *LegalWalletRechargeModel) {
	this.PaymentInstrument = PaymentInstrument
	return this
}
func (this *LegalWalletRechargeModel) GetPaymentInstrument() (PaymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletRechargeModel) SetUnit(Unit string) (result *LegalWalletRechargeModel) {
	this.Unit = Unit
	return this
}
func (this *LegalWalletRechargeModel) GetUnit() (Unit string) {
	return this.Unit
}
func (this *LegalWalletRechargeModel) SetPayMode(PayMode *PayMode.PayMode) (result *LegalWalletRechargeModel) {
	this.PayMode = PayMode
	return this
}
func (this *LegalWalletRechargeModel) GetPayMode() (PayMode *PayMode.PayMode) {
	return this.PayMode
}
func (this *LegalWalletRechargeModel) SetRemark(Remark string) (result *LegalWalletRechargeModel) {
	this.Remark = Remark
	return this
}
func (this *LegalWalletRechargeModel) GetRemark() (Remark string) {
	return this.Remark
}

type LegalWalletRechargeModel struct {
	Amount            *math.BigDecimal `gorm:"column:amount" json:"amount"`
	PaymentInstrument string           `gorm:"column:payment_instrument" json:"paymentInstrument"`
	Unit              string           `gorm:"column:unit" json:"unit"`
	PayMode           *PayMode.PayMode `gorm:"column:pay_mode" json:"payMode"`
	Remark            string           `gorm:"column:remark" json:"remark"`
}
