package dao

import (
	"bitrade/core/constant/AppealStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type AppealDao interface {
	GetBusinessAppealInitiatorIdStatistics(memberId int64) (result int64, err error)
	GetBusinessAppealAssociateIdStatistics(memberId int64) (result int64, err error)
	CountAllByStatus(status AppealStatus.AppealStatus) (result int64, err error)
	Save(m *entity.Appeal) (result *entity.Appeal, err error)
	FindById(id int64) (result *entity.Appeal, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Appeal, err error)
}
type appealDao struct {
	*db.DB
}

func NewAppealDao(db *db.DB) (dao AppealDao) {
	dao = &appealDao{db}
	return
}
func (this *appealDao) GetBusinessAppealInitiatorIdStatistics(memberId int64) (result int64, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select count(a.id) as complainantNum from Appeal a where a.initiatorId = ?", memberId)
	err = eng.Error
	return
}
func (this *appealDao) GetBusinessAppealAssociateIdStatistics(memberId int64) (result int64, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select count(a.id) as defendantNum from Appeal a where a.associateId = ?", memberId)
	err = eng.Error
	return
}
func (this *appealDao) CountAllByStatus(status AppealStatus.AppealStatus) (result int64, err error) {
	return
}
func (this *appealDao) Save(m *entity.Appeal) (result *entity.Appeal, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *appealDao) FindById(id int64) (result *entity.Appeal, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *appealDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Appeal{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *appealDao) FindAll(qp *types.QueryParam) (result []*entity.Appeal, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
