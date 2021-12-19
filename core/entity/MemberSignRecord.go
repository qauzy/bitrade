package entity

func (this *MemberSignRecord) SetId(id int64) (result *MemberSignRecord) {
	this.Id = id
	return this
}
func (this *MemberSignRecord) GetId() (id int64) {
	return this.Id
}
func (this *MemberSignRecord) SetMember(member Member) (result *MemberSignRecord) {
	this.Member = member
	return this
}
func (this *MemberSignRecord) GetMember() (member Member) {
	return this.Member
}
func (this *MemberSignRecord) SetSign(sign Sign) (result *MemberSignRecord) {
	this.Sign = sign
	return this
}
func (this *MemberSignRecord) GetSign() (sign Sign) {
	return this.Sign
}
func (this *MemberSignRecord) SetCoin(coin Coin) (result *MemberSignRecord) {
	this.Coin = coin
	return this
}
func (this *MemberSignRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *MemberSignRecord) SetAmount(amount math.BigDecimal) (result *MemberSignRecord) {
	this.Amount = amount
	return this
}
func (this *MemberSignRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberSignRecord) SetCreationTime(creationTime time.Time) (result *MemberSignRecord) {
	this.CreationTime = creationTime
	return this
}
func (this *MemberSignRecord) GetCreationTime() (creationTime time.Time) {
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
	Amount       math.BigDecimal
	CreationTime time.Time
}
