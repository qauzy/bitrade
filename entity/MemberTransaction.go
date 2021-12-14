
package entity

func (this *MemberTransaction) SetId(id int64) {
	this.Id = id
}
func (this *MemberTransaction) GetId() (id int64) {
	return this.Id
}
func (this *MemberTransaction) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *MemberTransaction) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberTransaction) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *MemberTransaction) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *MemberTransaction) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberTransaction) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *MemberTransaction) SetType(type TransactionType) {
	this.Type = type
}
func (this *MemberTransaction) GetType() (type TransactionType) {
	return this.Type
}
func (this *MemberTransaction) SetSymbol(symbol string) {
	this.Symbol = symbol
}
func (this *MemberTransaction) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *MemberTransaction) SetAddress(address string) {
	this.Address = address
}
func (this *MemberTransaction) GetAddress() (address string) {
	return this.Address
}
func (this *MemberTransaction) SetFee(fee BigDecimal) {
	this.Fee = fee
}
func (this *MemberTransaction) GetFee() (fee BigDecimal) {
	return this.Fee
}
func (this *MemberTransaction) SetFlag(flag int) {
	this.Flag = flag
}
func (this *MemberTransaction) GetFlag() (flag int) {
	return this.Flag
}
func (this *MemberTransaction) SetAirdropId(airdropId int64) {
	this.AirdropId = airdropId
}
func (this *MemberTransaction) GetAirdropId() (airdropId int64) {
	return this.AirdropId
}

type MemberTransaction struct {
	Id         int64
	MemberId   int64
	Amount     BigDecimal
	CreateTime Date
	Type       TransactionType
	Symbol     string
	Address    string
	Fee        BigDecimal
	Flag       int
	AirdropId  int64
}

