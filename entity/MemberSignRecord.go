package entity

func (this *MemberSignRecord) SetId(id int64) {
	this.Id = id
}
func (this *MemberSignRecord) GetId() (id int64) {
	return this.Id
}
func (this *MemberSignRecord) SetMember(member Member) {
	this.Member = member
}
func (this *MemberSignRecord) GetMember() (member Member) {
	return this.Member
}
func (this *MemberSignRecord) SetSign(sign Sign) {
	this.Sign = sign
}
func (this *MemberSignRecord) GetSign() (sign Sign) {
	return this.Sign
}
func (this *MemberSignRecord) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *MemberSignRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *MemberSignRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *MemberSignRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *MemberSignRecord) SetCreationTime(creationTime Date) {
	this.CreationTime = creationTime
}
func (this *MemberSignRecord) GetCreationTime() (creationTime Date) {
	return this.CreationTime
}
func NewMemberSignRecord() (this *MemberSignRecord) {
	this = new(MemberSignRecord)
	return
}
func NewMemberSignRecord(member Member, sign Sign) (this *MemberSignRecord) {
	this = new(MemberSignRecord)
	this.Member = member
	this.Sign = sign
	this.Coin = sign.GetCoin()
	//防止sign信息修改
	this.Amount = sign.GetAmount()
	//防止sign信息修改
	return
}

type MemberSignRecord struct {
	Id           int64
	Member       Member
	Sign         Sign
	Coin         Coin
	Amount       BigDecimal
	CreationTime Date
}
