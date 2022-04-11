package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/chocolate/xtime"
)

func (this *MemberAddress) SetId(Id int64) (result *MemberAddress) {
	this.Id = Id
	return this
}
func (this *MemberAddress) GetId() (Id int64) {
	return this.Id
}
func (this *MemberAddress) SetCreateTime(CreateTime xtime.Xtime) (result *MemberAddress) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberAddress) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *MemberAddress) SetDeleteTime(DeleteTime xtime.Xtime) (result *MemberAddress) {
	this.DeleteTime = DeleteTime
	return this
}
func (this *MemberAddress) GetDeleteTime() (DeleteTime xtime.Xtime) {
	return this.DeleteTime
}
func (this *MemberAddress) SetCoin(Coin *Coin) (result *MemberAddress) {
	this.Coin = Coin
	return this
}
func (this *MemberAddress) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *MemberAddress) SetAddress(Address string) (result *MemberAddress) {
	this.Address = Address
	return this
}
func (this *MemberAddress) GetAddress() (Address string) {
	return this.Address
}
func (this *MemberAddress) SetStatus(Status CommonStatus.CommonStatus) (result *MemberAddress) {
	this.Status = Status
	return this
}
func (this *MemberAddress) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *MemberAddress) SetMemberId(MemberId int64) (result *MemberAddress) {
	this.MemberId = MemberId
	return this
}
func (this *MemberAddress) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *MemberAddress) SetRemark(Remark string) (result *MemberAddress) {
	this.Remark = Remark
	return this
}
func (this *MemberAddress) GetRemark() (Remark string) {
	return this.Remark
}
func NewMemberAddress(id int64, createTime xtime.Xtime, deleteTime xtime.Xtime, coin *Coin, address string, status CommonStatus.CommonStatus, memberId int64, remark string) (ret *MemberAddress) {
	ret = new(MemberAddress)
	ret.Id = id
	ret.CreateTime = createTime
	ret.DeleteTime = deleteTime
	ret.Coin = coin
	ret.Address = address
	ret.Status = status
	ret.MemberId = memberId
	ret.Remark = remark
	return
}

type MemberAddress struct {
	Id         int64                     `gorm:"column:id" json:"id"`
	CreateTime xtime.Xtime               `gorm:"column:create_time" json:"createTime"`
	DeleteTime xtime.Xtime               `gorm:"column:delete_time" json:"deleteTime"`
	Coin       *Coin                     `gorm:"column:coin" json:"coin"`
	Address    string                    `gorm:"column:address" json:"address"`
	Status     CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
	MemberId   int64                     `gorm:"column:member_id" json:"memberId"`
	Remark     string                    `gorm:"column:remark" json:"remark"`
}
