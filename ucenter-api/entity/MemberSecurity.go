package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/MemberLevelEnum"
	"time"
)

func (this *MemberSecurity) SetUsername(Username string) (result *MemberSecurity) {
	this.Username = Username
	return this
}
func (this *MemberSecurity) GetUsername() (Username string) {
	return this.Username
}
func (this *MemberSecurity) SetId(Id int64) (result *MemberSecurity) {
	this.Id = Id
	return this
}
func (this *MemberSecurity) GetId() (Id int64) {
	return this.Id
}
func (this *MemberSecurity) SetCreateTime(CreateTime time.Time) (result *MemberSecurity) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberSecurity) GetCreateTime() (CreateTime time.Time) {
	return this.CreateTime
}
func (this *MemberSecurity) SetRealVerified(RealVerified *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.RealVerified = RealVerified
	return this
}
func (this *MemberSecurity) GetRealVerified() (RealVerified *BooleanEnum.BooleanEnum) {
	return this.RealVerified
}
func (this *MemberSecurity) SetEmailVerified(EmailVerified *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.EmailVerified = EmailVerified
	return this
}
func (this *MemberSecurity) GetEmailVerified() (EmailVerified *BooleanEnum.BooleanEnum) {
	return this.EmailVerified
}
func (this *MemberSecurity) SetPhoneVerified(PhoneVerified *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.PhoneVerified = PhoneVerified
	return this
}
func (this *MemberSecurity) GetPhoneVerified() (PhoneVerified *BooleanEnum.BooleanEnum) {
	return this.PhoneVerified
}
func (this *MemberSecurity) SetLoginVerified(LoginVerified *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.LoginVerified = LoginVerified
	return this
}
func (this *MemberSecurity) GetLoginVerified() (LoginVerified *BooleanEnum.BooleanEnum) {
	return this.LoginVerified
}
func (this *MemberSecurity) SetFundsVerified(FundsVerified *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.FundsVerified = FundsVerified
	return this
}
func (this *MemberSecurity) GetFundsVerified() (FundsVerified *BooleanEnum.BooleanEnum) {
	return this.FundsVerified
}
func (this *MemberSecurity) SetRealAuditing(RealAuditing *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.RealAuditing = RealAuditing
	return this
}
func (this *MemberSecurity) GetRealAuditing() (RealAuditing *BooleanEnum.BooleanEnum) {
	return this.RealAuditing
}
func (this *MemberSecurity) SetMobilePhone(MobilePhone string) (result *MemberSecurity) {
	this.MobilePhone = MobilePhone
	return this
}
func (this *MemberSecurity) GetMobilePhone() (MobilePhone string) {
	return this.MobilePhone
}
func (this *MemberSecurity) SetEmail(Email string) (result *MemberSecurity) {
	this.Email = Email
	return this
}
func (this *MemberSecurity) GetEmail() (Email string) {
	return this.Email
}
func (this *MemberSecurity) SetRealName(RealName string) (result *MemberSecurity) {
	this.RealName = RealName
	return this
}
func (this *MemberSecurity) GetRealName() (RealName string) {
	return this.RealName
}
func (this *MemberSecurity) SetRealNameRejectReason(RealNameRejectReason string) (result *MemberSecurity) {
	this.RealNameRejectReason = RealNameRejectReason
	return this
}
func (this *MemberSecurity) GetRealNameRejectReason() (RealNameRejectReason string) {
	return this.RealNameRejectReason
}
func (this *MemberSecurity) SetIdCard(IdCard string) (result *MemberSecurity) {
	this.IdCard = IdCard
	return this
}
func (this *MemberSecurity) GetIdCard() (IdCard string) {
	return this.IdCard
}
func (this *MemberSecurity) SetAvatar(Avatar string) (result *MemberSecurity) {
	this.Avatar = Avatar
	return this
}
func (this *MemberSecurity) GetAvatar() (Avatar string) {
	return this.Avatar
}
func (this *MemberSecurity) SetAccountVerified(AccountVerified *BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.AccountVerified = AccountVerified
	return this
}
func (this *MemberSecurity) GetAccountVerified() (AccountVerified *BooleanEnum.BooleanEnum) {
	return this.AccountVerified
}
func (this *MemberSecurity) SetGoogleStatus(GoogleStatus int) (result *MemberSecurity) {
	this.GoogleStatus = GoogleStatus
	return this
}
func (this *MemberSecurity) GetGoogleStatus() (GoogleStatus int) {
	return this.GoogleStatus
}
func (this *MemberSecurity) SetTransactions(Transactions int) (result *MemberSecurity) {
	this.Transactions = Transactions
	return this
}
func (this *MemberSecurity) GetTransactions() (Transactions int) {
	return this.Transactions
}
func (this *MemberSecurity) SetTransactionTime(TransactionTime time.Time) (result *MemberSecurity) {
	this.TransactionTime = TransactionTime
	return this
}
func (this *MemberSecurity) GetTransactionTime() (TransactionTime time.Time) {
	return this.TransactionTime
}
func (this *MemberSecurity) SetLevel(Level int) (result *MemberSecurity) {
	this.Level = Level
	return this
}
func (this *MemberSecurity) GetLevel() (Level int) {
	return this.Level
}
func (this *MemberSecurity) SetIntegration(Integration int64) (result *MemberSecurity) {
	this.Integration = Integration
	return this
}
func (this *MemberSecurity) GetIntegration() (Integration int64) {
	return this.Integration
}
func (this *MemberSecurity) SetKycStatus(KycStatus int) (result *MemberSecurity) {
	this.KycStatus = KycStatus
	return this
}
func (this *MemberSecurity) GetKycStatus() (KycStatus int) {
	return this.KycStatus
}
func (this *MemberSecurity) SetMemberGradeId(MemberGradeId int64) (result *MemberSecurity) {
	this.MemberGradeId = MemberGradeId
	return this
}
func (this *MemberSecurity) GetMemberGradeId() (MemberGradeId int64) {
	return this.MemberGradeId
}
func (this *MemberSecurity) SetGoogleState(GoogleState int) (result *MemberSecurity) {
	this.GoogleState = GoogleState
	return this
}
func (this *MemberSecurity) GetGoogleState() (GoogleState int) {
	return this.GoogleState
}
func (this *MemberSecurity) SetMemberLevel(MemberLevel *MemberLevelEnum.MemberLevelEnum) (result *MemberSecurity) {
	this.MemberLevel = MemberLevel
	return this
}
func (this *MemberSecurity) GetMemberLevel() (MemberLevel *MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}

