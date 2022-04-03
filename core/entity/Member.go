package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/constant/RealNameStatus"
	"bitrade/core/vo"
	"time"
)

func (this *Member) SetId(id int64) (result *Member) {
	this.Id = id
	return this
}
func (this *Member) GetId() (id int64) {
	return this.Id
}
func (this *Member) SetSalt(salt string) (result *Member) {
	this.Salt = salt
	return this
}
func (this *Member) GetSalt() (salt string) {
	return this.Salt
}
func (this *Member) SetUsername(username string) (result *Member) {
	this.Username = username
	return this
}
func (this *Member) GetUsername() (username string) {
	return this.Username
}
func (this *Member) SetPassword(password string) (result *Member) {
	this.Password = password
	return this
}
func (this *Member) GetPassword() (password string) {
	return this.Password
}
func (this *Member) SetMargin(margin string) (result *Member) {
	this.Margin = margin
	return this
}
func (this *Member) GetMargin() (margin string) {
	return this.Margin
}
func (this *Member) SetGoogleKey(googleKey string) (result *Member) {
	this.GoogleKey = googleKey
	return this
}
func (this *Member) GetGoogleKey() (googleKey string) {
	return this.GoogleKey
}
func (this *Member) SetGoogleState(googleState int) (result *Member) {
	this.GoogleState = googleState
	return this
}
func (this *Member) GetGoogleState() (googleState int) {
	return this.GoogleState
}
func (this *Member) SetGoogleDate(googleDate time.Time) (result *Member) {
	this.GoogleDate = googleDate
	return this
}
func (this *Member) GetGoogleDate() (googleDate time.Time) {
	return this.GoogleDate
}
func (this *Member) SetJyPassword(jyPassword string) (result *Member) {
	this.JyPassword = jyPassword
	return this
}
func (this *Member) GetJyPassword() (jyPassword string) {
	return this.JyPassword
}
func (this *Member) SetRealName(realName string) (result *Member) {
	this.RealName = realName
	return this
}
func (this *Member) GetRealName() (realName string) {
	return this.RealName
}
func (this *Member) SetIdNumber(idNumber string) (result *Member) {
	this.IdNumber = idNumber
	return this
}
func (this *Member) GetIdNumber() (idNumber string) {
	return this.IdNumber
}
func (this *Member) SetEmail(email string) (result *Member) {
	this.Email = email
	return this
}
func (this *Member) GetEmail() (email string) {
	return this.Email
}
func (this *Member) SetMobilePhone(mobilePhone string) (result *Member) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *Member) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *Member) SetLocation(location *Location) (result *Member) {
	this.Location = location
	return this
}
func (this *Member) GetLocation() (location *Location) {
	return this.Location
}
func (this *Member) SetMemberLevel(memberLevel MemberLevelEnum.MemberLevelEnum) (result *Member) {
	this.MemberLevel = memberLevel
	return this
}
func (this *Member) GetMemberLevel() (memberLevel MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *Member) SetStatus(status CommonStatus.CommonStatus) (result *Member) {
	this.Status = status
	return this
}
func (this *Member) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *Member) SetRegistrationTime(registrationTime time.Time) (result *Member) {
	this.RegistrationTime = registrationTime
	return this
}
func (this *Member) GetRegistrationTime() (registrationTime time.Time) {
	return this.RegistrationTime
}
func (this *Member) SetLastLoginTime(lastLoginTime time.Time) (result *Member) {
	this.LastLoginTime = lastLoginTime
	return this
}
func (this *Member) GetLastLoginTime() (lastLoginTime time.Time) {
	return this.LastLoginTime
}
func (this *Member) SetToken(token string) (result *Member) {
	this.Token = token
	return this
}
func (this *Member) GetToken() (token string) {
	return this.Token
}
func (this *Member) SetTransactions(transactions int) (result *Member) {
	this.Transactions = transactions
	return this
}
func (this *Member) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *Member) SetBankInfo(bankInfo *BankInfo) (result *Member) {
	this.BankInfo = bankInfo
	return this
}
func (this *Member) GetBankInfo() (bankInfo *BankInfo) {
	return this.BankInfo
}
func (this *Member) SetAlipay(alipay *Alipay) (result *Member) {
	this.Alipay = alipay
	return this
}
func (this *Member) GetAlipay() (alipay *Alipay) {
	return this.Alipay
}
func (this *Member) SetWechatPay(wechatPay *WechatPay) (result *Member) {
	this.WechatPay = wechatPay
	return this
}
func (this *Member) GetWechatPay() (wechatPay *WechatPay) {
	return this.WechatPay
}
func (this *Member) SetAppealTimes(appealTimes int) (result *Member) {
	this.AppealTimes = appealTimes
	return this
}
func (this *Member) GetAppealTimes() (appealTimes int) {
	return this.AppealTimes
}
func (this *Member) SetAppealSuccessTimes(appealSuccessTimes int) (result *Member) {
	this.AppealSuccessTimes = appealSuccessTimes
	return this
}
func (this *Member) GetAppealSuccessTimes() (appealSuccessTimes int) {
	return this.AppealSuccessTimes
}
func (this *Member) SetInviterId(inviterId int64) (result *Member) {
	this.InviterId = inviterId
	return this
}
func (this *Member) GetInviterId() (inviterId int64) {
	return this.InviterId
}
func (this *Member) SetPromotionCode(promotionCode string) (result *Member) {
	this.PromotionCode = promotionCode
	return this
}
func (this *Member) GetPromotionCode() (promotionCode string) {
	return this.PromotionCode
}
func (this *Member) SetFirstLevel(firstLevel int) (result *Member) {
	this.FirstLevel = firstLevel
	return this
}
func (this *Member) GetFirstLevel() (firstLevel int) {
	return this.FirstLevel
}
func (this *Member) SetSecondLevel(secondLevel int) (result *Member) {
	this.SecondLevel = secondLevel
	return this
}
func (this *Member) GetSecondLevel() (secondLevel int) {
	return this.SecondLevel
}
func (this *Member) SetThirdLevel(thirdLevel int) (result *Member) {
	this.ThirdLevel = thirdLevel
	return this
}
func (this *Member) GetThirdLevel() (thirdLevel int) {
	return this.ThirdLevel
}
func (this *Member) SetCountry(country *Country) (result *Member) {
	this.Country = country
	return this
}
func (this *Member) GetCountry() (country *Country) {
	return this.Country
}
func (this *Member) SetRealNameStatus(realNameStatus RealNameStatus.RealNameStatus) (result *Member) {
	this.RealNameStatus = realNameStatus
	return this
}
func (this *Member) GetRealNameStatus() (realNameStatus RealNameStatus.RealNameStatus) {
	return this.RealNameStatus
}
func (this *Member) SetLoginCount(loginCount int) (result *Member) {
	this.LoginCount = loginCount
	return this
}
func (this *Member) GetLoginCount() (loginCount int) {
	return this.LoginCount
}
func (this *Member) SetCertifiedBusinessStatus(certifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus) (result *Member) {
	this.CertifiedBusinessStatus = certifiedBusinessStatus
	return this
}
func (this *Member) GetCertifiedBusinessStatus() (certifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *Member) SetCertifiedBusinessApplyTime(certifiedBusinessApplyTime time.Time) (result *Member) {
	this.CertifiedBusinessApplyTime = certifiedBusinessApplyTime
	return this
}
func (this *Member) GetCertifiedBusinessApplyTime() (certifiedBusinessApplyTime time.Time) {
	return this.CertifiedBusinessApplyTime
}
func (this *Member) SetApplicationTime(applicationTime time.Time) (result *Member) {
	this.ApplicationTime = applicationTime
	return this
}
func (this *Member) GetApplicationTime() (applicationTime time.Time) {
	return this.ApplicationTime
}
func (this *Member) SetCertifiedBusinessCheckTime(certifiedBusinessCheckTime time.Time) (result *Member) {
	this.CertifiedBusinessCheckTime = certifiedBusinessCheckTime
	return this
}
func (this *Member) GetCertifiedBusinessCheckTime() (certifiedBusinessCheckTime time.Time) {
	return this.CertifiedBusinessCheckTime
}
func (this *Member) SetAvatar(avatar string) (result *Member) {
	this.Avatar = avatar
	return this
}
func (this *Member) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *Member) SetTokenExpireTime(tokenExpireTime time.Time) (result *Member) {
	this.TokenExpireTime = tokenExpireTime
	return this
}
func (this *Member) GetTokenExpireTime() (tokenExpireTime time.Time) {
	return this.TokenExpireTime
}
func (this *Member) SetPublishAdvertise(publishAdvertise BooleanEnum.BooleanEnum) (result *Member) {
	this.PublishAdvertise = publishAdvertise
	return this
}
func (this *Member) GetPublishAdvertise() (publishAdvertise BooleanEnum.BooleanEnum) {
	return this.PublishAdvertise
}
func (this *Member) SetTransactionStatus(transactionStatus BooleanEnum.BooleanEnum) (result *Member) {
	this.TransactionStatus = transactionStatus
	return this
}
func (this *Member) GetTransactionStatus() (transactionStatus BooleanEnum.BooleanEnum) {
	return this.TransactionStatus
}
func (this *Member) SetSignInAbility(signInAbility bool) (result *Member) {
	this.SignInAbility = signInAbility
	return this
}
func (this *Member) GetSignInAbility() (signInAbility bool) {
	return this.SignInAbility
}
func (this *Member) SetTransactionTime(transactionTime time.Time) (result *Member) {
	this.TransactionTime = transactionTime
	return this
}
func (this *Member) GetTransactionTime() (transactionTime time.Time) {
	return this.TransactionTime
}
func (this *Member) SetChannelId(channelId int64) (result *Member) {
	this.ChannelId = channelId
	return this
}
func (this *Member) GetChannelId() (channelId int64) {
	return this.ChannelId
}
func (this *Member) SetIsChannel(isChannel BooleanEnum.BooleanEnum) (result *Member) {
	this.IsChannel = isChannel
	return this
}
func (this *Member) GetIsChannel() (isChannel BooleanEnum.BooleanEnum) {
	return this.IsChannel
}
func (this *Member) SetChannelVO(channelVO *vo.ChannelVO) (result *Member) {
	this.ChannelVO = channelVO
	return this
}
func (this *Member) GetChannelVO() (channelVO *vo.ChannelVO) {
	return this.ChannelVO
}
func (this *Member) SetLoginLock(loginLock BooleanEnum.BooleanEnum) (result *Member) {
	this.LoginLock = loginLock
	return this
}
func (this *Member) GetLoginLock() (loginLock BooleanEnum.BooleanEnum) {
	return this.LoginLock
}
func (this *Member) SetIntegration(integration int64) (result *Member) {
	this.Integration = integration
	return this
}
func (this *Member) GetIntegration() (integration int64) {
	return this.Integration
}
func (this *Member) SetMemberGradeId(memberGradeId int64) (result *Member) {
	this.MemberGradeId = memberGradeId
	return this
}
func (this *Member) GetMemberGradeId() (memberGradeId int64) {
	return this.MemberGradeId
}
func (this *Member) SetKycStatus(kycStatus int) (result *Member) {
	this.KycStatus = kycStatus
	return this
}
func (this *Member) GetKycStatus() (kycStatus int) {
	return this.KycStatus
}
func (this *Member) SetGeneralizeTotal(generalizeTotal int64) (result *Member) {
	this.GeneralizeTotal = generalizeTotal
	return this
}
func (this *Member) GetGeneralizeTotal() (generalizeTotal int64) {
	return this.GeneralizeTotal
}
func (this *Member) SetInviterParentId(inviterParentId int64) (result *Member) {
	this.InviterParentId = inviterParentId
	return this
}
func (this *Member) GetInviterParentId() (inviterParentId int64) {
	return this.InviterParentId
}

