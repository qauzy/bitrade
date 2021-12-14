package entity

func (this *TransferRecord) SetId(id int64) {
	this.Id = id
}
func (this *TransferRecord) GetId() (id int64) {
	return this.Id
}
func (this *TransferRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *TransferRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *TransferRecord) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *TransferRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *TransferRecord) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *TransferRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *TransferRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *TransferRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *TransferRecord) SetFee(fee BigDecimal) {
	this.Fee = fee
}
func (this *TransferRecord) GetFee() (fee BigDecimal) {
	return this.Fee
}
func (this *TransferRecord) SetOrderSn(orderSn string) {
	this.OrderSn = orderSn
}
func (this *TransferRecord) GetOrderSn() (orderSn string) {
	return this.OrderSn
}
func (this *TransferRecord) SetAddress(address string) {
	this.Address = address
}
func (this *TransferRecord) GetAddress() (address string) {
	return this.Address
}
func (this *TransferRecord) SetRemark(remark string) {
	this.Remark = remark
}
func (this *TransferRecord) GetRemark() (remark string) {
	return this.Remark
}

type TransferRecord struct {
	Id         int64
	Amount     BigDecimal
	MemberId   int64
	Coin       Coin
	CreateTime Date
	Fee        BigDecimal
	OrderSn    string
	Address    string
	Remark     string
}
