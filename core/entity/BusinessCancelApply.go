package entity

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"time"
)

func (this *BusinessCancelApply) SetId(id int64) (result *BusinessCancelApply) {
	this.Id = id
	return this
}
func (this *BusinessCancelApply) GetId() (id int64) {
	return this.Id
}
func (this *BusinessCancelApply) SetMember(member *Member) (result *BusinessCancelApply) {
	this.Member = member
	return this
}
func (this *BusinessCancelApply) GetMember() (member *Member) {
	return this.Member
}
func (this *BusinessCancelApply) SetStatus(status CertifiedBusinessStatus.CertifiedBusinessStatus) (result *BusinessCancelApply) {
	this.Status = status
	return this
}
func (this *BusinessCancelApply) GetStatus() (status CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.Status
}
func (this *BusinessCancelApply) SetDepositRecordId(depositRecordId string) (result *BusinessCancelApply) {
	this.DepositRecordId = depositRecordId
	return this
}
func (this *BusinessCancelApply) GetDepositRecordId() (depositRecordId string) {
	return this.DepositRecordId
}
func (this *BusinessCancelApply) SetCancelApplyTime(cancelApplyTime time.Time) (result *BusinessCancelApply) {
	this.CancelApplyTime = cancelApplyTime
	return this
}
func (this *BusinessCancelApply) GetCancelApplyTime() (cancelApplyTime time.Time) {
	return this.CancelApplyTime
}
func (this *BusinessCancelApply) SetHandleTime(handleTime time.Time) (result *BusinessCancelApply) {
	this.HandleTime = handleTime
	return this
}
func (this *BusinessCancelApply) GetHandleTime() (handleTime time.Time) {
	return this.HandleTime
}
func (this *BusinessCancelApply) SetReason(reason string) (result *BusinessCancelApply) {
	this.Reason = reason
	return this
}
func (this *BusinessCancelApply) GetReason() (reason string) {
	return this.Reason
}
func (this *BusinessCancelApply) SetDetail(detail string) (result *BusinessCancelApply) {
	this.Detail = detail
	return this
}
func (this *BusinessCancelApply) GetDetail() (detail string) {
	return this.Detail
}
func (this *BusinessCancelApply) SetDepositRecord(depositRecord *DepositRecord) (result *BusinessCancelApply) {
	this.DepositRecord = depositRecord
	return this
}
func (this *BusinessCancelApply) GetDepositRecord() (depositRecord *DepositRecord) {
	return this.DepositRecord
}
func NewBusinessCancelApply() (this *BusinessCancelApply) {
	this = new(BusinessCancelApply)
	return
}

type BusinessCancelApply struct {
	Id              int64
	Member          *Member
	Status          CertifiedBusinessStatus.CertifiedBusinessStatus
	DepositRecordId string
	CancelApplyTime time.Time
	HandleTime      time.Time
	Reason          string
	Detail          string
	DepositRecord   *DepositRecord
}