type MemberSecurity struct {
	Username             string                           `gorm:"column:username" json:"username"`
	Id                   int64                            `gorm:"column:id" json:"id"`
	CreateTime           time.Time                        `gorm:"column:create_time" json:"createTime"`
	RealVerified         *BooleanEnum.BooleanEnum         `gorm:"column:real_verified" json:"realVerified"`
	EmailVerified        *BooleanEnum.BooleanEnum         `gorm:"column:email_verified" json:"emailVerified"`
	PhoneVerified        *BooleanEnum.BooleanEnum         `gorm:"column:phone_verified" json:"phoneVerified"`
	LoginVerified        *BooleanEnum.BooleanEnum         `gorm:"column:login_verified" json:"loginVerified"`
	FundsVerified        *BooleanEnum.BooleanEnum         `gorm:"column:funds_verified" json:"fundsVerified"`
	RealAuditing         *BooleanEnum.BooleanEnum         `gorm:"column:real_auditing" json:"realAuditing"`
	MobilePhone          string                           `gorm:"column:mobile_phone" json:"mobilePhone"`
	Email                string                           `gorm:"column:email" json:"email"`
	RealName             string                           `gorm:"column:real_name" json:"realName"`
	RealNameRejectReason string                           `gorm:"column:real_name_reject_reason" json:"realNameRejectReason"`
	IdCard               string                           `gorm:"column:id_card" json:"idCard"`
	Avatar               string                           `gorm:"column:avatar" json:"avatar"`
	AccountVerified      *BooleanEnum.BooleanEnum         `gorm:"column:account_verified" json:"accountVerified"`
	GoogleStatus         int                              `gorm:"column:google_status" json:"googleStatus"`
	Transactions         int                              `gorm:"column:transactions" json:"transactions"`
	TransactionTime      time.Time                        `gorm:"column:transaction_time" json:"transactionTime"`
	Level                int                              `gorm:"column:level" json:"level"`
	Integration          int64                            `gorm:"column:integration" json:"integration"`
	KycStatus            int                              `gorm:"column:kyc_status" json:"kycStatus"`
	MemberGradeId        int64                            `gorm:"column:member_grade_id" json:"memberGradeId"`
	GoogleState          int                              `gorm:"column:google_state" json:"googleState"`
	MemberLevel          *MemberLevelEnum.MemberLevelEnum `gorm:"column:member_level" json:"memberLevel"`
}
