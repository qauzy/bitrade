package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *LockPositionRecord) SetId(Id int64) (result *LockPositionRecord) {
	this.Id = Id
	return this
}
func (this *LockPositionRecord) GetId() (Id int64) {
	return this.Id
}
func (this *LockPositionRecord) SetMemberId(MemberId int64) (result *LockPositionRecord) {
	this.MemberId = MemberId
	return this
}
func (this *LockPositionRecord) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *LockPositionRecord) SetMemberName(MemberName string) (result *LockPositionRecord) {
	this.MemberName = MemberName
	return this
}
func (this *LockPositionRecord) GetMemberName() (MemberName string) {
	return this.MemberName
}
func (this *LockPositionRecord) SetCoin(Coin *Coin) (result *LockPositionRecord) {
	this.Coin = Coin
	return this
}
func (this *LockPositionRecord) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *LockPositionRecord) SetCreateTime(CreateTime xtime.Xtime) (result *LockPositionRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *LockPositionRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *LockPositionRecord) SetStatus(Status CommonStatus.CommonStatus) (result *LockPositionRecord) {
	this.Status = Status
	return this
}
func (this *LockPositionRecord) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *LockPositionRecord) SetUnlockTime(UnlockTime xtime.Xtime) (result *LockPositionRecord) {
	this.UnlockTime = UnlockTime
	return this
}
func (this *LockPositionRecord) GetUnlockTime() (UnlockTime xtime.Xtime) {
	return this.UnlockTime
}
func (this *LockPositionRecord) SetReason(Reason string) (result *LockPositionRecord) {
	this.Reason = Reason
	return this
}
func (this *LockPositionRecord) GetReason() (Reason string) {
	return this.Reason
}
func (this *LockPositionRecord) SetAmount(Amount math.BigDecimal) (result *LockPositionRecord) {
	this.Amount = Amount
	return this
}
func (this *LockPositionRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *LockPositionRecord) SetWalletId(WalletId int64) (result *LockPositionRecord) {
	this.WalletId = WalletId
	return this
}
func (this *LockPositionRecord) GetWalletId() (WalletId int64) {
	return this.WalletId
}
func NewLockPositionRecord(id int64, memberId int64, memberName string, coin *Coin, createTime xtime.Xtime, status CommonStatus.CommonStatus, unlockTime xtime.Xtime, reason string, amount math.BigDecimal, walletId int64) (ret *LockPositionRecord) {
	ret = new(LockPositionRecord)
	ret.Id = id
	ret.MemberId = memberId
	ret.MemberName = memberName
	ret.Coin = coin
	ret.CreateTime = createTime
	ret.Status = status
	ret.UnlockTime = unlockTime
	ret.Reason = reason
	ret.Amount = amount
	ret.WalletId = walletId
	return
}

type LockPositionRecord struct {
	Id         int64                     `gorm:"column:id" json:"id"`
	MemberId   int64                     `gorm:"column:member_id" json:"memberId"`
	MemberName string                    `gorm:"column:member_name" json:"memberName"`
	Coin       *Coin                     `gorm:"column:coin" json:"coin"`
	CreateTime xtime.Xtime               `gorm:"column:create_time" json:"createTime"`
	Status     CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
	UnlockTime xtime.Xtime               `gorm:"column:unlock_time" json:"unlockTime"`
	Reason     string                    `gorm:"column:reason" json:"reason"`
	Amount     math.BigDecimal           `gorm:"column:amount" json:"amount"`
	WalletId   int64                     `gorm:"column:wallet_id" json:"walletId"`
}
