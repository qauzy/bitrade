package entity

import (
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/entity"
)

func (this *LoginInfo) SetUsername(Username string) (result *LoginInfo) {
	this.Username = Username
	return this
}
func (this *LoginInfo) GetUsername() (Username string) {
	return this.Username
}
func (this *LoginInfo) SetLocation(Location *entity.Location) (result *LoginInfo) {
	this.Location = Location
	return this
}
func (this *LoginInfo) GetLocation() (Location *entity.Location) {
	return this.Location
}
func (this *LoginInfo) SetMemberLevel(MemberLevel *MemberLevelEnum.MemberLevelEnum) (result *LoginInfo) {
	this.MemberLevel = MemberLevel
	return this
}
func (this *LoginInfo) GetMemberLevel() (MemberLevel *MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *LoginInfo) SetToken(Token string) (result *LoginInfo) {
	this.Token = Token
	return this
}
func (this *LoginInfo) GetToken() (Token string) {
	return this.Token
}
func (this *LoginInfo) SetRealName(RealName string) (result *LoginInfo) {
	this.RealName = RealName
	return this
}
func (this *LoginInfo) GetRealName() (RealName string) {
	return this.RealName
}
func (this *LoginInfo) SetCountry(Country *entity.Country) (result *LoginInfo) {
	this.Country = Country
	return this
}
func (this *LoginInfo) GetCountry() (Country *entity.Country) {
	return this.Country
}
func (this *LoginInfo) SetAvatar(Avatar string) (result *LoginInfo) {
	this.Avatar = Avatar
	return this
}
func (this *LoginInfo) GetAvatar() (Avatar string) {
	return this.Avatar
}
func (this *LoginInfo) SetPromotionCode(PromotionCode string) (result *LoginInfo) {
	this.PromotionCode = PromotionCode
	return this
}
func (this *LoginInfo) GetPromotionCode() (PromotionCode string) {
	return this.PromotionCode
}
func (this *LoginInfo) SetId(Id int64) (result *LoginInfo) {
	this.Id = Id
	return this
}
func (this *LoginInfo) GetId() (Id int64) {
	return this.Id
}
func (this *LoginInfo) SetGoogleState(GoogleState int) (result *LoginInfo) {
	this.GoogleState = GoogleState
	return this
}
func (this *LoginInfo) GetGoogleState() (GoogleState int) {
	return this.GoogleState
}
func (this *LoginInfo) SetKycStatus(KycStatus int) (result *LoginInfo) {
	this.KycStatus = KycStatus
	return this
}
func (this *LoginInfo) GetKycStatus() (KycStatus int) {
	return this.KycStatus
}
func (this *LoginInfo) SetPromotionPrefix(PromotionPrefix string) (result *LoginInfo) {
	this.PromotionPrefix = PromotionPrefix
	return this
}
func (this *LoginInfo) GetPromotionPrefix() (PromotionPrefix string) {
	return this.PromotionPrefix
}
func (this *LoginInfo) SetSignInAbility(SignInAbility bool) (result *LoginInfo) {
	this.SignInAbility = SignInAbility
	return this
}
func (this *LoginInfo) GetSignInAbility() (SignInAbility bool) {
	return this.SignInAbility
}
func (this *LoginInfo) SetSignInActivity(SignInActivity bool) (result *LoginInfo) {
	this.SignInActivity = SignInActivity
	return this
}
func (this *LoginInfo) GetSignInActivity() (SignInActivity bool) {
	return this.SignInActivity
}
func (this *LoginInfo) SetMemberGradeId(MemberGradeId int64) (result *LoginInfo) {
	this.MemberGradeId = MemberGradeId
	return this
}
func (this *LoginInfo) GetMemberGradeId() (MemberGradeId int64) {
	return this.MemberGradeId
}
func (this *LoginInfo) SetIntegration(Integration int64) (result *LoginInfo) {
	this.Integration = Integration
	return this
}
func (this *LoginInfo) GetIntegration() (Integration int64) {
	return this.Integration
}
func (this *LoginInfo) SetMobile(Mobile string) (result *LoginInfo) {
	this.Mobile = Mobile
	return this
}
func (this *LoginInfo) GetMobile() (Mobile string) {
	return this.Mobile
}
func GetLoginInfo(member entity.Member, token string, signInActivity bool, prefix string) (result *LoginInfo) {
	return new(LoginInfo).SetLocation(member.GetLocation()).SetMemberLevel(member.GetMemberLevel()).SetUsername(member.GetUsername()).SetToken(token).SetRealName(member.GetRealName()).SetCountry(member.GetCountry()).SetAvatar(member.GetAvatar()).SetPromotionCode(member.GetPromotionCode()).SetId(member.GetId()).SetPromotionPrefix(prefix).SetSignInAbility(member.GetSignInAbility()).SetSignInActivity(signInActivity).SetMemberGradeId(member.GetMemberGradeId()).SetGoogleState(member.GetGoogleState()).SetIntegration(member.GetIntegration()).SetKycStatus(member.GetKycStatus()).SetMobile(member.GetMobilePhone())
}

type LoginInfo struct {
	Username        string                           `gorm:"column:username" json:"username"`
	Location        *entity.Location                 `gorm:"column:location" json:"location"`
	MemberLevel     *MemberLevelEnum.MemberLevelEnum `gorm:"column:member_level" json:"memberLevel"`
	Token           string                           `gorm:"column:token" json:"token"`
	RealName        string                           `gorm:"column:real_name" json:"realName"`
	Country         *entity.Country                  `gorm:"column:country" json:"country"`
	Avatar          string                           `gorm:"column:avatar" json:"avatar"`
	PromotionCode   string                           `gorm:"column:promotion_code" json:"promotionCode"`
	Id              int64                            `gorm:"column:id" json:"id"`
	GoogleState     int                              `gorm:"column:google_state" json:"googleState"`
	KycStatus       int                              `gorm:"column:kyc_status" json:"kycStatus"`
	PromotionPrefix string                           `gorm:"column:promotion_prefix" json:"promotionPrefix"`
	SignInAbility   bool                             `gorm:"column:sign_in_ability" json:"signInAbility"`
	SignInActivity  bool                             `gorm:"column:sign_in_activity" json:"signInActivity"`
	MemberGradeId   int64                            `gorm:"column:member_grade_id" json:"memberGradeId"`
	Integration     int64                            `gorm:"column:integration" json:"integration"`
	Mobile          string                           `gorm:"column:mobile" json:"mobile"`
}
