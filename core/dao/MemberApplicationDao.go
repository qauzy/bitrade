package dao

import (
	"bitrade/core/constant/AuditStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type MemberApplicationDao interface {
	FindMemberApplicationByMemberAndAuditStatusOrderByIdDesc(var1 entity.Member, var2 AuditStatus.AuditStatus) (result []entity.MemberApplication, err error)
	CountAllByAuditStatus(auditStatus AuditStatus.AuditStatus) (result int64, err error)
	FindSuccessMemberApplicationsByIdCard(idCard string, sta1 AuditStatus.AuditStatus, sta2 AuditStatus.AuditStatus) (result []entity.MemberApplication, err error)
	FindMemberApplicationByKycStatusInAndMember(kycStatus []int64, member entity.Member) (result entity.MemberApplication, err error)
	Save(m *entity.MemberApplication) (result *entity.MemberApplication, err error)
	FindById(id int64) (result *entity.MemberApplication, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.MemberApplication, err error)
}
type memberApplicationDao struct {
	*db.DB
}

func NewMemberApplicationDao(db *db.DB) (dao MemberApplicationDao) {
	dao = &memberApplicationDao{db}
	return
}
func (this *memberApplicationDao) FindMemberApplicationByMemberAndAuditStatusOrderByIdDesc(var1 entity.Member, var2 AuditStatus.AuditStatus) (result []entity.MemberApplication, err error) {
	err = this.DBRead().Where("id_desc = ?", var1).First(&result).Error
	return
}
func (this *memberApplicationDao) CountAllByAuditStatus(auditStatus AuditStatus.AuditStatus) (result int64, err error) {
	return
}
func (this *memberApplicationDao) FindSuccessMemberApplicationsByIdCard(idCard string, sta1 AuditStatus.AuditStatus, sta2 AuditStatus.AuditStatus) (result []entity.MemberApplication, err error) {
	err = this.DBRead().Where("id_card = ?", idCard).First(&result).Error
	return
}
func (this *memberApplicationDao) FindMemberApplicationByKycStatusInAndMember(kycStatus []int64, member entity.Member) (result entity.MemberApplication, err error) {
	err = this.DBRead().Where("kyc_status_in = ?", kycStatus).Where("member = ?", member).First(&result).Error
	return
}
func (this *memberApplicationDao) Save(m *entity.MemberApplication) (result *entity.MemberApplication, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberApplicationDao) FindById(id int64) (result *entity.MemberApplication, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberApplicationDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberApplication{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberApplicationDao) FindAll(qp *types.QueryParam) (result []*entity.MemberApplication, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
