package entity

func (this *WithdrawWalletInfo) SetUnit(unit string) (result *WithdrawWalletInfo) {
	this.Unit = unit
	return this
}
func (this *WithdrawWalletInfo) GetUnit() (unit string) {
	return this.Unit
}
func (this *WithdrawWalletInfo) SetThreshold(threshold math.BigDecimal) (result *WithdrawWalletInfo) {
	this.Threshold = threshold
	return this
}
func (this *WithdrawWalletInfo) GetThreshold() (threshold math.BigDecimal) {
	return this.Threshold
}
func (this *WithdrawWalletInfo) SetMinAmount(minAmount math.BigDecimal) (result *WithdrawWalletInfo) {
	this.MinAmount = minAmount
	return this
}
func (this *WithdrawWalletInfo) GetMinAmount() (minAmount math.BigDecimal) {
	return this.MinAmount
}
func (this *WithdrawWalletInfo) SetMaxAmount(maxAmount math.BigDecimal) (result *WithdrawWalletInfo) {
	this.MaxAmount = maxAmount
	return this
}
func (this *WithdrawWalletInfo) GetMaxAmount() (maxAmount math.BigDecimal) {
	return this.MaxAmount
}
func (this *WithdrawWalletInfo) SetMinTxFee(minTxFee float64) (result *WithdrawWalletInfo) {
	this.MinTxFee = minTxFee
	return this
}
func (this *WithdrawWalletInfo) GetMinTxFee() (minTxFee float64) {
	return this.MinTxFee
}
func (this *WithdrawWalletInfo) SetMaxTxFee(maxTxFee float64) (result *WithdrawWalletInfo) {
	this.MaxTxFee = maxTxFee
	return this
}
func (this *WithdrawWalletInfo) GetMaxTxFee() (maxTxFee float64) {
	return this.MaxTxFee
}
func (this *WithdrawWalletInfo) SetNameCn(nameCn string) (result *WithdrawWalletInfo) {
	this.NameCn = nameCn
	return this
}
func (this *WithdrawWalletInfo) GetNameCn() (nameCn string) {
	return this.NameCn
}
func (this *WithdrawWalletInfo) SetName(name string) (result *WithdrawWalletInfo) {
	this.Name = name
	return this
}
func (this *WithdrawWalletInfo) GetName() (name string) {
	return this.Name
}
func (this *WithdrawWalletInfo) SetBalance(balance math.BigDecimal) (result *WithdrawWalletInfo) {
	this.Balance = balance
	return this
}
func (this *WithdrawWalletInfo) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *WithdrawWalletInfo) SetCanAutoWithdraw(canAutoWithdraw constant.BooleanEnum) (result *WithdrawWalletInfo) {
	this.CanAutoWithdraw = canAutoWithdraw
	return this
}
func (this *WithdrawWalletInfo) GetCanAutoWithdraw() (canAutoWithdraw constant.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *WithdrawWalletInfo) SetWithdrawScale(withdrawScale int) (result *WithdrawWalletInfo) {
	this.WithdrawScale = withdrawScale
	return this
}
func (this *WithdrawWalletInfo) GetWithdrawScale() (withdrawScale int) {
	return this.WithdrawScale
}
func (this *WithdrawWalletInfo) SetAddresses(addresses []map[string]string) (result *WithdrawWalletInfo) {
	this.Addresses = addresses
	return this
}
func (this *WithdrawWalletInfo) GetAddresses() (addresses []map[string]string) {
	return this.Addresses
}

type WithdrawWalletInfo struct {
	Unit            string
	Threshold       math.BigDecimal
	MinAmount       math.BigDecimal
	MaxAmount       math.BigDecimal
	MinTxFee        float64
	MaxTxFee        float64
	NameCn          string
	Name            string
	Balance         math.BigDecimal
	CanAutoWithdraw constant.BooleanEnum
	WithdrawScale   int
	Addresses       []map[string]string
}
