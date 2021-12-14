
package entity

func (this *RobotTransaction) SetId(id int64) {
	this.Id = id
}
func (this *RobotTransaction) GetId() (id int64) {
	return this.Id
}
func (this *RobotTransaction) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *RobotTransaction) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *RobotTransaction) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *RobotTransaction) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *RobotTransaction) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *RobotTransaction) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *RobotTransaction) SetType(type TransactionType) {
	this.Type = type
}
func (this *RobotTransaction) GetType() (type TransactionType) {
	return this.Type
}
func (this *RobotTransaction) SetSymbol(symbol string) {
	this.Symbol = symbol
}
func (this *RobotTransaction) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *RobotTransaction) SetFee(fee BigDecimal) {
	this.Fee = fee
}
func (this *RobotTransaction) GetFee() (fee BigDecimal) {
	return this.Fee
}

type RobotTransaction struct {
	Id         int64
	MemberId   int64
	Amount     BigDecimal
	CreateTime Date
	Type       TransactionType
	Symbol     string
	Fee        BigDecimal
}

