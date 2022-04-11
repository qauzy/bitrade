package entity

import (
	"bitrade/core/constant/TransactionType"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *RobotTransaction) SetId(Id int64) (result *RobotTransaction) {
	this.Id = Id
	return this
}
func (this *RobotTransaction) GetId() (Id int64) {
	return this.Id
}
func (this *RobotTransaction) SetMemberId(MemberId int64) (result *RobotTransaction) {
	this.MemberId = MemberId
	return this
}
func (this *RobotTransaction) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *RobotTransaction) SetAmount(Amount math.BigDecimal) (result *RobotTransaction) {
	this.Amount = Amount
	return this
}
func (this *RobotTransaction) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *RobotTransaction) SetCreateTime(CreateTime xtime.Xtime) (result *RobotTransaction) {
	this.CreateTime = CreateTime
	return this
}
func (this *RobotTransaction) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *RobotTransaction) SetType(Type TransactionType.TransactionType) (result *RobotTransaction) {
	this.Type = Type
	return this
}
func (this *RobotTransaction) GetType() (Type TransactionType.TransactionType) {
	return this.Type
}
func (this *RobotTransaction) SetSymbol(Symbol string) (result *RobotTransaction) {
	this.Symbol = Symbol
	return this
}
func (this *RobotTransaction) GetSymbol() (Symbol string) {
	return this.Symbol
}
func (this *RobotTransaction) SetFee(Fee math.BigDecimal) (result *RobotTransaction) {
	this.Fee = Fee
	return this
}
func (this *RobotTransaction) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func NewRobotTransaction(id int64, memberId int64, amount math.BigDecimal, createTime xtime.Xtime, oType TransactionType.TransactionType, symbol string, fee math.BigDecimal) (ret *RobotTransaction) {
	ret = new(RobotTransaction)
	ret.Id = id
	ret.MemberId = memberId
	ret.Amount = amount
	ret.CreateTime = createTime
	ret.Type = oType
	ret.Symbol = symbol
	ret.Fee = fee
	return
}

type RobotTransaction struct {
	Id         int64                           `gorm:"column:id" json:"id"`
	MemberId   int64                           `gorm:"column:member_id" json:"memberId"`
	Amount     math.BigDecimal                 `gorm:"column:amount" json:"amount"`
	CreateTime xtime.Xtime                     `gorm:"column:create_time" json:"createTime"`
	Type       TransactionType.TransactionType `gorm:"column:type" json:"type"`
	Symbol     string                          `gorm:"column:symbol" json:"symbol"`
	Fee        math.BigDecimal                 `gorm:"column:fee" json:"fee"`
}
