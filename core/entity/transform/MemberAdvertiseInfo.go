package transform

import (
	"bitrade/core/constant"
	"time"
)

func (this *MemberAdvertiseInfo) SetBuy(buy []ScanAdvertise) (result *MemberAdvertiseInfo) {
	this.Buy = buy
	return this
}
func (this *MemberAdvertiseInfo) GetBuy() (buy []ScanAdvertise) {
	return this.Buy
}
func (this *MemberAdvertiseInfo) SetSell(sell []ScanAdvertise) (result *MemberAdvertiseInfo) {
	this.Sell = sell
	return this
}
func (this *MemberAdvertiseInfo) GetSell() (sell []ScanAdvertise) {
	return this.Sell
}
func (this *MemberAdvertiseInfo) SetUsername(username string) (result *MemberAdvertiseInfo) {
	this.Username = username
	return this
}
func (this *MemberAdvertiseInfo) GetUsername() (username string) {
	return this.Username
}
func (this *MemberAdvertiseInfo) SetAvatar(avatar string) (result *MemberAdvertiseInfo) {
	this.Avatar = avatar
	return this
}
func (this *MemberAdvertiseInfo) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *MemberAdvertiseInfo) SetRealVerified(realVerified constant.BooleanEnum) (result *MemberAdvertiseInfo) {
	this.RealVerified = realVerified
	return this
}
func (this *MemberAdvertiseInfo) GetRealVerified() (realVerified constant.BooleanEnum) {
	return this.RealVerified
}
func (this *MemberAdvertiseInfo) SetEmailVerified(emailVerified constant.BooleanEnum) (result *MemberAdvertiseInfo) {
	this.EmailVerified = emailVerified
	return this
}
func (this *MemberAdvertiseInfo) GetEmailVerified() (emailVerified constant.BooleanEnum) {
	return this.EmailVerified
}
func (this *MemberAdvertiseInfo) SetPhoneVerified(phoneVerified constant.BooleanEnum) (result *MemberAdvertiseInfo) {
	this.PhoneVerified = phoneVerified
	return this
}
func (this *MemberAdvertiseInfo) GetPhoneVerified() (phoneVerified constant.BooleanEnum) {
	return this.PhoneVerified
}
func (this *MemberAdvertiseInfo) SetTransactions(transactions int) (result *MemberAdvertiseInfo) {
	this.Transactions = transactions
	return this
}
func (this *MemberAdvertiseInfo) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *MemberAdvertiseInfo) SetCreateTime(createTime time.Time) (result *MemberAdvertiseInfo) {
	this.CreateTime = createTime
	return this
}
func (this *MemberAdvertiseInfo) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type MemberAdvertiseInfo struct {
	Buy           []ScanAdvertise
	Sell          []ScanAdvertise
	Username      string
	Avatar        string
	RealVerified  constant.BooleanEnum
	EmailVerified constant.BooleanEnum
	PhoneVerified constant.BooleanEnum
	Transactions  int
	CreateTime    time.Time
}