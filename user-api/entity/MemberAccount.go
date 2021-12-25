package entity

func (this *MemberAccount) SetRealName(realName string) (result *MemberAccount) {
	this.RealName = realName
	return this
}
func (this *MemberAccount) GetRealName() (realName string) {
	return this.RealName
}
func (this *MemberAccount) SetBankVerified(bankVerified constant.BooleanEnum) (result *MemberAccount) {
	this.BankVerified = bankVerified
	return this
}
func (this *MemberAccount) GetBankVerified() (bankVerified constant.BooleanEnum) {
	return this.BankVerified
}
func (this *MemberAccount) SetAliVerified(aliVerified constant.BooleanEnum) (result *MemberAccount) {
	this.AliVerified = aliVerified
	return this
}
func (this *MemberAccount) GetAliVerified() (aliVerified constant.BooleanEnum) {
	return this.AliVerified
}
func (this *MemberAccount) SetWechatVerified(wechatVerified constant.BooleanEnum) (result *MemberAccount) {
	this.WechatVerified = wechatVerified
	return this
}
func (this *MemberAccount) GetWechatVerified() (wechatVerified constant.BooleanEnum) {
	return this.WechatVerified
}
func (this *MemberAccount) SetBankInfo(bankInfo BankInfo) (result *MemberAccount) {
	this.BankInfo = bankInfo
	return this
}
func (this *MemberAccount) GetBankInfo() (bankInfo BankInfo) {
	return this.BankInfo
}
func (this *MemberAccount) SetAlipay(alipay Alipay) (result *MemberAccount) {
	this.Alipay = alipay
	return this
}
func (this *MemberAccount) GetAlipay() (alipay Alipay) {
	return this.Alipay
}
func (this *MemberAccount) SetWechatPay(wechatPay WechatPay) (result *MemberAccount) {
	this.WechatPay = wechatPay
	return this
}
func (this *MemberAccount) GetWechatPay() (wechatPay WechatPay) {
	return this.WechatPay
}
func (this *MemberAccount) SetMemberLevel(memberLevel int64) (result *MemberAccount) {
	this.MemberLevel = memberLevel
	return this
}
func (this *MemberAccount) GetMemberLevel() (memberLevel int64) {
	return this.MemberLevel
}

type MemberAccount struct {
	RealName       string
	BankVerified   constant.BooleanEnum
	AliVerified    constant.BooleanEnum
	WechatVerified constant.BooleanEnum
	BankInfo       BankInfo
	Alipay         Alipay
	WechatPay      WechatPay
	MemberLevel    int64
}
