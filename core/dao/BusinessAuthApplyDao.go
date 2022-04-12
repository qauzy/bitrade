package dao

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type BusinessAuthApplyDao interface {
	FindByMemberOrderByIdDesc(member *entity.Member) (result arraylist.List[entity.BusinessAuthApply], err error)
	FindByMemberAndCertifiedBusinessStatusOrderByIdDesc(member *entity.Member, certifiedBusinessStatus *CertifiedBusinessStatus.CertifiedBusinessStatus) (result arraylist.List[entity.BusinessAuthApply], err error)
	CountAllByCertifiedBusinessStatus(status *CertifiedBusinessStatus.CertifiedBusinessStatus) (result int64, err error)
	Save(m *entity.BusinessAuthApply) (result *entity.BusinessAuthApply, err error)
	FindById(id int64) (result *entity.BusinessAuthApply, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.BusinessAuthApply], err error)
}
type businessAuthApplyDao struct {
	*db.DB
}

func NewBusinessAuthApplyDao(db *db.DB) (dao BusinessAuthApplyDao) {
	dao = &businessAuthApplyDao{db}
	return
}
func (this *businessAuthApplyDao) FindByMemberOrderByIdDesc(member *entity.Member) (result arraylist.List[entity.BusinessAuthApply], err error) {
	err = this.DBRead().Where("id_desc = ?", member).First(&result).Error
	return
}
func (this *businessAuthApplyDao) FindByMemberAndCertifiedBusinessStatusOrderByIdDesc(member *entity.Member, certifiedBusinessStatus *CertifiedBusinessStatus.CertifiedBusinessStatus) (result arraylist.List[entity.BusinessAuthApply], err error) {
	err = this.DBRead().Where("id_desc = ?", member).First(&result).Error
	return
}
func (this *businessAuthApplyDao) CountAllByCertifiedBusinessStatus(status *CertifiedBusinessStatus.CertifiedBusinessStatus) (result int64, err error) {
	return
}
func (this *businessAuthApplyDao) Save(m *entity.BusinessAuthApply) (result *entity.BusinessAuthApply, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *businessAuthApplyDao) FindById(id int64) (result *entity.BusinessAuthApply, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *businessAuthApplyDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.BusinessAuthApply{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *businessAuthApplyDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.BusinessAuthApply], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
