package entity

func (this *BusinessAuthDeposit) SetId(id int64) {
	this.Id = id
}
func (this *BusinessAuthDeposit) GetId() (id int64) {
	return this.Id
}
func (this *BusinessAuthDeposit) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *BusinessAuthDeposit) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *BusinessAuthDeposit) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *BusinessAuthDeposit) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthDeposit) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *BusinessAuthDeposit) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *BusinessAuthDeposit) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *BusinessAuthDeposit) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *BusinessAuthDeposit) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *BusinessAuthDeposit) GetStatus() (status CommonStatus) {
	return this.Status
}

type BusinessAuthDeposit struct {
	Id         int64
	Coin       Coin
	Amount     BigDecimal
	CreateTime Date
	Admin      Admin
	Status     CommonStatus
}
