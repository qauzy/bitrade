package transform

import (
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/xtime"
)

func (this *MemberAdvertiseInfo) SetBuy(Buy arraylist.List[ScanAdvertise]) (result *MemberAdvertiseInfo) {
	this.Buy = Buy
	return this
}
func (this *MemberAdvertiseInfo) GetBuy() (Buy arraylist.List[ScanAdvertise]) {
	return this.Buy
}
func (this *MemberAdvertiseInfo) SetSell(Sell arraylist.List[ScanAdvertise]) (result *MemberAdvertiseInfo) {
	this.Sell = Sell
	return this
}
func (this *MemberAdvertiseInfo) GetSell() (Sell arraylist.List[ScanAdvertise]) {
	return this.Sell
}
func (this *MemberAdvertiseInfo) SetUsername(Username string) (result *MemberAdvertiseInfo) {
	this.Username = Username
	return this
}
func (this *MemberAdvertiseInfo) GetUsername() (Username string) {
	return this.Username
}
func (this *MemberAdvertiseInfo) SetAvatar(Avatar string) (result *MemberAdvertiseInfo) {
	this.Avatar = Avatar
	return this
}
func (this *MemberAdvertiseInfo) GetAvatar() (Avatar string) {
	return this.Avatar
}
func (this *MemberAdvertiseInfo) SetRealVerified(RealVerified BooleanEnum.BooleanEnum) (result *MemberAdvertiseInfo) {
	this.RealVerified = RealVerified
	return this
}
func (this *MemberAdvertiseInfo) GetRealVerified() (RealVerified BooleanEnum.BooleanEnum) {
	return this.RealVerified
}
func (this *MemberAdvertiseInfo) SetEmailVerified(EmailVerified BooleanEnum.BooleanEnum) (result *MemberAdvertiseInfo) {
	this.EmailVerified = EmailVerified
	return this
}
func (this *MemberAdvertiseInfo) GetEmailVerified() (EmailVerified BooleanEnum.BooleanEnum) {
	return this.EmailVerified
}
func (this *MemberAdvertiseInfo) SetPhoneVerified(PhoneVerified BooleanEnum.BooleanEnum) (result *MemberAdvertiseInfo) {
	this.PhoneVerified = PhoneVerified
	return this
}
func (this *MemberAdvertiseInfo) GetPhoneVerified() (PhoneVerified BooleanEnum.BooleanEnum) {
	return this.PhoneVerified
}
func (this *MemberAdvertiseInfo) SetTransactions(Transactions int) (result *MemberAdvertiseInfo) {
	this.Transactions = Transactions
	return this
}
func (this *MemberAdvertiseInfo) GetTransactions() (Transactions int) {
	return this.Transactions
}
func (this *MemberAdvertiseInfo) SetCreateTime(CreateTime xtime.Xtime) (result *MemberAdvertiseInfo) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberAdvertiseInfo) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func NewMemberAdvertiseInfo(buy arraylist.List[ScanAdvertise], sell arraylist.List[ScanAdvertise], username string, avatar string, realVerified BooleanEnum.BooleanEnum, emailVerified BooleanEnum.BooleanEnum, phoneVerified BooleanEnum.BooleanEnum, transactions int, createTime xtime.Xtime) (ret *MemberAdvertiseInfo) {
	ret = new(MemberAdvertiseInfo)
	ret.Buy = buy
	ret.Sell = sell
	ret.Username = username
	ret.Avatar = avatar
	ret.RealVerified = realVerified
	ret.EmailVerified = emailVerified
	ret.PhoneVerified = phoneVerified
	ret.Transactions = transactions
	ret.CreateTime = createTime
	return
}

type MemberAdvertiseInfo struct {
	Buy           arraylist.List[ScanAdvertise] `gorm:"column:buy" json:"buy"`
	Sell          arraylist.List[ScanAdvertise] `gorm:"column:sell" json:"sell"`
	Username      string                        `gorm:"column:username" json:"username"`
	Avatar        string                        `gorm:"column:avatar" json:"avatar"`
	RealVerified  BooleanEnum.BooleanEnum       `gorm:"column:real_verified" json:"realVerified"`
	EmailVerified BooleanEnum.BooleanEnum       `gorm:"column:email_verified" json:"emailVerified"`
	PhoneVerified BooleanEnum.BooleanEnum       `gorm:"column:phone_verified" json:"phoneVerified"`
	Transactions  int                           `gorm:"column:transactions" json:"transactions"`
	CreateTime    xtime.Xtime                   `gorm:"column:create_time" json:"createTime"`
}
