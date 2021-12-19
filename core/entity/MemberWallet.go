package entity

func (this *MemberWallet) SetId(id int64) (result *MemberWallet) {
	this.Id = id
	return this
}
func (this *MemberWallet) GetId() (id int64) {
	return this.Id
}
func (this *MemberWallet) SetMemberId(memberId int64) (result *MemberWallet) {
	this.MemberId = memberId
	return this
}
func (this *MemberWallet) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberWallet) SetCoin(coin Coin) (result *MemberWallet) {
	this.Coin = coin
	return this
}
func (this *MemberWallet) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *MemberWallet) SetBalance(balance math.BigDecimal) (result *MemberWallet) {
	this.Balance = balance
	return this
}
func (this *MemberWallet) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *MemberWallet) SetFrozenBalance(frozenBalance math.BigDecimal) (result *MemberWallet) {
	this.FrozenBalance = frozenBalance
	return this
}
func (this *MemberWallet) GetFrozenBalance() (frozenBalance math.BigDecimal) {
	return this.FrozenBalance
}
func (this *MemberWallet) SetAddress(address string) (result *MemberWallet) {
	this.Address = address
	return this
}
func (this *MemberWallet) GetAddress() (address string) {
	return this.Address
}
func (this *MemberWallet) SetToken(token string) (result *MemberWallet) {
	this.Token = token
	return this
}
func (this *MemberWallet) GetToken() (token string) {
	return this.Token
}
func (this *MemberWallet) SetContract(contract string) (result *MemberWallet) {
	this.Contract = contract
	return this
}
func (this *MemberWallet) GetContract() (contract string) {
	return this.Contract
}
func (this *MemberWallet) SetVersion(version int) (result *MemberWallet) {
	this.Version = version
	return this
}
func (this *MemberWallet) GetVersion() (version int) {
	return this.Version
}
func (this *MemberWallet) SetIsLock(isLock constant.BooleanEnum) (result *MemberWallet) {
	this.IsLock = isLock
	return this
}
func (this *MemberWallet) GetIsLock() (isLock constant.BooleanEnum) {
	return this.IsLock
}
func (this *MemberWallet) SetReleaseBalance(releaseBalance math.BigDecimal) (result *MemberWallet) {
	this.ReleaseBalance = releaseBalance
	return this
}
func (this *MemberWallet) GetReleaseBalance() (releaseBalance math.BigDecimal) {
	return this.ReleaseBalance
}

type MemberWallet struct {
	Id             int64
	MemberId       int64
	Coin           Coin
	Balance        math.BigDecimal
	FrozenBalance  math.BigDecimal
	Address        string
	Token          string
	Contract       string
	Version        int
	IsLock         constant.BooleanEnum
	ReleaseBalance math.BigDecimal
}
