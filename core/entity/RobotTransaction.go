package entity

func (this *RobotTransaction) SetId(id int64) (result *RobotTransaction) {
	this.Id = id
	return this
}
func (this *RobotTransaction) GetId() (id int64) {
	return this.Id
}
func (this *RobotTransaction) SetMemberId(memberId int64) (result *RobotTransaction) {
	this.MemberId = memberId
	return this
}
func (this *RobotTransaction) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *RobotTransaction) SetAmount(amount math.BigDecimal) (result *RobotTransaction) {
	this.Amount = amount
	return this
}
func (this *RobotTransaction) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *RobotTransaction) SetCreateTime(createTime time.Time) (result *RobotTransaction) {
	this.CreateTime = createTime
	return this
}
func (this *RobotTransaction) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *RobotTransaction) SetType(oType constant.TransactionType) (result *RobotTransaction) {
	this.Type = oType
	return this
}
func (this *RobotTransaction) GetType() (oType constant.TransactionType) {
	return this.Type
}
func (this *RobotTransaction) SetSymbol(symbol string) (result *RobotTransaction) {
	this.Symbol = symbol
	return this
}
func (this *RobotTransaction) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *RobotTransaction) SetFee(fee math.BigDecimal) (result *RobotTransaction) {
	this.Fee = fee
	return this
}
func (this *RobotTransaction) GetFee() (fee math.BigDecimal) {
	return this.Fee
}

type RobotTransaction struct {
	Id         int64
	MemberId   int64
	Amount     math.BigDecimal
	CreateTime time.Time
	Type       constant.TransactionType
	Symbol     string
	Fee        math.BigDecimal
}
