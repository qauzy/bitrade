package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/MemberLevelEnum"
	"time"
)

func (this *MemberSecurity) SetUsername(username string) (result *MemberSecurity) {
	this.Username = username
	return this
}
func (this *MemberSecurity) GetUsername() (username string) {
	return this.Username
}
func (this *MemberSecurity) SetId(id int64) (result *MemberSecurity) {
	this.Id = id
	return this
}
func (this *MemberSecurity) GetId() (id int64) {
	return this.Id
}
func (this *MemberSecurity) SetCreateTime(createTime time.Time) (result *MemberSecurity) {
	this.CreateTime = createTime
	return this
}
func (this *MemberSecurity) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberSecurity) SetRealVerified(realVerified BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.RealVerified = realVerified
	return this
}
func (this *MemberSecurity) GetRealVerified() (realVerified BooleanEnum.BooleanEnum) {
	return this.RealVerified
}
func (this *MemberSecurity) SetEmailVerified(emailVerified BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.EmailVerified = emailVerified
	return this
}
func (this *MemberSecurity) GetEmailVerified() (emailVerified BooleanEnum.BooleanEnum) {
	return this.EmailVerified
}
func (this *MemberSecurity) SetPhoneVerified(phoneVerified BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.PhoneVerified = phoneVerified
	return this
}
func (this *MemberSecurity) GetPhoneVerified() (phoneVerified BooleanEnum.BooleanEnum) {
	return this.PhoneVerified
}
func (this *MemberSecurity) SetLoginVerified(loginVerified BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.LoginVerified = loginVerified
	return this
}
func (this *MemberSecurity) GetLoginVerified() (loginVerified BooleanEnum.BooleanEnum) {
	return this.LoginVerified
}
func (this *MemberSecurity) SetFundsVerified(fundsVerified BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.FundsVerified = fundsVerified
	return this
}
func (this *MemberSecurity) GetFundsVerified() (fundsVerified BooleanEnum.BooleanEnum) {
	return this.FundsVerified
}
func (this *MemberSecurity) SetRealAuditing(realAuditing BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.RealAuditing = realAuditing
	return this
}
func (this *MemberSecurity) GetRealAuditing() (realAuditing BooleanEnum.BooleanEnum) {
	return this.RealAuditing
}
func (this *MemberSecurity) SetMobilePhone(mobilePhone string) (result *MemberSecurity) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *MemberSecurity) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *MemberSecurity) SetEmail(email string) (result *MemberSecurity) {
	this.Email = email
	return this
}
func (this *MemberSecurity) GetEmail() (email string) {
	return this.Email
}
func (this *MemberSecurity) SetRealName(realName string) (result *MemberSecurity) {
	this.RealName = realName
	return this
}
func (this *MemberSecurity) GetRealName() (realName string) {
	return this.RealName
}
func (this *MemberSecurity) SetRealNameRejectReason(realNameRejectReason string) (result *MemberSecurity) {
	this.RealNameRejectReason = realNameRejectReason
	return this
}
func (this *MemberSecurity) GetRealNameRejectReason() (realNameRejectReason string) {
	return this.RealNameRejectReason
}
func (this *MemberSecurity) SetIdCard(idCard string) (result *MemberSecurity) {
	this.IdCard = idCard
	return this
}
func (this *MemberSecurity) GetIdCard() (idCard string) {
	return this.IdCard
}
func (this *MemberSecurity) SetAvatar(avatar string) (result *MemberSecurity) {
	this.Avatar = avatar
	return this
}
func (this *MemberSecurity) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *MemberSecurity) SetAccountVerified(accountVerified BooleanEnum.BooleanEnum) (result *MemberSecurity) {
	this.AccountVerified = accountVerified
	return this
}
func (this *MemberSecurity) GetAccountVerified() (accountVerified BooleanEnum.BooleanEnum) {
	return this.AccountVerified
}
func (this *MemberSecurity) SetGoogleStatus(googleStatus int64) (result *MemberSecurity) {
	this.GoogleStatus = googleStatus
	return this
}
func (this *MemberSecurity) GetGoogleStatus() (googleStatus int64) {
	return this.GoogleStatus
}
func (this *MemberSecurity) SetTransactions(transactions int) (result *MemberSecurity) {
	this.Transactions = transactions
	return this
}
func (this *MemberSecurity) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *MemberSecurity) SetTransactionTime(transactionTime time.Time) (result *MemberSecurity) {
	this.TransactionTime = transactionTime
	return this
}
func (this *MemberSecurity) GetTransactionTime() (transactionTime time.Time) {
	return this.TransactionTime
}
func (this *MemberSecurity) SetLevel(level int64) (result *MemberSecurity) {
	this.Level = level
	return this
}
func (this *MemberSecurity) GetLevel() (level int64) {
	return this.Level
}
func (this *MemberSecurity) SetIntegration(integration int64) (result *MemberSecurity) {
	this.Integration = integration
	return this
}
func (this *MemberSecurity) GetIntegration() (integration int64) {
	return this.Integration
}
func (this *MemberSecurity) SetKycStatus(kycStatus int64) (result *MemberSecurity) {
	this.KycStatus = kycStatus
	return this
}
func (this *MemberSecurity) GetKycStatus() (kycStatus int64) {
	return this.KycStatus
}
func (this *MemberSecurity) SetMemberGradeId(memberGradeId int64) (result *MemberSecurity) {
	this.MemberGradeId = memberGradeId
	return this
}
func (this *MemberSecurity) GetMemberGradeId() (memberGradeId int64) {
	return this.MemberGradeId
}
func (this *MemberSecurity) SetGoogleState(googleState int64) (result *MemberSecurity) {
	this.GoogleState = googleState
	return this
}
func (this *MemberSecurity) GetGoogleState() (googleState int64) {
	return this.GoogleState
}
func (this *MemberSecurity) SetMemberLevel(memberLevel MemberLevelEnum.MemberLevelEnum) (result *MemberSecurity) {
	this.MemberLevel = memberLevel
	return this
}
func (this *MemberSecurity) GetMemberLevel() (memberLevel MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}

type MemberSecurity struct {
	Username             string
	Id                   int64
	CreateTime           time.Time
	RealVerified         BooleanEnum.BooleanEnum
	EmailVerified        BooleanEnum.BooleanEnum
	PhoneVerified        BooleanEnum.BooleanEnum
	LoginVerified        BooleanEnum.BooleanEnum
	FundsVerified        BooleanEnum.BooleanEnum
	RealAuditing         BooleanEnum.BooleanEnum
	MobilePhone          string
	Email                string
	RealName             string
	RealNameRejectReason string
	IdCard               string
	Avatar               string
	AccountVerified      BooleanEnum.BooleanEnum
	GoogleStatus         int64
	Transactions         int
	TransactionTime      time.Time
	Level                int64
	Integration          int64
	KycStatus            int64
	MemberGradeId        int64
	GoogleState          int64
	MemberLevel          MemberLevelEnum.MemberLevelEnum
}
