package entity

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/constant/MemberLevelEnum"
)

func (this *CertifiedBusinessInfo) SetMemberLevel(MemberLevel *MemberLevelEnum.MemberLevelEnum) (result *CertifiedBusinessInfo) {
	this.MemberLevel = MemberLevel
	return this
}
func (this *CertifiedBusinessInfo) GetMemberLevel() (MemberLevel *MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *CertifiedBusinessInfo) SetCertifiedBusinessStatus(CertifiedBusinessStatus *CertifiedBusinessStatus.CertifiedBusinessStatus) (result *CertifiedBusinessInfo) {
	this.CertifiedBusinessStatus = CertifiedBusinessStatus
	return this
}
func (this *CertifiedBusinessInfo) GetCertifiedBusinessStatus() (CertifiedBusinessStatus *CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *CertifiedBusinessInfo) SetEmail(Email string) (result *CertifiedBusinessInfo) {
	this.Email = Email
	return this
}
func (this *CertifiedBusinessInfo) GetEmail() (Email string) {
	return this.Email
}
func (this *CertifiedBusinessInfo) SetDetail(Detail string) (result *CertifiedBusinessInfo) {
	this.Detail = Detail
	return this
}
func (this *CertifiedBusinessInfo) GetDetail() (Detail string) {
	return this.Detail
}
func (this *CertifiedBusinessInfo) SetReason(Reason string) (result *CertifiedBusinessInfo) {
	this.Reason = Reason
	return this
}
func (this *CertifiedBusinessInfo) GetReason() (Reason string) {
	return this.Reason
}

type CertifiedBusinessInfo struct {
	MemberLevel             *MemberLevelEnum.MemberLevelEnum                 `gorm:"column:member_level" json:"memberLevel"`
	CertifiedBusinessStatus *CertifiedBusinessStatus.CertifiedBusinessStatus `gorm:"column:certified_business_status" json:"certifiedBusinessStatus"`
	Email                   string                                           `gorm:"column:email" json:"email"`
	Detail                  string                                           `gorm:"column:detail" json:"detail"`
	Reason                  string                                           `gorm:"column:reason" json:"reason"`
}
