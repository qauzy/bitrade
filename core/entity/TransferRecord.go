package entity

func (this *TransferRecord) SetId(id int64) (result *TransferRecord) {
	this.Id = id
	return this
}
func (this *TransferRecord) GetId() (id int64) {
	return this.Id
}
func (this *TransferRecord) SetAmount(amount math.BigDecimal) (result *TransferRecord) {
	this.Amount = amount
	return this
}
func (this *TransferRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *TransferRecord) SetMemberId(memberId int64) (result *TransferRecord) {
	this.MemberId = memberId
	return this
}
func (this *TransferRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *TransferRecord) SetCoin(coin Coin) (result *TransferRecord) {
	this.Coin = coin
	return this
}
func (this *TransferRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *TransferRecord) SetCreateTime(createTime time.Time) (result *TransferRecord) {
	this.CreateTime = createTime
	return this
}
func (this *TransferRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *TransferRecord) SetFee(fee math.BigDecimal) (result *TransferRecord) {
	this.Fee = fee
	return this
}
func (this *TransferRecord) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *TransferRecord) SetOrderSn(orderSn string) (result *TransferRecord) {
	this.OrderSn = orderSn
	return this
}
func (this *TransferRecord) GetOrderSn() (orderSn string) {
	return this.OrderSn
}
func (this *TransferRecord) SetAddress(address string) (result *TransferRecord) {
	this.Address = address
	return this
}
func (this *TransferRecord) GetAddress() (address string) {
	return this.Address
}
func (this *TransferRecord) SetRemark(remark string) (result *TransferRecord) {
	this.Remark = remark
	return this
}
func (this *TransferRecord) GetRemark() (remark string) {
	return this.Remark
}

type TransferRecord struct {
	Id         int64
	Amount     math.BigDecimal
	MemberId   int64
	Coin       Coin
	CreateTime time.Time
	Fee        math.BigDecimal
	OrderSn    string
	Address    string
	Remark     string
}
