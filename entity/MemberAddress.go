package entity

func (this *MemberAddress) SetId(id int64) {
	this.Id = id
}
func (this *MemberAddress) GetId() (id int64) {
	return this.Id
}
func (this *MemberAddress) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberAddress) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *MemberAddress) SetDeleteTime(deleteTime Date) {
	this.DeleteTime = deleteTime
}
func (this *MemberAddress) GetDeleteTime() (deleteTime Date) {
	return this.DeleteTime
}
func (this *MemberAddress) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *MemberAddress) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *MemberAddress) SetAddress(address string) {
	this.Address = address
}
func (this *MemberAddress) GetAddress() (address string) {
	return this.Address
}
func (this *MemberAddress) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *MemberAddress) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *MemberAddress) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *MemberAddress) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberAddress) SetRemark(remark string) {
	this.Remark = remark
}
func (this *MemberAddress) GetRemark() (remark string) {
	return this.Remark
}

type MemberAddress struct {
	Id         int64
	CreateTime Date
	DeleteTime Date
	Coin       Coin
	Address    string
	Status     CommonStatus
	MemberId   int64
	Remark     string
}
