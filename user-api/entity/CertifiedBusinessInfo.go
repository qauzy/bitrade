package entity

import "bitrade/core/constant"

func (this *CertifiedBusinessInfo) SetMemberLevel(memberLevel constant.MemberLevelEnum) (result *CertifiedBusinessInfo) {
	this.MemberLevel = memberLevel
	return this
}
func (this *CertifiedBusinessInfo) GetMemberLevel() (memberLevel constant.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *CertifiedBusinessInfo) SetCertifiedBusinessStatus(certifiedBusinessStatus constant.CertifiedBusinessStatus) (result *CertifiedBusinessInfo) {
	this.CertifiedBusinessStatus = certifiedBusinessStatus
	return this
}
func (this *CertifiedBusinessInfo) GetCertifiedBusinessStatus() (certifiedBusinessStatus constant.CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *CertifiedBusinessInfo) SetEmail(email string) (result *CertifiedBusinessInfo) {
	this.Email = email
	return this
}
func (this *CertifiedBusinessInfo) GetEmail() (email string) {
	return this.Email
}
func (this *CertifiedBusinessInfo) SetDetail(detail string) (result *CertifiedBusinessInfo) {
	this.Detail = detail
	return this
}
func (this *CertifiedBusinessInfo) GetDetail() (detail string) {
	return this.Detail
}
func (this *CertifiedBusinessInfo) SetReason(reason string) (result *CertifiedBusinessInfo) {
	this.Reason = reason
	return this
}
func (this *CertifiedBusinessInfo) GetReason() (reason string) {
	return this.Reason
}

type CertifiedBusinessInfo struct {
	MemberLevel             constant.MemberLevelEnum
	CertifiedBusinessStatus constant.CertifiedBusinessStatus
	Email                   string
	Detail                  string
	Reason                  string
}
