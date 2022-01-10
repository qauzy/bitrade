package entity

import (
	"bitrade/core/constant/RewardRecordType"
	"github.com/qauzy/math"
	"time"
)

func (this *RewardRecord) SetId(id int64) (result *RewardRecord) {
	this.Id = id
	return this
}
func (this *RewardRecord) GetId() (id int64) {
	return this.Id
}
func (this *RewardRecord) SetCoin(coin *Coin) (result *RewardRecord) {
	this.Coin = coin
	return this
}
func (this *RewardRecord) GetCoin() (coin *Coin) {
	return this.Coin
}
func (this *RewardRecord) SetRemark(remark string) (result *RewardRecord) {
	this.Remark = remark
	return this
}
func (this *RewardRecord) GetRemark() (remark string) {
	return this.Remark
}
func (this *RewardRecord) SetType(oType RewardRecordType.RewardRecordType) (result *RewardRecord) {
	this.Type = oType
	return this
}
func (this *RewardRecord) GetType() (oType RewardRecordType.RewardRecordType) {
	return this.Type
}
func (this *RewardRecord) SetAmount(amount math.BigDecimal) (result *RewardRecord) {
	this.Amount = amount
	return this
}
func (this *RewardRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *RewardRecord) SetMember(member *Member) (result *RewardRecord) {
	this.Member = member
	return this
}
func (this *RewardRecord) GetMember() (member *Member) {
	return this.Member
}
func (this *RewardRecord) SetCreateTime(createTime time.Time) (result *RewardRecord) {
	this.CreateTime = createTime
	return this
}
func (this *RewardRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type RewardRecord struct {
	Id         int64
	Coin       *Coin
	Remark     string
	Type       RewardRecordType.RewardRecordType
	Amount     math.BigDecimal
	Member     *Member
	CreateTime time.Time
}