type Member struct {
	Id                         int64
	Salt                       string
	Username                   string
	Password                   string
	Margin                     string
	GoogleKey                  string
	GoogleState                int
	GoogleDate                 time.Time
	JyPassword                 string
	RealName                   string
	IdNumber                   string
	Email                      string
	MobilePhone                string
	Location                   *Location
	MemberLevel                MemberLevelEnum.MemberLevelEnum
	Status                     CommonStatus.CommonStatus
	RegistrationTime           time.Time
	LastLoginTime              time.Time
	Token                      string
	Transactions               int
	BankInfo                   *BankInfo
	Alipay                     *Alipay
	WechatPay                  *WechatPay
	AppealTimes                int
	AppealSuccessTimes         int
	InviterId                  int64
	PromotionCode              string
	FirstLevel                 int
	SecondLevel                int
	ThirdLevel                 int
	Country                    *Country
	RealNameStatus             RealNameStatus.RealNameStatus
	LoginCount                 int
	CertifiedBusinessStatus    CertifiedBusinessStatus.CertifiedBusinessStatus
	CertifiedBusinessApplyTime time.Time
	ApplicationTime            time.Time
	CertifiedBusinessCheckTime time.Time
	Avatar                     string
	TokenExpireTime            time.Time
	PublishAdvertise           BooleanEnum.BooleanEnum
	TransactionStatus          BooleanEnum.BooleanEnum
	SignInAbility              bool
	TransactionTime            time.Time
	ChannelId                  int64
	IsChannel                  BooleanEnum.BooleanEnum
	ChannelVO                  *vo.ChannelVO
	LoginLock                  BooleanEnum.BooleanEnum
	Integration                int64
	MemberGradeId              int64
	KycStatus                  int
	GeneralizeTotal            int64
	InviterParentId            int64
}
