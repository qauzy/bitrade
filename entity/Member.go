package entity

func (this *Member) SetId(id int64) {
	this.Id = id
}
func (this *Member) GetId() (id int64) {
	return this.Id
}
func (this *Member) SetSalt(salt string) {
	this.Salt = salt
}
func (this *Member) GetSalt() (salt string) {
	return this.Salt
}
func (this *Member) SetUsername(username string) {
	this.Username = username
}
func (this *Member) GetUsername() (username string) {
	return this.Username
}
func (this *Member) SetPassword(password string) {
	this.Password = password
}
func (this *Member) GetPassword() (password string) {
	return this.Password
}
func (this *Member) SetMargin(margin string) {
	this.Margin = margin
}
func (this *Member) GetMargin() (margin string) {
	return this.Margin
}
func (this *Member) SetGoogleKey(googleKey string) {
	this.GoogleKey = googleKey
}
func (this *Member) GetGoogleKey() (googleKey string) {
	return this.GoogleKey
}
func (this *Member) SetGoogleState(googleState int64) {
	this.GoogleState = googleState
}
func (this *Member) GetGoogleState() (googleState int64) {
	return this.GoogleState
}
func (this *Member) SetGoogleDate(googleDate Date) {
	this.GoogleDate = googleDate
}
func (this *Member) GetGoogleDate() (googleDate Date) {
	return this.GoogleDate
}
func (this *Member) SetJyPassword(jyPassword string) {
	this.JyPassword = jyPassword
}
func (this *Member) GetJyPassword() (jyPassword string) {
	return this.JyPassword
}
func (this *Member) SetRealName(realName string) {
	this.RealName = realName
}
func (this *Member) GetRealName() (realName string) {
	return this.RealName
}
func (this *Member) SetIdNumber(idNumber string) {
	this.IdNumber = idNumber
}
func (this *Member) GetIdNumber() (idNumber string) {
	return this.IdNumber
}
func (this *Member) SetEmail(email string) {
	this.Email = email
}
func (this *Member) GetEmail() (email string) {
	return this.Email
}
func (this *Member) SetMobilePhone(mobilePhone string) {
	this.MobilePhone = mobilePhone
}
func (this *Member) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *Member) SetLocation(location Location) {
	this.Location = location
}
func (this *Member) GetLocation() (location Location) {
	return this.Location
}
func (this *Member) SetMemberLevel(memberLevel MemberLevelEnum) {
	this.MemberLevel = memberLevel
}
func (this *Member) GetMemberLevel() (memberLevel MemberLevelEnum) {
	return this.MemberLevel
}
func (this *Member) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *Member) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *Member) SetRegistrationTime(registrationTime Date) {
	this.RegistrationTime = registrationTime
}
func (this *Member) GetRegistrationTime() (registrationTime Date) {
	return this.RegistrationTime
}
func (this *Member) SetLastLoginTime(lastLoginTime Date) {
	this.LastLoginTime = lastLoginTime
}
func (this *Member) GetLastLoginTime() (lastLoginTime Date) {
	return this.LastLoginTime
}
func (this *Member) SetToken(token string) {
	this.Token = token
}
func (this *Member) GetToken() (token string) {
	return this.Token
}
func (this *Member) SetTransactions(transactions int) {
	this.Transactions = transactions
}
func (this *Member) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *Member) SetBankInfo(bankInfo BankInfo) {
	this.BankInfo = bankInfo
}
func (this *Member) GetBankInfo() (bankInfo BankInfo) {
	return this.BankInfo
}
func (this *Member) SetAlipay(alipay Alipay) {
	this.Alipay = alipay
}
func (this *Member) GetAlipay() (alipay Alipay) {
	return this.Alipay
}
func (this *Member) SetWechatPay(wechatPay WechatPay) {
	this.WechatPay = wechatPay
}
func (this *Member) GetWechatPay() (wechatPay WechatPay) {
	return this.WechatPay
}
func (this *Member) SetAppealTimes(appealTimes int) {
	this.AppealTimes = appealTimes
}
func (this *Member) GetAppealTimes() (appealTimes int) {
	return this.AppealTimes
}
func (this *Member) SetAppealSuccessTimes(appealSuccessTimes int) {
	this.AppealSuccessTimes = appealSuccessTimes
}
func (this *Member) GetAppealSuccessTimes() (appealSuccessTimes int) {
	return this.AppealSuccessTimes
}
func (this *Member) SetInviterId(inviterId int64) {
	this.InviterId = inviterId
}
func (this *Member) GetInviterId() (inviterId int64) {
	return this.InviterId
}
func (this *Member) SetPromotionCode(promotionCode string) {
	this.PromotionCode = promotionCode
}
func (this *Member) GetPromotionCode() (promotionCode string) {
	return this.PromotionCode
}
func (this *Member) SetFirstLevel(firstLevel int) {
	this.FirstLevel = firstLevel
}
func (this *Member) GetFirstLevel() (firstLevel int) {
	return this.FirstLevel
}
func (this *Member) SetSecondLevel(secondLevel int) {
	this.SecondLevel = secondLevel
}
func (this *Member) GetSecondLevel() (secondLevel int) {
	return this.SecondLevel
}
func (this *Member) SetThirdLevel(thirdLevel int) {
	this.ThirdLevel = thirdLevel
}
func (this *Member) GetThirdLevel() (thirdLevel int) {
	return this.ThirdLevel
}
func (this *Member) SetCountry(country Country) {
	this.Country = country
}
func (this *Member) GetCountry() (country Country) {
	return this.Country
}
func (this *Member) SetRealNameStatus(realNameStatus RealNameStatus) {
	this.RealNameStatus = realNameStatus
}
func (this *Member) GetRealNameStatus() (realNameStatus RealNameStatus) {
	return this.RealNameStatus
}
func (this *Member) SetLoginCount(loginCount int) {
	this.LoginCount = loginCount
}
func (this *Member) GetLoginCount() (loginCount int) {
	return this.LoginCount
}
func (this *Member) SetCertifiedBusinessStatus(certifiedBusinessStatus CertifiedBusinessStatus) {
	this.CertifiedBusinessStatus = certifiedBusinessStatus
}
func (this *Member) GetCertifiedBusinessStatus() (certifiedBusinessStatus CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *Member) SetCertifiedBusinessApplyTime(certifiedBusinessApplyTime Date) {
	this.CertifiedBusinessApplyTime = certifiedBusinessApplyTime
}
func (this *Member) GetCertifiedBusinessApplyTime() (certifiedBusinessApplyTime Date) {
	return this.CertifiedBusinessApplyTime
}
func (this *Member) SetApplicationTime(applicationTime Date) {
	this.ApplicationTime = applicationTime
}
func (this *Member) GetApplicationTime() (applicationTime Date) {
	return this.ApplicationTime
}
func (this *Member) SetCertifiedBusinessCheckTime(certifiedBusinessCheckTime Date) {
	this.CertifiedBusinessCheckTime = certifiedBusinessCheckTime
}
func (this *Member) GetCertifiedBusinessCheckTime() (certifiedBusinessCheckTime Date) {
	return this.CertifiedBusinessCheckTime
}
func (this *Member) SetAvatar(avatar string) {
	this.Avatar = avatar
}
func (this *Member) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *Member) SetTokenExpireTime(tokenExpireTime Date) {
	this.TokenExpireTime = tokenExpireTime
}
func (this *Member) GetTokenExpireTime() (tokenExpireTime Date) {
	return this.TokenExpireTime
}
func (this *Member) SetPublishAdvertise(publishAdvertise BooleanEnum) {
	this.PublishAdvertise = publishAdvertise
}
func (this *Member) GetPublishAdvertise() (publishAdvertise BooleanEnum) {
	return this.PublishAdvertise
}
func (this *Member) SetTransactionStatus(transactionStatus BooleanEnum) {
	this.TransactionStatus = transactionStatus
}
func (this *Member) GetTransactionStatus() (transactionStatus BooleanEnum) {
	return this.TransactionStatus
}
func (this *Member) SetSignInAbility(signInAbility bool) {
	this.SignInAbility = signInAbility
}
func (this *Member) GetSignInAbility() (signInAbility bool) {
	return this.SignInAbility
}
func (this *Member) SetTransactionTime(transactionTime Date) {
	this.TransactionTime = transactionTime
}
func (this *Member) GetTransactionTime() (transactionTime Date) {
	return this.TransactionTime
}
func (this *Member) SetChannelId(channelId int64) {
	this.ChannelId = channelId
}
func (this *Member) GetChannelId() (channelId int64) {
	return this.ChannelId
}
func (this *Member) SetIsChannel(isChannel BooleanEnum) {
	this.IsChannel = isChannel
}
func (this *Member) GetIsChannel() (isChannel BooleanEnum) {
	return this.IsChannel
}
func (this *Member) SetChannelVO(channelVO ChannelVO) {
	this.ChannelVO = channelVO
}
func (this *Member) GetChannelVO() (channelVO ChannelVO) {
	return this.ChannelVO
}
func (this *Member) SetLoginLock(loginLock BooleanEnum) {
	this.LoginLock = loginLock
}
func (this *Member) GetLoginLock() (loginLock BooleanEnum) {
	return this.LoginLock
}
func (this *Member) SetIntegration(integration int64) {
	this.Integration = integration
}
func (this *Member) GetIntegration() (integration int64) {
	return this.Integration
}
func (this *Member) SetMemberGradeId(memberGradeId int64) {
	this.MemberGradeId = memberGradeId
}
func (this *Member) GetMemberGradeId() (memberGradeId int64) {
	return this.MemberGradeId
}
func (this *Member) SetKycStatus(kycStatus int64) {
	this.KycStatus = kycStatus
}
func (this *Member) GetKycStatus() (kycStatus int64) {
	return this.KycStatus
}
func (this *Member) SetGeneralizeTotal(generalizeTotal int64) {
	this.GeneralizeTotal = generalizeTotal
}
func (this *Member) GetGeneralizeTotal() (generalizeTotal int64) {
	return this.GeneralizeTotal
}
func (this *Member) SetInviterParentId(inviterParentId int64) {
	this.InviterParentId = inviterParentId
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
	GoogleState                int64
	GoogleDate                 Date
	JyPassword                 string
	RealName                   string
	IdNumber                   string
	Email                      string
	MobilePhone                string
	Location                   Location
	MemberLevel                MemberLevelEnum
	Status                     CommonStatus
	RegistrationTime           Date
	LastLoginTime              Date
	Token                      string
	Transactions               int
	BankInfo                   BankInfo
	Alipay                     Alipay
	WechatPay                  WechatPay
	AppealTimes                int
	AppealSuccessTimes         int
	InviterId                  int64
	PromotionCode              string
	FirstLevel                 int
	SecondLevel                int
	ThirdLevel                 int
	Country                    Country
	RealNameStatus             RealNameStatus
	LoginCount                 int
	CertifiedBusinessStatus    CertifiedBusinessStatus
	CertifiedBusinessApplyTime Date
	ApplicationTime            Date
	CertifiedBusinessCheckTime Date
	Avatar                     string
	TokenExpireTime            Date
	PublishAdvertise           BooleanEnum
	TransactionStatus          BooleanEnum
	SignInAbility              bool
	TransactionTime            Date
	ChannelId                  int64
	IsChannel                  BooleanEnum
	ChannelVO                  ChannelVO
	LoginLock                  BooleanEnum
	Integration                int64
	MemberGradeId              int64
	KycStatus                  int64
	GeneralizeTotal            int64
	InviterParentId            int64
}
