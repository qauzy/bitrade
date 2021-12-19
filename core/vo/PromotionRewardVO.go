package vo

import "github.com/qauzy/math"

func (this *PromotionRewardVO) SetUsername(username string) (result *PromotionRewardVO) {
	this.Username = username
	return this
}
func (this *PromotionRewardVO) GetUsername() (username string) {
	return this.Username
}
func (this *PromotionRewardVO) SetRealName(realName string) (result *PromotionRewardVO) {
	this.RealName = realName
	return this
}
func (this *PromotionRewardVO) GetRealName() (realName string) {
	return this.RealName
}
func (this *PromotionRewardVO) SetMobilePhone(mobilePhone string) (result *PromotionRewardVO) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *PromotionRewardVO) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *PromotionRewardVO) SetEmail(email string) (result *PromotionRewardVO) {
	this.Email = email
	return this
}
func (this *PromotionRewardVO) GetEmail() (email string) {
	return this.Email
}
func (this *PromotionRewardVO) SetPromotionCode(promotionCode string) (result *PromotionRewardVO) {
	this.PromotionCode = promotionCode
	return this
}
func (this *PromotionRewardVO) GetPromotionCode() (promotionCode string) {
	return this.PromotionCode
}
func (this *PromotionRewardVO) SetPromotionNum(promotionNum int) (result *PromotionRewardVO) {
	this.PromotionNum = promotionNum
	return this
}
func (this *PromotionRewardVO) GetPromotionNum() (promotionNum int) {
	return this.PromotionNum
}
func (this *PromotionRewardVO) SetPromotionReward(promotionReward math.BigDecimal) (result *PromotionRewardVO) {
	this.PromotionReward = promotionReward
	return this
}
func (this *PromotionRewardVO) GetPromotionReward() (promotionReward math.BigDecimal) {
	return this.PromotionReward
}

type PromotionRewardVO struct {
	Username        string
	RealName        string
	MobilePhone     string
	Email           string
	PromotionCode   string
	PromotionNum    int
	PromotionReward math.BigDecimal
}
