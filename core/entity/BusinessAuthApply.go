package entity

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
	"github.com/valyala/fastjson"
	"time"
)

func (this *BusinessAuthApply) SetId(id int64) (result *BusinessAuthApply) {
	this.Id = id
	return this
}
func (this *BusinessAuthApply) GetId() (id int64) {
	return this.Id
}
func (this *BusinessAuthApply) SetMember(member Member) (result *BusinessAuthApply) {
	this.Member = member
	return this
}
func (this *BusinessAuthApply) GetMember() (member Member) {
	return this.Member
}
func (this *BusinessAuthApply) SetCertifiedBusinessStatus(certifiedBusinessStatus constant.CertifiedBusinessStatus) (result *BusinessAuthApply) {
	this.CertifiedBusinessStatus = certifiedBusinessStatus
	return this
}
func (this *BusinessAuthApply) GetCertifiedBusinessStatus() (certifiedBusinessStatus constant.CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *BusinessAuthApply) SetDetail(detail string) (result *BusinessAuthApply) {
	this.Detail = detail
	return this
}
func (this *BusinessAuthApply) GetDetail() (detail string) {
	return this.Detail
}
func (this *BusinessAuthApply) SetCreateTime(createTime time.Time) (result *BusinessAuthApply) {
	this.CreateTime = createTime
	return this
}
func (this *BusinessAuthApply) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *BusinessAuthApply) SetAuditingTime(auditingTime time.Time) (result *BusinessAuthApply) {
	this.AuditingTime = auditingTime
	return this
}
func (this *BusinessAuthApply) GetAuditingTime() (auditingTime time.Time) {
	return this.AuditingTime
}
func (this *BusinessAuthApply) SetAuthInfo(authInfo string) (result *BusinessAuthApply) {
	this.AuthInfo = authInfo
	return this
}
func (this *BusinessAuthApply) GetAuthInfo() (authInfo string) {
	return this.AuthInfo
}
func (this *BusinessAuthApply) SetBusinessAuthDeposit(businessAuthDeposit BusinessAuthDeposit) (result *BusinessAuthApply) {
	this.BusinessAuthDeposit = businessAuthDeposit
	return this
}
func (this *BusinessAuthApply) GetBusinessAuthDeposit() (businessAuthDeposit BusinessAuthDeposit) {
	return this.BusinessAuthDeposit
}
func (this *BusinessAuthApply) SetDepositRecordId(depositRecordId string) (result *BusinessAuthApply) {
	this.DepositRecordId = depositRecordId
	return this
}
func (this *BusinessAuthApply) GetDepositRecordId() (depositRecordId string) {
	return this.DepositRecordId
}
func (this *BusinessAuthApply) SetAmount(amount math.BigDecimal) (result *BusinessAuthApply) {
	this.Amount = amount
	return this
}
func (this *BusinessAuthApply) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthApply) SetUpdateTime(updateTime time.Time) (result *BusinessAuthApply) {
	this.UpdateTime = updateTime
	return this
}
func (this *BusinessAuthApply) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *BusinessAuthApply) SetInfo(info fastjson.Object) (result *BusinessAuthApply) {
	this.Info = info
	return this
}
func (this *BusinessAuthApply) GetInfo() (info fastjson.Object) {
	return this.Info
}

type BusinessAuthApply struct {
	Id                      int64
	Member                  Member
	CertifiedBusinessStatus constant.CertifiedBusinessStatus
	Detail                  string
	CreateTime              time.Time
	AuditingTime            time.Time
	AuthInfo                string
	BusinessAuthDeposit     BusinessAuthDeposit
	DepositRecordId         string
	Amount                  math.BigDecimal
	UpdateTime              time.Time
	Info                    fastjson.Object
}
