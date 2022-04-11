package entity

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/math"
)

func (this *BusinessAuthApply) SetId(Id int64) (result *BusinessAuthApply) {
	this.Id = Id
	return this
}
func (this *BusinessAuthApply) GetId() (Id int64) {
	return this.Id
}
func (this *BusinessAuthApply) SetMember(Member *Member) (result *BusinessAuthApply) {
	this.Member = Member
	return this
}
func (this *BusinessAuthApply) GetMember() (Member *Member) {
	return this.Member
}
func (this *BusinessAuthApply) SetCertifiedBusinessStatus(CertifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus) (result *BusinessAuthApply) {
	this.CertifiedBusinessStatus = CertifiedBusinessStatus
	return this
}
func (this *BusinessAuthApply) GetCertifiedBusinessStatus() (CertifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.CertifiedBusinessStatus
}
func (this *BusinessAuthApply) SetDetail(Detail string) (result *BusinessAuthApply) {
	this.Detail = Detail
	return this
}
func (this *BusinessAuthApply) GetDetail() (Detail string) {
	return this.Detail
}
func (this *BusinessAuthApply) SetCreateTime(CreateTime xtime.Xtime) (result *BusinessAuthApply) {
	this.CreateTime = CreateTime
	return this
}
func (this *BusinessAuthApply) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *BusinessAuthApply) SetAuditingTime(AuditingTime xtime.Xtime) (result *BusinessAuthApply) {
	this.AuditingTime = AuditingTime
	return this
}
func (this *BusinessAuthApply) GetAuditingTime() (AuditingTime xtime.Xtime) {
	return this.AuditingTime
}
func (this *BusinessAuthApply) SetAuthInfo(AuthInfo string) (result *BusinessAuthApply) {
	this.AuthInfo = AuthInfo
	return this
}
func (this *BusinessAuthApply) GetAuthInfo() (AuthInfo string) {
	return this.AuthInfo
}
func (this *BusinessAuthApply) SetBusinessAuthDeposit(BusinessAuthDeposit *BusinessAuthDeposit) (result *BusinessAuthApply) {
	this.BusinessAuthDeposit = BusinessAuthDeposit
	return this
}
func (this *BusinessAuthApply) GetBusinessAuthDeposit() (BusinessAuthDeposit *BusinessAuthDeposit) {
	return this.BusinessAuthDeposit
}
func (this *BusinessAuthApply) SetDepositRecordId(DepositRecordId string) (result *BusinessAuthApply) {
	this.DepositRecordId = DepositRecordId
	return this
}
func (this *BusinessAuthApply) GetDepositRecordId() (DepositRecordId string) {
	return this.DepositRecordId
}
func (this *BusinessAuthApply) SetAmount(Amount math.BigDecimal) (result *BusinessAuthApply) {
	this.Amount = Amount
	return this
}
func (this *BusinessAuthApply) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthApply) SetUpdateTime(UpdateTime xtime.Xtime) (result *BusinessAuthApply) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *BusinessAuthApply) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *BusinessAuthApply) SetInfo(Info *fastjson.JSONObject) (result *BusinessAuthApply) {
	this.Info = Info
	return this
}
func (this *BusinessAuthApply) GetInfo() (Info *fastjson.JSONObject) {
	return this.Info
}
func NewBusinessAuthApply(id int64, member *Member, certifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus, detail string, createTime xtime.Xtime, auditingTime xtime.Xtime, authInfo string, businessAuthDeposit *BusinessAuthDeposit, depositRecordId string, amount math.BigDecimal, updateTime xtime.Xtime, info *fastjson.JSONObject) (ret *BusinessAuthApply) {
	ret = new(BusinessAuthApply)
	ret.Id = id
	ret.Member = member
	ret.CertifiedBusinessStatus = certifiedBusinessStatus
	ret.Detail = detail
	ret.CreateTime = createTime
	ret.AuditingTime = auditingTime
	ret.AuthInfo = authInfo
	ret.BusinessAuthDeposit = businessAuthDeposit
	ret.DepositRecordId = depositRecordId
	ret.Amount = amount
	ret.UpdateTime = updateTime
	ret.Info = info
	return
}

type BusinessAuthApply struct {
	Id                      int64                                           `gorm:"column:id" json:"id"`
	Member                  *Member                                         `gorm:"column:member" json:"member"`
	CertifiedBusinessStatus CertifiedBusinessStatus.CertifiedBusinessStatus `gorm:"column:certified_business_status" json:"certifiedBusinessStatus"`
	Detail                  string                                          `gorm:"column:detail" json:"detail"`
	CreateTime              xtime.Xtime                                     `gorm:"column:create_time" json:"createTime"`
	AuditingTime            xtime.Xtime                                     `gorm:"column:auditing_time" json:"auditingTime"`
	AuthInfo                string                                          `gorm:"column:auth_info" json:"authInfo"`
	BusinessAuthDeposit     *BusinessAuthDeposit                            `gorm:"column:business_auth_deposit" json:"businessAuthDeposit"`
	DepositRecordId         string                                          `gorm:"column:deposit_record_id" json:"depositRecordId"`
	Amount                  math.BigDecimal                                 `gorm:"column:amount" json:"amount"`
	UpdateTime              xtime.Xtime                                     `gorm:"column:update_time" json:"updateTime"`
	Info                    *fastjson.JSONObject                            `gorm:"column:info" json:"info"`
}
