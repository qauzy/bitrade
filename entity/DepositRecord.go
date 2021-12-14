package entity

func (this *DepositRecord) SetId(id string) {
	this.Id = id
}
func (this *DepositRecord) GetId() (id string) {
	return this.Id
}
func (this *DepositRecord) SetMember(member Member) {
	this.Member = member
}
func (this *DepositRecord) GetMember() (member Member) {
	return this.Member
}
func (this *DepositRecord) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *DepositRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *DepositRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *DepositRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *DepositRecord) SetStatus(status DepositStatusEnum) {
	this.Status = status
}
func (this *DepositRecord) GetStatus() (status DepositStatusEnum) {
	return this.Status
}

type DepositRecord struct {
	Id     string
	Member Member
	Coin   Coin
	Amount BigDecimal
	Status DepositStatusEnum
}
