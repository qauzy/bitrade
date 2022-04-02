package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/entity"
)

func (this *MemberAccount) SetRealName(RealName string) (result *MemberAccount) {
	this.RealName = RealName
	return this
}
func (this *MemberAccount) GetRealName() (RealName string) {
	return this.RealName
}
func (this *MemberAccount) SetBankVerified(BankVerified *BooleanEnum.BooleanEnum) (result *MemberAccount) {
	this.BankVerified = BankVerified
	return this
}
func (this *MemberAccount) GetBankVerified() (BankVerified *BooleanEnum.BooleanEnum) {
	return this.BankVerified
}
func (this *MemberAccount) SetAliVerified(AliVerified *BooleanEnum.BooleanEnum) (result *MemberAccount) {
	this.AliVerified = AliVerified
	return this
}
func (this *MemberAccount) GetAliVerified() (AliVerified *BooleanEnum.BooleanEnum) {
	return this.AliVerified
}
func (this *MemberAccount) SetWechatVerified(WechatVerified *BooleanEnum.BooleanEnum) (result *MemberAccount) {
	this.WechatVerified = WechatVerified
	return this
}
func (this *MemberAccount) GetWechatVerified() (WechatVerified *BooleanEnum.BooleanEnum) {
	return this.WechatVerified
}
func (this *MemberAccount) SetBankInfo(BankInfo *entity.BankInfo) (result *MemberAccount) {
	this.BankInfo = BankInfo
	return this
}
func (this *MemberAccount) GetBankInfo() (BankInfo *entity.BankInfo) {
	return this.BankInfo
}
func (this *MemberAccount) SetAlipay(Alipay *entity.Alipay) (result *MemberAccount) {
	this.Alipay = Alipay
	return this
}
func (this *MemberAccount) GetAlipay() (Alipay *entity.Alipay) {
	return this.Alipay
}
func (this *MemberAccount) SetWechatPay(WechatPay *entity.WechatPay) (result *MemberAccount) {
	this.WechatPay = WechatPay
	return this
}
func (this *MemberAccount) GetWechatPay() (WechatPay *entity.WechatPay) {
	return this.WechatPay
}
func (this *MemberAccount) SetMemberLevel(MemberLevel int) (result *MemberAccount) {
	this.MemberLevel = MemberLevel
	return this
}
func (this *MemberAccount) GetMemberLevel() (MemberLevel int) {
	return this.MemberLevel
}

type MemberAccount struct {
	RealName       string                   `gorm:"column:real_name" json:"realName"`
	BankVerified   *BooleanEnum.BooleanEnum `gorm:"column:bank_verified" json:"bankVerified"`
	AliVerified    *BooleanEnum.BooleanEnum `gorm:"column:ali_verified" json:"aliVerified"`
	WechatVerified *BooleanEnum.BooleanEnum `gorm:"column:wechat_verified" json:"wechatVerified"`
	BankInfo       *entity.BankInfo         `gorm:"column:bank_info" json:"bankInfo"`
	Alipay         *entity.Alipay           `gorm:"column:alipay" json:"alipay"`
	WechatPay      *entity.WechatPay        `gorm:"column:wechat_pay" json:"wechatPay"`
	MemberLevel    int                      `gorm:"column:member_level" json:"memberLevel"`
}
