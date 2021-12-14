package entity

func (this *BusinessAuthApply) SetId(id int64) {
	this.Id = id
}
func (this *BusinessAuthApply) GetId() (id int64) {
	return this.Id
}
func (this *BusinessAuthApply) SetMember(member Member) {
	this.Member = member
}
func (this *BusinessAuthApply) GetMember() (member Member) {
	return this.Member
}
func (this *BusinessAuthApply) SetCertifiedBusinessStatus(certifiedBusinessStatus CertifiedBusinessStatus) {
	this.CertifiedBusinessStatus = certifiedBusinessStatus
}
func (this *BusinessAuthApply) GetCertifiedBusinessStatus() (certifiedBusinessStatus CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *BusinessAuthApply) SetDetail(detail string) {
	this.Detail = detail
}
func (this *BusinessAuthApply) GetDetail() (detail string) {
	return this.Detail
}
func (this *BusinessAuthApply) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *BusinessAuthApply) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *BusinessAuthApply) SetAuditingTime(auditingTime Date) {
	this.AuditingTime = auditingTime
}
func (this *BusinessAuthApply) GetAuditingTime() (auditingTime Date) {
	return this.AuditingTime
}
func (this *BusinessAuthApply) SetAuthInfo(authInfo string) {
	this.AuthInfo = authInfo
}
func (this *BusinessAuthApply) GetAuthInfo() (authInfo string) {
	return this.AuthInfo
}
func (this *BusinessAuthApply) SetBusinessAuthDeposit(businessAuthDeposit BusinessAuthDeposit) {
	this.BusinessAuthDeposit = businessAuthDeposit
}
func (this *BusinessAuthApply) GetBusinessAuthDeposit() (businessAuthDeposit BusinessAuthDeposit) {
	return this.BusinessAuthDeposit
}
func (this *BusinessAuthApply) SetDepositRecordId(depositRecordId string) {
	this.DepositRecordId = depositRecordId
}
func (this *BusinessAuthApply) GetDepositRecordId() (depositRecordId string) {
	return this.DepositRecordId
}
func (this *BusinessAuthApply) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *BusinessAuthApply) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthApply) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *BusinessAuthApply) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *BusinessAuthApply) SetInfo(info JSONObject) {
	this.Info = info
}
func (this *BusinessAuthApply) GetInfo() (info JSONObject) {
	return this.Info
}

type BusinessAuthApply struct {
	Id                      int64
	Member                  Member
	CertifiedBusinessStatus CertifiedBusinessStatus
	Detail                  string
	CreateTime              Date
	AuditingTime            Date
	AuthInfo                string
	BusinessAuthDeposit     BusinessAuthDeposit
	DepositRecordId         string
	Amount                  BigDecimal
	UpdateTime              Date
	Info                    JSONObject
}
