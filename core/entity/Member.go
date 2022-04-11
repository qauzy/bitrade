package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/constant/RealNameStatus"
	"bitrade/core/vo"
	"github.com/qauzy/chocolate/xtime"
)

func (this *Member) SetId(Id int64) (result *Member) {
	this.Id = Id
	return this
}
func (this *Member) GetId() (Id int64) {
	return this.Id
}
func (this *Member) SetSalt(Salt string) (result *Member) {
	this.Salt = Salt
	return this
}
func (this *Member) GetSalt() (Salt string) {
	return this.Salt
}
func (this *Member) SetUsername(Username string) (result *Member) {
	this.Username = Username
	return this
}
func (this *Member) GetUsername() (Username string) {
	return this.Username
}
func (this *Member) SetPassword(Password string) (result *Member) {
	this.Password = Password
	return this
}
func (this *Member) GetPassword() (Password string) {
	return this.Password
}
func (this *Member) SetMargin(Margin string) (result *Member) {
	this.Margin = Margin
	return this
}
func (this *Member) GetMargin() (Margin string) {
	return this.Margin
}
func (this *Member) SetGoogleKey(GoogleKey string) (result *Member) {
	this.GoogleKey = GoogleKey
	return this
}
func (this *Member) GetGoogleKey() (GoogleKey string) {
	return this.GoogleKey
}
func (this *Member) SetGoogleState(GoogleState int) (result *Member) {
	this.GoogleState = GoogleState
	return this
}
func (this *Member) GetGoogleState() (GoogleState int) {
	return this.GoogleState
}
func (this *Member) SetGoogleDate(GoogleDate xtime.Xtime) (result *Member) {
	this.GoogleDate = GoogleDate
	return this
}
func (this *Member) GetGoogleDate() (GoogleDate xtime.Xtime) {
	return this.GoogleDate
}
func (this *Member) SetJyPassword(JyPassword string) (result *Member) {
	this.JyPassword = JyPassword
	return this
}
func (this *Member) GetJyPassword() (JyPassword string) {
	return this.JyPassword
}
func (this *Member) SetRealName(RealName string) (result *Member) {
	this.RealName = RealName
	return this
}
func (this *Member) GetRealName() (RealName string) {
	return this.RealName
}
func (this *Member) SetIdNumber(IdNumber string) (result *Member) {
	this.IdNumber = IdNumber
	return this
}
func (this *Member) GetIdNumber() (IdNumber string) {
	return this.IdNumber
}
func (this *Member) SetEmail(Email string) (result *Member) {
	this.Email = Email
	return this
}
func (this *Member) GetEmail() (Email string) {
	return this.Email
}
func (this *Member) SetMobilePhone(MobilePhone string) (result *Member) {
	this.MobilePhone = MobilePhone
	return this
}
func (this *Member) GetMobilePhone() (MobilePhone string) {
	return this.MobilePhone
}
func (this *Member) SetLocation(Location *Location) (result *Member) {
	this.Location = Location
	return this
}
func (this *Member) GetLocation() (Location *Location) {
	return this.Location
}
func (this *Member) SetMemberLevel(MemberLevel MemberLevelEnum.MemberLevelEnum) (result *Member) {
	this.MemberLevel = MemberLevel
	return this
}
func (this *Member) GetMemberLevel() (MemberLevel MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *Member) SetStatus(Status CommonStatus.CommonStatus) (result *Member) {
	this.Status = Status
	return this
}
func (this *Member) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *Member) SetRegistrationTime(RegistrationTime xtime.Xtime) (result *Member) {
	this.RegistrationTime = RegistrationTime
	return this
}
func (this *Member) GetRegistrationTime() (RegistrationTime xtime.Xtime) {
	return this.RegistrationTime
}
func (this *Member) SetLastLoginTime(LastLoginTime xtime.Xtime) (result *Member) {
	this.LastLoginTime = LastLoginTime
	return this
}
func (this *Member) GetLastLoginTime() (LastLoginTime xtime.Xtime) {
	return this.LastLoginTime
}
func (this *Member) SetToken(Token string) (result *Member) {
	this.Token = Token
	return this
}
func (this *Member) GetToken() (Token string) {
	return this.Token
}
func (this *Member) SetTransactions(Transactions int) (result *Member) {
	this.Transactions = Transactions
	return this
}
func (this *Member) GetTransactions() (Transactions int) {
	return this.Transactions
}
func (this *Member) SetBankInfo(BankInfo *BankInfo) (result *Member) {
	this.BankInfo = BankInfo
	return this
}
func (this *Member) GetBankInfo() (BankInfo *BankInfo) {
	return this.BankInfo
}
func (this *Member) SetAlipay(Alipay *Alipay) (result *Member) {
	this.Alipay = Alipay
	return this
}
func (this *Member) GetAlipay() (Alipay *Alipay) {
	return this.Alipay
}
func (this *Member) SetWechatPay(WechatPay *WechatPay) (result *Member) {
	this.WechatPay = WechatPay
	return this
}
func (this *Member) GetWechatPay() (WechatPay *WechatPay) {
	return this.WechatPay
}
func (this *Member) SetAppealTimes(AppealTimes int) (result *Member) {
	this.AppealTimes = AppealTimes
	return this
}
func (this *Member) GetAppealTimes() (AppealTimes int) {
	return this.AppealTimes
}
func (this *Member) SetAppealSuccessTimes(AppealSuccessTimes int) (result *Member) {
	this.AppealSuccessTimes = AppealSuccessTimes
	return this
}
func (this *Member) GetAppealSuccessTimes() (AppealSuccessTimes int) {
	return this.AppealSuccessTimes
}
func (this *Member) SetInviterId(InviterId int64) (result *Member) {
	this.InviterId = InviterId
	return this
}
func (this *Member) GetInviterId() (InviterId int64) {
	return this.InviterId
}
func (this *Member) SetPromotionCode(PromotionCode string) (result *Member) {
	this.PromotionCode = PromotionCode
	return this
}
func (this *Member) GetPromotionCode() (PromotionCode string) {
	return this.PromotionCode
}
func (this *Member) SetFirstLevel(FirstLevel int) (result *Member) {
	this.FirstLevel = FirstLevel
	return this
}
func (this *Member) GetFirstLevel() (FirstLevel int) {
	return this.FirstLevel
}
func (this *Member) SetSecondLevel(SecondLevel int) (result *Member) {
	this.SecondLevel = SecondLevel
	return this
}
func (this *Member) GetSecondLevel() (SecondLevel int) {
	return this.SecondLevel
}
func (this *Member) SetThirdLevel(ThirdLevel int) (result *Member) {
	this.ThirdLevel = ThirdLevel
	return this
}
func (this *Member) GetThirdLevel() (ThirdLevel int) {
	return this.ThirdLevel
}
func (this *Member) SetCountry(Country *Country) (result *Member) {
	this.Country = Country
	return this
}
func (this *Member) GetCountry() (Country *Country) {
	return this.Country
}
func (this *Member) SetRealNameStatus(RealNameStatus RealNameStatus.RealNameStatus) (result *Member) {
	this.RealNameStatus = RealNameStatus
	return this
}
func (this *Member) GetRealNameStatus() (RealNameStatus RealNameStatus.RealNameStatus) {
	return this.RealNameStatus
}
func (this *Member) SetLoginCount(LoginCount int) (result *Member) {
	this.LoginCount = LoginCount
	return this
}
func (this *Member) GetLoginCount() (LoginCount int) {
	return this.LoginCount
}
func (this *Member) SetCertifiedBusinessStatus(CertifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus) (result *Member) {
	this.CertifiedBusinessStatus = CertifiedBusinessStatus
	return this
}
func (this *Member) GetCertifiedBusinessStatus() (CertifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *Member) SetCertifiedBusinessApplyTime(CertifiedBusinessApplyTime xtime.Xtime) (result *Member) {
	this.CertifiedBusinessApplyTime = CertifiedBusinessApplyTime
	return this
}
func (this *Member) GetCertifiedBusinessApplyTime() (CertifiedBusinessApplyTime xtime.Xtime) {
	return this.CertifiedBusinessApplyTime
}
func (this *Member) SetApplicationTime(ApplicationTime xtime.Xtime) (result *Member) {
	this.ApplicationTime = ApplicationTime
	return this
}
func (this *Member) GetApplicationTime() (ApplicationTime xtime.Xtime) {
	return this.ApplicationTime
}
func (this *Member) SetCertifiedBusinessCheckTime(CertifiedBusinessCheckTime xtime.Xtime) (result *Member) {
	this.CertifiedBusinessCheckTime = CertifiedBusinessCheckTime
	return this
}
func (this *Member) GetCertifiedBusinessCheckTime() (CertifiedBusinessCheckTime xtime.Xtime) {
	return this.CertifiedBusinessCheckTime
}
func (this *Member) SetAvatar(Avatar string) (result *Member) {
	this.Avatar = Avatar
	return this
}
func (this *Member) GetAvatar() (Avatar string) {
	return this.Avatar
}
func (this *Member) SetTokenExpireTime(TokenExpireTime xtime.Xtime) (result *Member) {
	this.TokenExpireTime = TokenExpireTime
	return this
}
func (this *Member) GetTokenExpireTime() (TokenExpireTime xtime.Xtime) {
	return this.TokenExpireTime
}
func (this *Member) SetPublishAdvertise(PublishAdvertise BooleanEnum.BooleanEnum) (result *Member) {
	this.PublishAdvertise = PublishAdvertise
	return this
}
func (this *Member) GetPublishAdvertise() (PublishAdvertise BooleanEnum.BooleanEnum) {
	return this.PublishAdvertise
}
func (this *Member) SetTransactionStatus(TransactionStatus BooleanEnum.BooleanEnum) (result *Member) {
	this.TransactionStatus = TransactionStatus
	return this
}
func (this *Member) GetTransactionStatus() (TransactionStatus BooleanEnum.BooleanEnum) {
	return this.TransactionStatus
}
func (this *Member) SetSignInAbility(SignInAbility bool) (result *Member) {
	this.SignInAbility = SignInAbility
	return this
}
func (this *Member) GetSignInAbility() (SignInAbility bool) {
	return this.SignInAbility
}
func (this *Member) SetTransactionTime(TransactionTime xtime.Xtime) (result *Member) {
	this.TransactionTime = TransactionTime
	return this
}
func (this *Member) GetTransactionTime() (TransactionTime xtime.Xtime) {
	return this.TransactionTime
}
func (this *Member) SetChannelId(ChannelId int64) (result *Member) {
	this.ChannelId = ChannelId
	return this
}
func (this *Member) GetChannelId() (ChannelId int64) {
	return this.ChannelId
}
func (this *Member) SetIsChannel(IsChannel BooleanEnum.BooleanEnum) (result *Member) {
	this.IsChannel = IsChannel
	return this
}
func (this *Member) GetIsChannel() (IsChannel BooleanEnum.BooleanEnum) {
	return this.IsChannel
}
func (this *Member) SetChannelVO(ChannelVO *vo.ChannelVO) (result *Member) {
	this.ChannelVO = ChannelVO
	return this
}
func (this *Member) GetChannelVO() (ChannelVO *vo.ChannelVO) {
	return this.ChannelVO
}
func (this *Member) SetLoginLock(LoginLock BooleanEnum.BooleanEnum) (result *Member) {
	this.LoginLock = LoginLock
	return this
}
func (this *Member) GetLoginLock() (LoginLock BooleanEnum.BooleanEnum) {
	return this.LoginLock
}
func (this *Member) SetIntegration(Integration int64) (result *Member) {
	this.Integration = Integration
	return this
}
func (this *Member) GetIntegration() (Integration int64) {
	return this.Integration
}
func (this *Member) SetMemberGradeId(MemberGradeId int64) (result *Member) {
	this.MemberGradeId = MemberGradeId
	return this
}
func (this *Member) GetMemberGradeId() (MemberGradeId int64) {
	return this.MemberGradeId
}
func (this *Member) SetKycStatus(KycStatus int) (result *Member) {
	this.KycStatus = KycStatus
	return this
}
func (this *Member) GetKycStatus() (KycStatus int) {
	return this.KycStatus
}
func (this *Member) SetGeneralizeTotal(GeneralizeTotal int64) (result *Member) {
	this.GeneralizeTotal = GeneralizeTotal
	return this
}
func (this *Member) GetGeneralizeTotal() (GeneralizeTotal int64) {
	return this.GeneralizeTotal
}
func (this *Member) SetInviterParentId(InviterParentId int64) (result *Member) {
	this.InviterParentId = InviterParentId
	return this
}
func (this *Member) GetInviterParentId() (InviterParentId int64) {
	return this.InviterParentId
}
func NewMember(id int64, salt string, username string, password string, margin string, googleKey string, googleState int, googleDate xtime.Xtime, jyPassword string, realName string, idNumber string, email string, mobilePhone string, location *Location, memberLevel MemberLevelEnum.MemberLevelEnum, status CommonStatus.CommonStatus, registrationTime xtime.Xtime, lastLoginTime xtime.Xtime, token string, transactions int, bankInfo *BankInfo, alipay *Alipay, wechatPay *WechatPay, appealTimes int, appealSuccessTimes int, inviterId int64, promotionCode string, firstLevel int, secondLevel int, thirdLevel int, country *Country, realNameStatus RealNameStatus.RealNameStatus, loginCount int, certifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus, certifiedBusinessApplyTime xtime.Xtime, applicationTime xtime.Xtime, certifiedBusinessCheckTime xtime.Xtime, avatar string, tokenExpireTime xtime.Xtime, publishAdvertise BooleanEnum.BooleanEnum, transactionStatus BooleanEnum.BooleanEnum, signInAbility bool, transactionTime xtime.Xtime, channelId int64, isChannel BooleanEnum.BooleanEnum, channelVO *vo.ChannelVO, loginLock BooleanEnum.BooleanEnum, integration int64, memberGradeId int64, kycStatus int, generalizeTotal int64, inviterParentId int64) (ret *Member) {
	ret = new(Member)
	ret.Id = id
	ret.Salt = salt
	ret.Username = username
	ret.Password = password
	ret.Margin = margin
	ret.GoogleKey = googleKey
	ret.GoogleState = googleState
	ret.GoogleDate = googleDate
	ret.JyPassword = jyPassword
	ret.RealName = realName
	ret.IdNumber = idNumber
	ret.Email = email
	ret.MobilePhone = mobilePhone
	ret.Location = location
	ret.MemberLevel = memberLevel
	ret.Status = status
	ret.RegistrationTime = registrationTime
	ret.LastLoginTime = lastLoginTime
	ret.Token = token
	ret.Transactions = transactions
	ret.BankInfo = bankInfo
	ret.Alipay = alipay
	ret.WechatPay = wechatPay
	ret.AppealTimes = appealTimes
	ret.AppealSuccessTimes = appealSuccessTimes
	ret.InviterId = inviterId
	ret.PromotionCode = promotionCode
	ret.FirstLevel = firstLevel
	ret.SecondLevel = secondLevel
	ret.ThirdLevel = thirdLevel
	ret.Country = country
	ret.RealNameStatus = realNameStatus
	ret.LoginCount = loginCount
	ret.CertifiedBusinessStatus = certifiedBusinessStatus
	ret.CertifiedBusinessApplyTime = certifiedBusinessApplyTime
	ret.ApplicationTime = applicationTime
	ret.CertifiedBusinessCheckTime = certifiedBusinessCheckTime
	ret.Avatar = avatar
	ret.TokenExpireTime = tokenExpireTime
	ret.PublishAdvertise = publishAdvertise
	ret.TransactionStatus = transactionStatus
	ret.SignInAbility = signInAbility
	ret.TransactionTime = transactionTime
	ret.ChannelId = channelId
	ret.IsChannel = isChannel
	ret.ChannelVO = channelVO
	ret.LoginLock = loginLock
	ret.Integration = integration
	ret.MemberGradeId = memberGradeId
	ret.KycStatus = kycStatus
	ret.GeneralizeTotal = generalizeTotal
	ret.InviterParentId = inviterParentId
	return
}

