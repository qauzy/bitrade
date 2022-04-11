package entity

import (
	"bitrade/core/constant/RewardRecordType"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *RewardRecord) SetId(Id int64) (result *RewardRecord) {
	this.Id = Id
	return this
}
func (this *RewardRecord) GetId() (Id int64) {
	return this.Id
}
func (this *RewardRecord) SetCoin(Coin *Coin) (result *RewardRecord) {
	this.Coin = Coin
	return this
}
func (this *RewardRecord) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *RewardRecord) SetRemark(Remark string) (result *RewardRecord) {
	this.Remark = Remark
	return this
}
func (this *RewardRecord) GetRemark() (Remark string) {
	return this.Remark
}
func (this *RewardRecord) SetType(Type RewardRecordType.RewardRecordType) (result *RewardRecord) {
	this.Type = Type
	return this
}
func (this *RewardRecord) GetType() (Type RewardRecordType.RewardRecordType) {
	return this.Type
}
func (this *RewardRecord) SetAmount(Amount math.BigDecimal) (result *RewardRecord) {
	this.Amount = Amount
	return this
}
func (this *RewardRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *RewardRecord) SetMember(Member *Member) (result *RewardRecord) {
	this.Member = Member
	return this
}
func (this *RewardRecord) GetMember() (Member *Member) {
	return this.Member
}
func (this *RewardRecord) SetCreateTime(CreateTime xtime.Xtime) (result *RewardRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *RewardRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func NewRewardRecord(id int64, coin *Coin, remark string, oType RewardRecordType.RewardRecordType, amount math.BigDecimal, member *Member, createTime xtime.Xtime) (ret *RewardRecord) {
	ret = new(RewardRecord)
	ret.Id = id
	ret.Coin = coin
	ret.Remark = remark
	ret.Type = oType
	ret.Amount = amount
	ret.Member = member
	ret.CreateTime = createTime
	return
}

type RewardRecord struct {
	Id         int64                             `gorm:"column:id" json:"id"`
	Coin       *Coin                             `gorm:"column:coin" json:"coin"`
	Remark     string                            `gorm:"column:remark" json:"remark"`
	Type       RewardRecordType.RewardRecordType `gorm:"column:type" json:"type"`
	Amount     math.BigDecimal                   `gorm:"column:amount" json:"amount"`
	Member     *Member                           `gorm:"column:member" json:"member"`
	CreateTime xtime.Xtime                       `gorm:"column:create_time" json:"createTime"`
}
