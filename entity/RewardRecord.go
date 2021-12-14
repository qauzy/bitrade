
package entity

func (this *RewardRecord) SetId(id int64) {
	this.Id = id
}
func (this *RewardRecord) GetId() (id int64) {
	return this.Id
}
func (this *RewardRecord) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *RewardRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *RewardRecord) SetRemark(remark string) {
	this.Remark = remark
}
func (this *RewardRecord) GetRemark() (remark string) {
	return this.Remark
}
func (this *RewardRecord) SetType(type RewardRecordType) {
	this.Type = type
}
func (this *RewardRecord) GetType() (type RewardRecordType) {
	return this.Type
}
func (this *RewardRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *RewardRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *RewardRecord) SetMember(member Member) {
	this.Member = member
}
func (this *RewardRecord) GetMember() (member Member) {
	return this.Member
}
func (this *RewardRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *RewardRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}

type RewardRecord struct {
	Id         int64
	Coin       Coin
	Remark     string
	Type       RewardRecordType
	Amount     BigDecimal
	Member     Member
	CreateTime Date
}

