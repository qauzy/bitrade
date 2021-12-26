package entity

func (this *MemberTransaction) SetId(id int64) (result *MemberTransaction) {
	this.Id = id
	return this
}
func (this *MemberTransaction) GetId() (id int64) {
	return this.Id
}
func (this *MemberTransaction) SetMemberId(memberId int64) (result *MemberTransaction) {
	this.MemberId = memberId
	return this
}
func (this *MemberTransaction) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberTransaction) SetAmount(amount math.BigDecimal) (result *MemberTransaction) {
	this.Amount = amount
	return this
}
func (this *MemberTransaction) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberTransaction) SetCreateTime(createTime time.Time) (result *MemberTransaction) {
	this.CreateTime = createTime
	return this
}
func (this *MemberTransaction) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberTransaction) SetType(oType TransactionType.TransactionType) (result *MemberTransaction) {
	this.Type = oType
	return this
}
func (this *MemberTransaction) GetType() (oType TransactionType.TransactionType) {
	return this.Type
}
func (this *MemberTransaction) SetSymbol(symbol string) (result *MemberTransaction) {
	this.Symbol = symbol
	return this
}
func (this *MemberTransaction) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *MemberTransaction) SetAddress(address string) (result *MemberTransaction) {
	this.Address = address
	return this
}
func (this *MemberTransaction) GetAddress() (address string) {
	return this.Address
}
func (this *MemberTransaction) SetFee(fee math.BigDecimal) (result *MemberTransaction) {
	this.Fee = fee
	return this
}
func (this *MemberTransaction) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *MemberTransaction) SetFlag(flag int) (result *MemberTransaction) {
	this.Flag = flag
	return this
}
func (this *MemberTransaction) GetFlag() (flag int) {
	return this.Flag
}
func (this *MemberTransaction) SetAirdropId(airdropId int64) (result *MemberTransaction) {
	this.AirdropId = airdropId
	return this
}
func (this *MemberTransaction) GetAirdropId() (airdropId int64) {
	return this.AirdropId
}

type MemberTransaction struct {
	Id         int64
	MemberId   int64
	Amount     math.BigDecimal
	CreateTime time.Time
	Type       TransactionType.TransactionType
	Symbol     string
	Address    string
	Fee        math.BigDecimal
	Flag       int
	AirdropId  int64
}
