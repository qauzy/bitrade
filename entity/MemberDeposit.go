package entity

func (this *MemberDeposit) SetId(id int64) {
	this.Id = id
}
func (this *MemberDeposit) GetId() (id int64) {
	return this.Id
}
func (this *MemberDeposit) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *MemberDeposit) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberDeposit) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *MemberDeposit) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *MemberDeposit) SetUnit(unit string) {
	this.Unit = unit
}
func (this *MemberDeposit) GetUnit() (unit string) {
	return this.Unit
}
func (this *MemberDeposit) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberDeposit) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *MemberDeposit) SetTxid(txid string) {
	this.Txid = txid
}
func (this *MemberDeposit) GetTxid() (txid string) {
	return this.Txid
}
func (this *MemberDeposit) SetAddress(address string) {
	this.Address = address
}
func (this *MemberDeposit) GetAddress() (address string) {
	return this.Address
}
func (this *MemberDeposit) SetUsername(username string) {
	this.Username = username
}
func (this *MemberDeposit) GetUsername() (username string) {
	return this.Username
}

type MemberDeposit struct {
	Id         int64
	MemberId   int64
	Amount     BigDecimal
	Unit       string
	CreateTime Date
	Txid       string
	Address    string
	Username   string
}
