package dao

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type BusinessCancelApplyDao interface {
	FindByMemberAndStatusOrderByIdDesc(member entity.Member, status CertifiedBusinessStatus.CertifiedBusinessStatus) (result []entity.BusinessCancelApply, err error)
	FindByMemberOrderByIdDesc(member entity.Member) (result []entity.BusinessCancelApply, err error)
	CountAllByStatus(status CertifiedBusinessStatus.CertifiedBusinessStatus) (result int64, err error)
	Save(m *entity.BusinessCancelApply) (result *entity.BusinessCancelApply, err error)
	FindById(id int64) (result *entity.BusinessCancelApply, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.BusinessCancelApply, err error)
}
type businessCancelApplyDao struct {
	*db.DB
}

func NewBusinessCancelApplyDao(db *db.DB) (dao BusinessCancelApplyDao) {
	dao = &businessCancelApplyDao{db}
	return
}
func (this *businessCancelApplyDao) FindByMemberAndStatusOrderByIdDesc(member entity.Member, status CertifiedBusinessStatus.CertifiedBusinessStatus) (result []entity.BusinessCancelApply, err error) {
	err = this.DBRead().Where("id_desc = ?", member).First(&result).Error
	return
}
func (this *businessCancelApplyDao) FindByMemberOrderByIdDesc(member entity.Member) (result []entity.BusinessCancelApply, err error) {
	err = this.DBRead().Where("id_desc = ?", member).First(&result).Error
	return
}
func (this *businessCancelApplyDao) CountAllByStatus(status CertifiedBusinessStatus.CertifiedBusinessStatus) (result int64, err error) {
	return
}
func (this *businessCancelApplyDao) Save(m *entity.BusinessCancelApply) (result *entity.BusinessCancelApply, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *businessCancelApplyDao) FindById(id int64) (result *entity.BusinessCancelApply, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *businessCancelApplyDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.BusinessCancelApply{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *businessCancelApplyDao) FindAll(qp *types.QueryParam) (result []*entity.BusinessCancelApply, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
