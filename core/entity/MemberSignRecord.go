package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *MemberSignRecord) SetId(Id int64) (result *MemberSignRecord) {
	this.Id = Id
	return this
}
func (this *MemberSignRecord) GetId() (Id int64) {
	return this.Id
}
func (this *MemberSignRecord) SetMember(Member *Member) (result *MemberSignRecord) {
	this.Member = Member
	return this
}
func (this *MemberSignRecord) GetMember() (Member *Member) {
	return this.Member
}
func (this *MemberSignRecord) SetSign(Sign *Sign) (result *MemberSignRecord) {
	this.Sign = Sign
	return this
}
func (this *MemberSignRecord) GetSign() (Sign *Sign) {
	return this.Sign
}
func (this *MemberSignRecord) SetCoin(Coin *Coin) (result *MemberSignRecord) {
	this.Coin = Coin
	return this
}
func (this *MemberSignRecord) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *MemberSignRecord) SetAmount(Amount math.BigDecimal) (result *MemberSignRecord) {
	this.Amount = Amount
	return this
}
func (this *MemberSignRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberSignRecord) SetCreationTime(CreationTime xtime.Xtime) (result *MemberSignRecord) {
	this.CreationTime = CreationTime
	return this
}
func (this *MemberSignRecord) GetCreationTime() (CreationTime xtime.Xtime) {
	return this.CreationTime
}
func NewMemberSignRecord() (this *MemberSignRecord) {
	this = new(MemberSignRecord)
	return
}
func NewMemberSignRecordV2(member *Member, sign *Sign) (this *MemberSignRecord) {
	this = new(MemberSignRecord)
	this.Member = this.Member
	this.Sign = this.Sign
	this.Coin = this.Sign.GetCoin()
	//防止sign信息修改
	this.Amount = this.Sign.GetAmount()
	//防止sign信息修改
	return
}
func NewMemberSignRecordEx(id int64, member *Member, sign *Sign, coin *Coin, amount math.BigDecimal, creationTime xtime.Xtime) (ret *MemberSignRecord) {
	ret = new(MemberSignRecord)
	ret.Id = id
	ret.Member = member
	ret.Sign = sign
	ret.Coin = coin
	ret.Amount = amount
	ret.CreationTime = creationTime
	return
}

type MemberSignRecord struct {
	Id           int64           `gorm:"column:id" json:"id"`
	Member       *Member         `gorm:"column:member" json:"member"`
	Sign         *Sign           `gorm:"column:sign" json:"sign"`
	Coin         *Coin           `gorm:"column:coin" json:"coin"`
	Amount       math.BigDecimal `gorm:"column:amount" json:"amount"`
	CreationTime xtime.Xtime     `gorm:"column:creation_time" json:"creationTime"`
}
