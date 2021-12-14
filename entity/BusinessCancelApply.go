package entity

func (this *BusinessCancelApply) SetId(id int64) {
	this.Id = id
}
func (this *BusinessCancelApply) GetId() (id int64) {
	return this.Id
}
func (this *BusinessCancelApply) SetMember(member Member) {
	this.Member = member
}
func (this *BusinessCancelApply) GetMember() (member Member) {
	return this.Member
}
func (this *BusinessCancelApply) SetStatus(status CertifiedBusinessStatus) {
	this.Status = status
}
func (this *BusinessCancelApply) GetStatus() (status CertifiedBusinessStatus) {
	return this.Status
}
func (this *BusinessCancelApply) SetDepositRecordId(depositRecordId string) {
	this.DepositRecordId = depositRecordId
}
func (this *BusinessCancelApply) GetDepositRecordId() (depositRecordId string) {
	return this.DepositRecordId
}
func (this *BusinessCancelApply) SetCancelApplyTime(cancelApplyTime Date) {
	this.CancelApplyTime = cancelApplyTime
}
func (this *BusinessCancelApply) GetCancelApplyTime() (cancelApplyTime Date) {
	return this.CancelApplyTime
}
func (this *BusinessCancelApply) SetHandleTime(handleTime Date) {
	this.HandleTime = handleTime
}
func (this *BusinessCancelApply) GetHandleTime() (handleTime Date) {
	return this.HandleTime
}
func (this *BusinessCancelApply) SetReason(reason string) {
	this.Reason = reason
}
func (this *BusinessCancelApply) GetReason() (reason string) {
	return this.Reason
}
func (this *BusinessCancelApply) SetDetail(detail string) {
	this.Detail = detail
}
func (this *BusinessCancelApply) GetDetail() (detail string) {
	return this.Detail
}
func (this *BusinessCancelApply) SetDepositRecord(depositRecord DepositRecord) {
	this.DepositRecord = depositRecord
}
func (this *BusinessCancelApply) GetDepositRecord() (depositRecord DepositRecord) {
	return this.DepositRecord
}
func NewBusinessCancelApply() (this *BusinessCancelApply) {
	this = new(BusinessCancelApply)
	return
}

type BusinessCancelApply struct {
	Id              int64
	Member          Member
	Status          CertifiedBusinessStatus
	DepositRecordId string
	CancelApplyTime Date
	HandleTime      Date
	Reason          string
	Detail          string
	DepositRecord   DepositRecord
}