type Member struct {
	Id                         int64                                           `gorm:"column:id" json:"id"`
	Salt                       string                                          `gorm:"column:salt" json:"salt"`
	Username                   string                                          `gorm:"column:username" json:"username"`
	Password                   string                                          `gorm:"column:password" json:"password"`
	Margin                     string                                          `gorm:"column:margin" json:"margin"`
	GoogleKey                  string                                          `gorm:"column:google_key" json:"googleKey"`
	GoogleState                int                                             `gorm:"column:google_state" json:"googleState"`
	GoogleDate                 xtime.Xtime                                     `gorm:"column:google_date" json:"googleDate"`
	JyPassword                 string                                          `gorm:"column:jy_password" json:"jyPassword"`
	RealName                   string                                          `gorm:"column:real_name" json:"realName"`
	IdNumber                   string                                          `gorm:"column:id_number" json:"idNumber"`
	Email                      string                                          `gorm:"column:email" json:"email"`
	MobilePhone                string                                          `gorm:"column:mobile_phone" json:"mobilePhone"`
	Location                   *Location                                       `gorm:"column:location" json:"location"`
	MemberLevel                MemberLevelEnum.MemberLevelEnum                 `gorm:"column:member_level" json:"memberLevel"`
	Status                     CommonStatus.CommonStatus                       `gorm:"column:status" json:"status"`
	RegistrationTime           xtime.Xtime                                     `gorm:"column:registration_time" json:"registrationTime"`
	LastLoginTime              xtime.Xtime                                     `gorm:"column:last_login_time" json:"lastLoginTime"`
	Token                      string                                          `gorm:"column:token" json:"token"`
	Transactions               int                                             `gorm:"column:transactions" json:"transactions"`
	BankInfo                   *BankInfo                                       `gorm:"column:bank_info" json:"bankInfo"`
	Alipay                     *Alipay                                         `gorm:"column:alipay" json:"alipay"`
	WechatPay                  *WechatPay                                      `gorm:"column:wechat_pay" json:"wechatPay"`
	AppealTimes                int                                             `gorm:"column:appeal_times" json:"appealTimes"`
	AppealSuccessTimes         int                                             `gorm:"column:appeal_success_times" json:"appealSuccessTimes"`
	InviterId                  int64                                           `gorm:"column:inviter_id" json:"inviterId"`
	PromotionCode              string                                          `gorm:"column:promotion_code" json:"promotionCode"`
	FirstLevel                 int                                             `gorm:"column:first_level" json:"firstLevel"`
	SecondLevel                int                                             `gorm:"column:second_level" json:"secondLevel"`
	ThirdLevel                 int                                             `gorm:"column:third_level" json:"thirdLevel"`
	Country                    *Country                                        `gorm:"column:country" json:"country"`
	RealNameStatus             RealNameStatus.RealNameStatus                   `gorm:"column:real_name_status" json:"realNameStatus"`
	LoginCount                 int                                             `gorm:"column:login_count" json:"loginCount"`
	CertifiedBusinessStatus    CertifiedBusinessStatus.CertifiedBusinessStatus `gorm:"column:certified_business_status" json:"certifiedBusinessStatus"`
	CertifiedBusinessApplyTime xtime.Xtime                                     `gorm:"column:certified_business_apply_time" json:"certifiedBusinessApplyTime"`
	ApplicationTime            xtime.Xtime                                     `gorm:"column:application_time" json:"applicationTime"`
	CertifiedBusinessCheckTime xtime.Xtime                                     `gorm:"column:certified_business_check_time" json:"certifiedBusinessCheckTime"`
	Avatar                     string                                          `gorm:"column:avatar" json:"avatar"`
	TokenExpireTime            xtime.Xtime                                     `gorm:"column:token_expire_time" json:"tokenExpireTime"`
	PublishAdvertise           BooleanEnum.BooleanEnum                         `gorm:"column:publish_advertise" json:"publishAdvertise"`
	TransactionStatus          BooleanEnum.BooleanEnum                         `gorm:"column:transaction_status" json:"transactionStatus"`
	SignInAbility              bool                                            `gorm:"column:sign_in_ability" json:"signInAbility"`
	TransactionTime            xtime.Xtime                                     `gorm:"column:transaction_time" json:"transactionTime"`
	ChannelId                  int64                                           `gorm:"column:channel_id" json:"channelId"`
	IsChannel                  BooleanEnum.BooleanEnum                         `gorm:"column:is_channel" json:"isChannel"`
	ChannelVO                  *vo.ChannelVO                                   `gorm:"column:channel_v_o" json:"channelVO"`
	LoginLock                  BooleanEnum.BooleanEnum                         `gorm:"column:login_lock" json:"loginLock"`
	Integration                int64                                           `gorm:"column:integration" json:"integration"`
	MemberGradeId              int64                                           `gorm:"column:member_grade_id" json:"memberGradeId"`
	KycStatus                  int                                             `gorm:"column:kyc_status" json:"kycStatus"`
	GeneralizeTotal            int64                                           `gorm:"column:generalize_total" json:"generalizeTotal"`
	InviterParentId            int64                                           `gorm:"column:inviter_parent_id" json:"inviterParentId"`
}
