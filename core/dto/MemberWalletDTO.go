package dto

import "github.com/qauzy/math"

func (this *MemberWalletDTO) SetId(Id int64) (result *MemberWalletDTO) {
	this.Id = Id
	return this
}
func (this *MemberWalletDTO) GetId() (Id int64) {
	return this.Id
}
func (this *MemberWalletDTO) SetUnit(Unit string) (result *MemberWalletDTO) {
	this.Unit = Unit
	return this
}
func (this *MemberWalletDTO) GetUnit() (Unit string) {
	return this.Unit
}
func (this *MemberWalletDTO) SetBalance(Balance math.BigDecimal) (result *MemberWalletDTO) {
	this.Balance = Balance
	return this
}
func (this *MemberWalletDTO) GetBalance() (Balance math.BigDecimal) {
	return this.Balance
}
func (this *MemberWalletDTO) SetFrozenBalance(FrozenBalance math.BigDecimal) (result *MemberWalletDTO) {
	this.FrozenBalance = FrozenBalance
	return this
}
func (this *MemberWalletDTO) GetFrozenBalance() (FrozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *MemberWalletDTO) SetAllBalance(AllBalance math.BigDecimal) (result *MemberWalletDTO) {
	this.AllBalance = AllBalance
	return this
}
func (this *MemberWalletDTO) GetAllBalance() (AllBalance math.BigDecimal) {
	return this.AllBalance
}
func (this *MemberWalletDTO) SetAddress(Address string) (result *MemberWalletDTO) {
	this.Address = Address
	return this
}
func (this *MemberWalletDTO) GetAddress() (Address string) {
	return this.Address
}

type MemberWalletDTO struct {
	Id            int64           `gorm:"column:id" json:"id"`
	Unit          string          `gorm:"column:unit" json:"unit"`
	Balance       math.BigDecimal `gorm:"column:balance" json:"balance"`
	FrozenBalance math.BigDecimal `gorm:"column:frozen_balance" json:"frozenBalance"`
	AllBalance    math.BigDecimal `gorm:"column:all_balance" json:"allBalance"`
	Address       string          `gorm:"column:address" json:"address"`
	BaseMemberDTO
}
