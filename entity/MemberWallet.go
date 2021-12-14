package entity

func (this *MemberWallet) SetId(id int64) {
	this.Id = id
}
func (this *MemberWallet) GetId() (id int64) {
	return this.Id
}
func (this *MemberWallet) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *MemberWallet) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberWallet) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *MemberWallet) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *MemberWallet) SetBalance(balance BigDecimal) {
	this.Balance = balance
}
func (this *MemberWallet) GetBalance() (balance BigDecimal) {
	return this.Balance
}
func (this *MemberWallet) SetFrozenBalance(frozenBalance BigDecimal) {
	this.FrozenBalance = frozenBalance
}
func (this *MemberWallet) GetFrozenBalance() (frozenBalance BigDecimal) {
	return this.FrozenBalance
}
func (this *MemberWallet) SetAddress(address string) {
	this.Address = address
}
func (this *MemberWallet) GetAddress() (address string) {
	return this.Address
}
func (this *MemberWallet) SetToken(token string) {
	this.Token = token
}
func (this *MemberWallet) GetToken() (token string) {
	return this.Token
}
func (this *MemberWallet) SetContract(contract string) {
	this.Contract = contract
}
func (this *MemberWallet) GetContract() (contract string) {
	return this.Contract
}
func (this *MemberWallet) SetVersion(version int) {
	this.Version = version
}
func (this *MemberWallet) GetVersion() (version int) {
	return this.Version
}
func (this *MemberWallet) SetIsLock(isLock BooleanEnum) {
	this.IsLock = isLock
}
func (this *MemberWallet) GetIsLock() (isLock BooleanEnum) {
	return this.IsLock
}
func (this *MemberWallet) SetReleaseBalance(releaseBalance BigDecimal) {
	this.ReleaseBalance = releaseBalance
}
func (this *MemberWallet) GetReleaseBalance() (releaseBalance BigDecimal) {
	return this.ReleaseBalance
}

type MemberWallet struct {
	Id             int64
	MemberId       int64
	Coin           Coin
	Balance        BigDecimal
	FrozenBalance  BigDecimal
	Address        string
	Token          string
	Contract       string
	Version        int
	IsLock         BooleanEnum
	ReleaseBalance BigDecimal
}
