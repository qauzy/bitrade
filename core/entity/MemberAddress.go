package entity

func (this *MemberAddress) SetId(id int64) (result *MemberAddress) {
	this.Id = id
	return this
}
func (this *MemberAddress) GetId() (id int64) {
	return this.Id
}
func (this *MemberAddress) SetCreateTime(createTime time.Time) (result *MemberAddress) {
	this.CreateTime = createTime
	return this
}
func (this *MemberAddress) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberAddress) SetDeleteTime(deleteTime time.Time) (result *MemberAddress) {
	this.DeleteTime = deleteTime
	return this
}
func (this *MemberAddress) GetDeleteTime() (deleteTime time.Time) {
	return this.DeleteTime
}
func (this *MemberAddress) SetCoin(coin Coin) (result *MemberAddress) {
	this.Coin = coin
	return this
}
func (this *MemberAddress) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *MemberAddress) SetAddress(address string) (result *MemberAddress) {
	this.Address = address
	return this
}
func (this *MemberAddress) GetAddress() (address string) {
	return this.Address
}
func (this *MemberAddress) SetStatus(status CommonStatus.CommonStatus) (result *MemberAddress) {
	this.Status = status
	return this
}
func (this *MemberAddress) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *MemberAddress) SetMemberId(memberId int64) (result *MemberAddress) {
	this.MemberId = memberId
	return this
}
func (this *MemberAddress) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberAddress) SetRemark(remark string) (result *MemberAddress) {
	this.Remark = remark
	return this
}
func (this *MemberAddress) GetRemark() (remark string) {
	return this.Remark
}

type MemberAddress struct {
	Id         int64
	CreateTime time.Time
	DeleteTime time.Time
	Coin       Coin
	Address    string
	Status     CommonStatus.CommonStatus
	MemberId   int64
	Remark     string
}
