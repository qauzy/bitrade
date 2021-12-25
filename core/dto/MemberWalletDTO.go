package dto

func (this *MemberWalletDTO) SetId(id int64) (result *MemberWalletDTO) {
	this.Id = id
	return this
}
func (this *MemberWalletDTO) GetId() (id int64) {
	return this.Id
}
func (this *MemberWalletDTO) SetUnit(unit string) (result *MemberWalletDTO) {
	this.Unit = unit
	return this
}
func (this *MemberWalletDTO) GetUnit() (unit string) {
	return this.Unit
}
func (this *MemberWalletDTO) SetBalance(balance math.BigDecimal) (result *MemberWalletDTO) {
	this.Balance = balance
	return this
}
func (this *MemberWalletDTO) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *MemberWalletDTO) SetFrozenBalance(frozenBalance math.BigDecimal) (result *MemberWalletDTO) {
	this.FrozenBalance = frozenBalance
	return this
}
func (this *MemberWalletDTO) GetFrozenBalance() (frozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *MemberWalletDTO) SetAllBalance(allBalance math.BigDecimal) (result *MemberWalletDTO) {
	this.AllBalance = allBalance
	return this
}
func (this *MemberWalletDTO) GetAllBalance() (allBalance math.BigDecimal) {
	return this.AllBalance
}
func (this *MemberWalletDTO) SetAddress(address string) (result *MemberWalletDTO) {
	this.Address = address
	return this
}
func (this *MemberWalletDTO) GetAddress() (address string) {
	return this.Address
}

type MemberWalletDTO struct {
	Id            int64
	Unit          string
	Balance       math.BigDecimal
	FrozenBalance math.BigDecimal
	AllBalance    math.BigDecimal
	Address       string
}
