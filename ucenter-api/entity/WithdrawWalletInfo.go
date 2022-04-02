package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/math"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/maps/hashmap"
)

func (this *WithdrawWalletInfo) SetUnit(Unit string) (result *WithdrawWalletInfo) {
	this.Unit = Unit
	return this
}
func (this *WithdrawWalletInfo) GetUnit() (Unit string) {
	return this.Unit
}
func (this *WithdrawWalletInfo) SetThreshold(Threshold *math.BigDecimal) (result *WithdrawWalletInfo) {
	this.Threshold = Threshold
	return this
}
func (this *WithdrawWalletInfo) GetThreshold() (Threshold *math.BigDecimal) {
	return this.Threshold
}
func (this *WithdrawWalletInfo) SetMinAmount(MinAmount *math.BigDecimal) (result *WithdrawWalletInfo) {
	this.MinAmount = MinAmount
	return this
}
func (this *WithdrawWalletInfo) GetMinAmount() (MinAmount *math.BigDecimal) {
	return this.MinAmount
}
func (this *WithdrawWalletInfo) SetMaxAmount(MaxAmount *math.BigDecimal) (result *WithdrawWalletInfo) {
	this.MaxAmount = MaxAmount
	return this
}
func (this *WithdrawWalletInfo) GetMaxAmount() (MaxAmount *math.BigDecimal) {
	return this.MaxAmount
}
func (this *WithdrawWalletInfo) SetMinTxFee(MinTxFee float64) (result *WithdrawWalletInfo) {
	this.MinTxFee = MinTxFee
	return this
}
func (this *WithdrawWalletInfo) GetMinTxFee() (MinTxFee float64) {
	return this.MinTxFee
}
func (this *WithdrawWalletInfo) SetMaxTxFee(MaxTxFee float64) (result *WithdrawWalletInfo) {
	this.MaxTxFee = MaxTxFee
	return this
}
func (this *WithdrawWalletInfo) GetMaxTxFee() (MaxTxFee float64) {
	return this.MaxTxFee
}
func (this *WithdrawWalletInfo) SetNameCn(NameCn string) (result *WithdrawWalletInfo) {
	this.NameCn = NameCn
	return this
}
func (this *WithdrawWalletInfo) GetNameCn() (NameCn string) {
	return this.NameCn
}
func (this *WithdrawWalletInfo) SetName(Name string) (result *WithdrawWalletInfo) {
	this.Name = Name
	return this
}
func (this *WithdrawWalletInfo) GetName() (Name string) {
	return this.Name
}
func (this *WithdrawWalletInfo) SetBalance(Balance *math.BigDecimal) (result *WithdrawWalletInfo) {
	this.Balance = Balance
	return this
}
func (this *WithdrawWalletInfo) GetBalance() (Balance *math.BigDecimal) {
	return this.Balance
}
func (this *WithdrawWalletInfo) SetCanAutoWithdraw(CanAutoWithdraw *BooleanEnum.BooleanEnum) (result *WithdrawWalletInfo) {
	this.CanAutoWithdraw = CanAutoWithdraw
	return this
}
func (this *WithdrawWalletInfo) GetCanAutoWithdraw() (CanAutoWithdraw *BooleanEnum.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *WithdrawWalletInfo) SetWithdrawScale(WithdrawScale int) (result *WithdrawWalletInfo) {
	this.WithdrawScale = WithdrawScale
	return this
}
func (this *WithdrawWalletInfo) GetWithdrawScale() (WithdrawScale int) {
	return this.WithdrawScale
}
func (this *WithdrawWalletInfo) SetAddresses(Addresses *arraylist.List[*hashmap.Map[string, string]]) (result *WithdrawWalletInfo) {
	this.Addresses = Addresses
	return this
}
func (this *WithdrawWalletInfo) GetAddresses() (Addresses *arraylist.List[*hashmap.Map[string, string]]) {
	return this.Addresses
}

type WithdrawWalletInfo struct {
	Unit            string                                        `gorm:"column:unit" json:"unit"`
	Threshold       *math.BigDecimal                              `gorm:"column:threshold" json:"threshold"`
	MinAmount       *math.BigDecimal                              `gorm:"column:min_amount" json:"minAmount"`
	MaxAmount       *math.BigDecimal                              `gorm:"column:max_amount" json:"maxAmount"`
	MinTxFee        float64                                       `gorm:"column:min_tx_fee" json:"minTxFee"`
	MaxTxFee        float64                                       `gorm:"column:max_tx_fee" json:"maxTxFee"`
	NameCn          string                                        `gorm:"column:name_cn" json:"nameCn"`
	Name            string                                        `gorm:"column:name" json:"name"`
	Balance         *math.BigDecimal                              `gorm:"column:balance" json:"balance"`
	CanAutoWithdraw *BooleanEnum.BooleanEnum                      `gorm:"column:can_auto_withdraw" json:"canAutoWithdraw"`
	WithdrawScale   int                                           `gorm:"column:withdraw_scale" json:"withdrawScale"`
	Addresses       *arraylist.List[*hashmap.Map[string, string]] `gorm:"column:addresses" json:"addresses"`
}
