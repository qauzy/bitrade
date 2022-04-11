package entity

import (
	"bitrade/core/constant/TransactionType"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *MemberTransaction) SetId(Id int64) (result *MemberTransaction) {
	this.Id = Id
	return this
}
func (this *MemberTransaction) GetId() (Id int64) {
	return this.Id
}
func (this *MemberTransaction) SetMemberId(MemberId int64) (result *MemberTransaction) {
	this.MemberId = MemberId
	return this
}
func (this *MemberTransaction) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *MemberTransaction) SetAmount(Amount math.BigDecimal) (result *MemberTransaction) {
	this.Amount = Amount
	return this
}
func (this *MemberTransaction) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberTransaction) SetCreateTime(CreateTime xtime.Xtime) (result *MemberTransaction) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberTransaction) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *MemberTransaction) SetType(Type TransactionType.TransactionType) (result *MemberTransaction) {
	this.Type = Type
	return this
}
func (this *MemberTransaction) GetType() (Type TransactionType.TransactionType) {
	return this.Type
}
func (this *MemberTransaction) SetSymbol(Symbol string) (result *MemberTransaction) {
	this.Symbol = Symbol
	return this
}
func (this *MemberTransaction) GetSymbol() (Symbol string) {
	return this.Symbol
}
func (this *MemberTransaction) SetAddress(Address string) (result *MemberTransaction) {
	this.Address = Address
	return this
}
func (this *MemberTransaction) GetAddress() (Address string) {
	return this.Address
}
func (this *MemberTransaction) SetFee(Fee math.BigDecimal) (result *MemberTransaction) {
	this.Fee = Fee
	return this
}
func (this *MemberTransaction) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func (this *MemberTransaction) SetFlag(Flag int) (result *MemberTransaction) {
	this.Flag = Flag
	return this
}
func (this *MemberTransaction) GetFlag() (Flag int) {
	return this.Flag
}
func (this *MemberTransaction) SetAirdropId(AirdropId int64) (result *MemberTransaction) {
	this.AirdropId = AirdropId
	return this
}
func (this *MemberTransaction) GetAirdropId() (AirdropId int64) {
	return this.AirdropId
}
func NewMemberTransaction(id int64, memberId int64, amount math.BigDecimal, createTime xtime.Xtime, oType TransactionType.TransactionType, symbol string, address string, fee math.BigDecimal, flag int, airdropId int64) (ret *MemberTransaction) {
	ret = new(MemberTransaction)
	ret.Id = id
	ret.MemberId = memberId
	ret.Amount = amount
	ret.CreateTime = createTime
	ret.Type = oType
	ret.Symbol = symbol
	ret.Address = address
	ret.Fee = fee
	ret.Flag = flag
	ret.AirdropId = airdropId
	return
}

type MemberTransaction struct {
	Id         int64                           `gorm:"column:id" json:"id"`
	MemberId   int64                           `gorm:"column:member_id" json:"memberId"`
	Amount     math.BigDecimal                 `gorm:"column:amount" json:"amount"`
	CreateTime xtime.Xtime                     `gorm:"column:create_time" json:"createTime"`
	Type       TransactionType.TransactionType `gorm:"column:type" json:"type"`
	Symbol     string                          `gorm:"column:symbol" json:"symbol"`
	Address    string                          `gorm:"column:address" json:"address"`
	Fee        math.BigDecimal                 `gorm:"column:fee" json:"fee"`
	Flag       int                             `gorm:"column:flag" json:"flag"`
	AirdropId  int64                           `gorm:"column:airdrop_id" json:"airdropId"`
}
