
package entity

import (
	"bitrade/core/constant/IntegrationRecordType"
	"github.com/qauzy/chocolate/xtime"
)

func (this *IntegrationRecord) SetId(Id int64) (result *IntegrationRecord) {
	this.Id = Id
	return this
}
func (this *IntegrationRecord) GetId() (Id int64) {
	return this.Id
}
func (this *IntegrationRecord) SetMemberId(MemberId int64) (result *IntegrationRecord) {
	this.MemberId = MemberId
	return this
}
func (this *IntegrationRecord) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *IntegrationRecord) SetType(Type IntegrationRecordType.IntegrationRecordType) (result *IntegrationRecord) {
	this.Type = Type
	return this
}
func (this *IntegrationRecord) GetType() (Type IntegrationRecordType.IntegrationRecordType) {
	return this.Type
}
func (this *IntegrationRecord) SetAmount(Amount int64) (result *IntegrationRecord) {
	this.Amount = Amount
	return this
}
func (this *IntegrationRecord) GetAmount() (Amount int64) {
	return this.Amount
}
func (this *IntegrationRecord) SetCreateTime(CreateTime xtime.Xtime) (result *IntegrationRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *IntegrationRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func NewIntegrationRecord(id int64, memberId int64, oType IntegrationRecordType.IntegrationRecordType, amount int64, createTime xtime.Xtime) (ret *IntegrationRecord) {
	ret = new(IntegrationRecord)
	ret.Id = id
	ret.MemberId = memberId
	ret.Type = type
	ret.Amount = amount
	ret.CreateTime = createTime
	return
}

type IntegrationRecord struct {
	Id         int64                                       `gorm:"column:id" json:"id"`
	MemberId   int64                                       `gorm:"column:member_id" json:"memberId"`
	Type       IntegrationRecordType.IntegrationRecordType `gorm:"column:type" json:"type"`
	Amount     int64                                       `gorm:"column:amount" json:"amount"`
	CreateTime xtime.Xtime                                 `gorm:"column:create_time" json:"createTime"`
}

