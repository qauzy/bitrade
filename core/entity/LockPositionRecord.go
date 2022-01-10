package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
	"time"
)

func (this *LockPositionRecord) SetId(id int64) (result *LockPositionRecord) {
	this.Id = id
	return this
}
func (this *LockPositionRecord) GetId() (id int64) {
	return this.Id
}
func (this *LockPositionRecord) SetMemberId(memberId int64) (result *LockPositionRecord) {
	this.MemberId = memberId
	return this
}
func (this *LockPositionRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *LockPositionRecord) SetMemberName(memberName string) (result *LockPositionRecord) {
	this.MemberName = memberName
	return this
}
func (this *LockPositionRecord) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *LockPositionRecord) SetCoin(coin *Coin) (result *LockPositionRecord) {
	this.Coin = coin
	return this
}
func (this *LockPositionRecord) GetCoin() (coin *Coin) {
	return this.Coin
}
func (this *LockPositionRecord) SetCreateTime(createTime time.Time) (result *LockPositionRecord) {
	this.CreateTime = createTime
	return this
}
func (this *LockPositionRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *LockPositionRecord) SetStatus(status CommonStatus.CommonStatus) (result *LockPositionRecord) {
	this.Status = status
	return this
}
func (this *LockPositionRecord) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *LockPositionRecord) SetUnlockTime(unlockTime time.Time) (result *LockPositionRecord) {
	this.UnlockTime = unlockTime
	return this
}
func (this *LockPositionRecord) GetUnlockTime() (unlockTime time.Time) {
	return this.UnlockTime
}
func (this *LockPositionRecord) SetReason(reason string) (result *LockPositionRecord) {
	this.Reason = reason
	return this
}
func (this *LockPositionRecord) GetReason() (reason string) {
	return this.Reason
}
func (this *LockPositionRecord) SetAmount(amount math.BigDecimal) (result *LockPositionRecord) {
	this.Amount = amount
	return this
}
func (this *LockPositionRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *LockPositionRecord) SetWalletId(walletId int64) (result *LockPositionRecord) {
	this.WalletId = walletId
	return this
}
func (this *LockPositionRecord) GetWalletId() (walletId int64) {
	return this.WalletId
}

type LockPositionRecord struct {
	Id         int64
	MemberId   int64
	MemberName string
	Coin       *Coin
	CreateTime time.Time
	Status     CommonStatus.CommonStatus
	UnlockTime time.Time
	Reason     string
	Amount     math.BigDecimal
	WalletId   int64
}
