package dto

func (this *OtcWalletDTO) SetId(id int64) (result *OtcWalletDTO) {
	this.Id = id
	return this
}
func (this *OtcWalletDTO) GetId() (id int64) {
	return this.Id
}
func (this *OtcWalletDTO) SetUnit(unit string) (result *OtcWalletDTO) {
	this.Unit = unit
	return this
}
func (this *OtcWalletDTO) GetUnit() (unit string) {
	return this.Unit
}
func (this *OtcWalletDTO) SetBalance(balance math.BigDecimal) (result *OtcWalletDTO) {
	this.Balance = balance
	return this
}
func (this *OtcWalletDTO) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *OtcWalletDTO) SetFrozenBalance(frozenBalance math.BigDecimal) (result *OtcWalletDTO) {
	this.FrozenBalance = frozenBalance
	return this
}
func (this *OtcWalletDTO) GetFrozenBalance() (frozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *OtcWalletDTO) SetAllBalance(allBalance math.BigDecimal) (result *OtcWalletDTO) {
	this.AllBalance = allBalance
	return this
}
func (this *OtcWalletDTO) GetAllBalance() (allBalance math.BigDecimal) {
	return this.AllBalance
}

type OtcWalletDTO struct {
	Id            int64
	Unit          string
	Balance       math.BigDecimal
	FrozenBalance math.BigDecimal
	AllBalance    math.BigDecimal
}
