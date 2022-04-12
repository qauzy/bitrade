package dto

import "github.com/qauzy/math"

func (this *OtcWalletDTO) SetId(Id int64) (result *OtcWalletDTO) {
	this.Id = Id
	return this
}
func (this *OtcWalletDTO) GetId() (Id int64) {
	return this.Id
}
func (this *OtcWalletDTO) SetUnit(Unit string) (result *OtcWalletDTO) {
	this.Unit = Unit
	return this
}
func (this *OtcWalletDTO) GetUnit() (Unit string) {
	return this.Unit
}
func (this *OtcWalletDTO) SetBalance(Balance math.BigDecimal) (result *OtcWalletDTO) {
	this.Balance = Balance
	return this
}
func (this *OtcWalletDTO) GetBalance() (Balance math.BigDecimal) {
	return this.Balance
}
func (this *OtcWalletDTO) SetFrozenBalance(FrozenBalance math.BigDecimal) (result *OtcWalletDTO) {
	this.FrozenBalance = FrozenBalance
	return this
}
func (this *OtcWalletDTO) GetFrozenBalance() (FrozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *OtcWalletDTO) SetAllBalance(AllBalance math.BigDecimal) (result *OtcWalletDTO) {
	this.AllBalance = AllBalance
	return this
}
func (this *OtcWalletDTO) GetAllBalance() (AllBalance math.BigDecimal) {
	return this.AllBalance
}

type OtcWalletDTO struct {
	Id            int64           `gorm:"column:id" json:"id"`
	Unit          string          `gorm:"column:unit" json:"unit"`
	Balance       math.BigDecimal `gorm:"column:balance" json:"balance"`
	FrozenBalance math.BigDecimal `gorm:"column:frozen_balance" json:"frozenBalance"`
	AllBalance    math.BigDecimal `gorm:"column:all_balance" json:"allBalance"`
	BaseMemberDTO
}
