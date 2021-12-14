package entity

func (this *Sign) SetId(id int64) {
	this.Id = id
}
func (this *Sign) GetId() (id int64) {
	return this.Id
}
func (this *Sign) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *Sign) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *Sign) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *Sign) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *Sign) SetEndDate(endDate Date) {
	this.EndDate = endDate
}
func (this *Sign) GetEndDate() (endDate Date) {
	return this.EndDate
}
func (this *Sign) SetCreationTime(creationTime Date) {
	this.CreationTime = creationTime
}
func (this *Sign) GetCreationTime() (creationTime Date) {
	return this.CreationTime
}
func (this *Sign) SetStatus(status SignStatus) {
	this.Status = status
}
func (this *Sign) GetStatus() (status SignStatus) {
	return this.Status
}

type Sign struct {
	Id           int64
	Coin         Coin
	Amount       BigDecimal
	EndDate      Date
	CreationTime Date
	Status       SignStatus
}
