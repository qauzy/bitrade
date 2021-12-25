package entity

func (this *LoginInfo) SetUsername(username string) (result *LoginInfo) {
	this.Username = username
	return this
}
func (this *LoginInfo) GetUsername() (username string) {
	return this.Username
}
func (this *LoginInfo) SetLocation(location Location) (result *LoginInfo) {
	this.Location = location
	return this
}
func (this *LoginInfo) GetLocation() (location Location) {
	return this.Location
}
func (this *LoginInfo) SetMemberLevel(memberLevel constant.MemberLevelEnum) (result *LoginInfo) {
	this.MemberLevel = memberLevel
	return this
}
func (this *LoginInfo) GetMemberLevel() (memberLevel constant.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *LoginInfo) SetToken(token string) (result *LoginInfo) {
	this.Token = token
	return this
}
func (this *LoginInfo) GetToken() (token string) {
	return this.Token
}
func (this *LoginInfo) SetRealName(realName string) (result *LoginInfo) {
	this.RealName = realName
	return this
}
func (this *LoginInfo) GetRealName() (realName string) {
	return this.RealName
}
func (this *LoginInfo) SetCountry(country Country) (result *LoginInfo) {
	this.Country = country
	return this
}
func (this *LoginInfo) GetCountry() (country Country) {
	return this.Country
}
func (this *LoginInfo) SetAvatar(avatar string) (result *LoginInfo) {
	this.Avatar = avatar
	return this
}
func (this *LoginInfo) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *LoginInfo) SetPromotionCode(promotionCode string) (result *LoginInfo) {
	this.PromotionCode = promotionCode
	return this
}
func (this *LoginInfo) GetPromotionCode() (promotionCode string) {
	return this.PromotionCode
}
func (this *LoginInfo) SetId(id int64) (result *LoginInfo) {
	this.Id = id
	return this
}
func (this *LoginInfo) GetId() (id int64) {
	return this.Id
}
func (this *LoginInfo) SetGoogleState(googleState int64) (result *LoginInfo) {
	this.GoogleState = googleState
	return this
}
func (this *LoginInfo) GetGoogleState() (googleState int64) {
	return this.GoogleState
}
func (this *LoginInfo) SetKycStatus(kycStatus int64) (result *LoginInfo) {
	this.KycStatus = kycStatus
	return this
}
func (this *LoginInfo) GetKycStatus() (kycStatus int64) {
	return this.KycStatus
}
func (this *LoginInfo) SetPromotionPrefix(promotionPrefix string) (result *LoginInfo) {
	this.PromotionPrefix = promotionPrefix
	return this
}
func (this *LoginInfo) GetPromotionPrefix() (promotionPrefix string) {
	return this.PromotionPrefix
}
func (this *LoginInfo) SetSignInAbility(signInAbility bool) (result *LoginInfo) {
	this.SignInAbility = signInAbility
	return this
}
func (this *LoginInfo) GetSignInAbility() (signInAbility bool) {
	return this.SignInAbility
}
func (this *LoginInfo) SetSignInActivity(signInActivity bool) (result *LoginInfo) {
	this.SignInActivity = signInActivity
	return this
}
func (this *LoginInfo) GetSignInActivity() (signInActivity bool) {
	return this.SignInActivity
}
func (this *LoginInfo) SetMemberGradeId(memberGradeId int64) (result *LoginInfo) {
	this.MemberGradeId = memberGradeId
	return this
}
func (this *LoginInfo) GetMemberGradeId() (memberGradeId int64) {
	return this.MemberGradeId
}
func (this *LoginInfo) SetIntegration(integration int64) (result *LoginInfo) {
	this.Integration = integration
	return this
}
func (this *LoginInfo) GetIntegration() (integration int64) {
	return this.Integration
}
func (this *LoginInfo) SetMobile(mobile string) (result *LoginInfo) {
	this.Mobile = mobile
	return this
}
func (this *LoginInfo) GetMobile() (mobile string) {
	return this.Mobile
}

/**
 *
 * @param member
 * @param token
 * @param signInActivity
 * @param prefix
 * @return
 */
func GetLoginInfo(member Member, token string, signInActivity bool, prefix string) (result LoginInfo) {
	return new(LoginInfo).SetLocation(member.GetLocation()).SetMemberLevel(member.GetMemberLevel()).SetUsername(member.GetUsername()).SetToken(token).SetRealName(member.GetRealName()).SetCountry(member.GetCountry()).SetAvatar(member.GetAvatar()).SetPromotionCode(member.GetPromotionCode()).SetId(member.GetId()).SetPromotionPrefix(prefix).SetSignInAbility(member.GetSignInAbility()).SetSignInActivity(signInActivity).SetMemberGradeId(member.GetMemberGradeId()).SetGoogleState(member.GetGoogleState()).SetIntegration(member.GetIntegration()).SetKycStatus(member.GetKycStatus()).SetMobile(member.GetMobilePhone())
}

type LoginInfo struct {
	Username        string
	Location        Location
	MemberLevel     constant.MemberLevelEnum
	Token           string
	RealName        string
	Country         Country
	Avatar          string
	PromotionCode   string
	Id              int64
	GoogleState     int64
	KycStatus       int64
	PromotionPrefix string
	SignInAbility   bool
	SignInActivity  bool
	MemberGradeId   int64
	Integration     int64
	Mobile          string
}
