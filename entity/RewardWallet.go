package entity

func (this *RewardWallet) SetId(id int64) {
	this.Id = id
}
func (this *RewardWallet) GetId() (id int64) {
	return this.Id
}
func (this *RewardWallet) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *RewardWallet) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *RewardWallet) SetCoinUnit(coinUnit string) {
	this.CoinUnit = coinUnit
}
func (this *RewardWallet) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *RewardWallet) SetBalance(balance BigDecimal) {
	this.Balance = balance
}
func (this *RewardWallet) GetBalance() (balance BigDecimal) {
	return this.Balance
}
func (this *RewardWallet) SetVersion(version int) {
	this.Version = version
}
func (this *RewardWallet) GetVersion() (version int) {
	return this.Version
}
func (this *RewardWallet) SetFrozenBalance(frozenBalance BigDecimal) {
	this.FrozenBalance = frozenBalance
}
func (this *RewardWallet) GetFrozenBalance() (frozenBalance BigDecimal) {
	return this.FrozenBalance
}
func (this *RewardWallet) SetIsLock(isLock BooleanEnum) {
	this.IsLock = isLock
}
func (this *RewardWallet) GetIsLock() (isLock BooleanEnum) {
	return this.IsLock
}

type RewardWallet struct {
	Id            int64
	MemberId      int64
	CoinUnit      string
	Balance       BigDecimal
	Version       int
	FrozenBalance BigDecimal
	IsLock        BooleanEnum
}
