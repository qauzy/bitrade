package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *MemberDeposit) SetId(Id int64) (result *MemberDeposit) {
	this.Id = Id
	return this
}
func (this *MemberDeposit) GetId() (Id int64) {
	return this.Id
}
func (this *MemberDeposit) SetMemberId(MemberId int64) (result *MemberDeposit) {
	this.MemberId = MemberId
	return this
}
func (this *MemberDeposit) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *MemberDeposit) SetAmount(Amount math.BigDecimal) (result *MemberDeposit) {
	this.Amount = Amount
	return this
}
func (this *MemberDeposit) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberDeposit) SetUnit(Unit string) (result *MemberDeposit) {
	this.Unit = Unit
	return this
}
func (this *MemberDeposit) GetUnit() (Unit string) {
	return this.Unit
}
func (this *MemberDeposit) SetCreateTime(CreateTime xtime.Xtime) (result *MemberDeposit) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberDeposit) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *MemberDeposit) SetTxid(Txid string) (result *MemberDeposit) {
	this.Txid = Txid
	return this
}
func (this *MemberDeposit) GetTxid() (Txid string) {
	return this.Txid
}
func (this *MemberDeposit) SetAddress(Address string) (result *MemberDeposit) {
	this.Address = Address
	return this
}
func (this *MemberDeposit) GetAddress() (Address string) {
	return this.Address
}
func (this *MemberDeposit) SetUsername(Username string) (result *MemberDeposit) {
	this.Username = Username
	return this
}
func (this *MemberDeposit) GetUsername() (Username string) {
	return this.Username
}
func NewMemberDeposit(id int64, memberId int64, amount math.BigDecimal, unit string, createTime xtime.Xtime, txid string, address string, username string) (ret *MemberDeposit) {
	ret = new(MemberDeposit)
	ret.Id = id
	ret.MemberId = memberId
	ret.Amount = amount
	ret.Unit = unit
	ret.CreateTime = createTime
	ret.Txid = txid
	ret.Address = address
	ret.Username = username
	return
}

type MemberDeposit struct {
	Id         int64           `gorm:"column:id" json:"id"`
	MemberId   int64           `gorm:"column:member_id" json:"memberId"`
	Amount     math.BigDecimal `gorm:"column:amount" json:"amount"`
	Unit       string          `gorm:"column:unit" json:"unit"`
	CreateTime xtime.Xtime     `gorm:"column:create_time" json:"createTime"`
	Txid       string          `gorm:"column:txid" json:"txid"`
	Address    string          `gorm:"column:address" json:"address"`
	Username   string          `gorm:"column:username" json:"username"`
}
