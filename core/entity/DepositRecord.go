package entity

func (this *DepositRecord) SetId(id string) (result *DepositRecord) {
	this.Id = id
	return this
}
func (this *DepositRecord) GetId() (id string) {
	return this.Id
}
func (this *DepositRecord) SetMember(member Member) (result *DepositRecord) {
	this.Member = member
	return this
}
func (this *DepositRecord) GetMember() (member Member) {
	return this.Member
}
func (this *DepositRecord) SetCoin(coin Coin) (result *DepositRecord) {
	this.Coin = coin
	return this
}
func (this *DepositRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *DepositRecord) SetAmount(amount math.BigDecimal) (result *DepositRecord) {
	this.Amount = amount
	return this
}
func (this *DepositRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *DepositRecord) SetStatus(status DepositStatusEnum.DepositStatusEnum) (result *DepositRecord) {
	this.Status = status
	return this
}
func (this *DepositRecord) GetStatus() (status DepositStatusEnum.DepositStatusEnum) {
	return this.Status
}

type DepositRecord struct {
	Id     string
	Member Member
	Coin   Coin
	Amount math.BigDecimal
	Status DepositStatusEnum.DepositStatusEnum
}
