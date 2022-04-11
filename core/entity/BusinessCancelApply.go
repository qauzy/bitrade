package entity

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"github.com/qauzy/chocolate/xtime"
)

func (this *BusinessCancelApply) SetId(Id int64) (result *BusinessCancelApply) {
	this.Id = Id
	return this
}
func (this *BusinessCancelApply) GetId() (Id int64) {
	return this.Id
}
func (this *BusinessCancelApply) SetMember(Member *Member) (result *BusinessCancelApply) {
	this.Member = Member
	return this
}
func (this *BusinessCancelApply) GetMember() (Member *Member) {
	return this.Member
}
func (this *BusinessCancelApply) SetStatus(Status CertifiedBusinessStatus.CertifiedBusinessStatus) (result *BusinessCancelApply) {
	this.Status = Status
	return this
}
func (this *BusinessCancelApply) GetStatus() (Status CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.Status
}
func (this *BusinessCancelApply) SetDepositRecordId(DepositRecordId string) (result *BusinessCancelApply) {
	this.DepositRecordId = DepositRecordId
	return this
}
func (this *BusinessCancelApply) GetDepositRecordId() (DepositRecordId string) {
	return this.DepositRecordId
}
func (this *BusinessCancelApply) SetCancelApplyTime(CancelApplyTime xtime.Xtime) (result *BusinessCancelApply) {
	this.CancelApplyTime = CancelApplyTime
	return this
}
func (this *BusinessCancelApply) GetCancelApplyTime() (CancelApplyTime xtime.Xtime) {
	return this.CancelApplyTime
}
func (this *BusinessCancelApply) SetHandleTime(HandleTime xtime.Xtime) (result *BusinessCancelApply) {
	this.HandleTime = HandleTime
	return this
}
func (this *BusinessCancelApply) GetHandleTime() (HandleTime xtime.Xtime) {
	return this.HandleTime
}
func (this *BusinessCancelApply) SetReason(Reason string) (result *BusinessCancelApply) {
	this.Reason = Reason
	return this
}
func (this *BusinessCancelApply) GetReason() (Reason string) {
	return this.Reason
}
func (this *BusinessCancelApply) SetDetail(Detail string) (result *BusinessCancelApply) {
	this.Detail = Detail
	return this
}
func (this *BusinessCancelApply) GetDetail() (Detail string) {
	return this.Detail
}
func (this *BusinessCancelApply) SetDepositRecord(DepositRecord *DepositRecord) (result *BusinessCancelApply) {
	this.DepositRecord = DepositRecord
	return this
}
func (this *BusinessCancelApply) GetDepositRecord() (DepositRecord *DepositRecord) {
	return this.DepositRecord
}
func NewBusinessCancelApply() (this *BusinessCancelApply) {
	this = new(BusinessCancelApply)
	return
}
func NewBusinessCancelApplyEx(id int64, member *Member, status CertifiedBusinessStatus.CertifiedBusinessStatus, depositRecordId string, cancelApplyTime xtime.Xtime, handleTime xtime.Xtime, reason string, detail string, depositRecord *DepositRecord) (ret *BusinessCancelApply) {
	ret = new(BusinessCancelApply)
	ret.Id = id
	ret.Member = member
	ret.Status = status
	ret.DepositRecordId = depositRecordId
	ret.CancelApplyTime = cancelApplyTime
	ret.HandleTime = handleTime
	ret.Reason = reason
	ret.Detail = detail
	ret.DepositRecord = depositRecord
	return
}

type BusinessCancelApply struct {
	Id              int64                                           `gorm:"column:id" json:"id"`
	Member          *Member                                         `gorm:"column:member" json:"member"`
	Status          CertifiedBusinessStatus.CertifiedBusinessStatus `gorm:"column:status" json:"status"`
	DepositRecordId string                                          `gorm:"column:deposit_record_id" json:"depositRecordId"`
	CancelApplyTime xtime.Xtime                                     `gorm:"column:cancel_apply_time" json:"cancelApplyTime"`
	HandleTime      xtime.Xtime                                     `gorm:"column:handle_time" json:"handleTime"`
	Reason          string                                          `gorm:"column:reason" json:"reason"`
	Detail          string                                          `gorm:"column:detail" json:"detail"`
	DepositRecord   *DepositRecord                                  `gorm:"column:deposit_record" json:"depositRecord"`
}
