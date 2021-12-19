package entity

func (this *MemberDeposit) SetId(id int64) (result *MemberDeposit) {
	this.Id = id
	return this
}
func (this *MemberDeposit) GetId() (id int64) {
	return this.Id
}
func (this *MemberDeposit) SetMemberId(memberId int64) (result *MemberDeposit) {
	this.MemberId = memberId
	return this
}
func (this *MemberDeposit) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberDeposit) SetAmount(amount math.BigDecimal) (result *MemberDeposit) {
	this.Amount = amount
	return this
}
func (this *MemberDeposit) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberDeposit) SetUnit(unit string) (result *MemberDeposit) {
	this.Unit = unit
	return this
}
func (this *MemberDeposit) GetUnit() (unit string) {
	return this.Unit
}
func (this *MemberDeposit) SetCreateTime(createTime time.Time) (result *MemberDeposit) {
	this.CreateTime = createTime
	return this
}
func (this *MemberDeposit) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberDeposit) SetTxid(txid string) (result *MemberDeposit) {
	this.Txid = txid
	return this
}
func (this *MemberDeposit) GetTxid() (txid string) {
	return this.Txid
}
func (this *MemberDeposit) SetAddress(address string) (result *MemberDeposit) {
	this.Address = address
	return this
}
func (this *MemberDeposit) GetAddress() (address string) {
	return this.Address
}
func (this *MemberDeposit) SetUsername(username string) (result *MemberDeposit) {
	this.Username = username
	return this
}
func (this *MemberDeposit) GetUsername() (username string) {
	return this.Username
}

type MemberDeposit struct {
	Id         int64
	MemberId   int64
	Amount     math.BigDecimal
	Unit       string
	CreateTime time.Time
	Txid       string
	Address    string
	Username   string
}
