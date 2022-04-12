package dao

import (
	"bitrade/core/constant/AppealStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type AppealDao interface {
	GetBusinessAppealInitiatorIdStatistics(memberId int64) (result int64, err error)
	GetBusinessAppealAssociateIdStatistics(memberId int64) (result int64, err error)
	CountAllByStatus(status *AppealStatus.AppealStatus) (result int64, err error)
	Save(m *entity.Appeal) (result *entity.Appeal, err error)
	FindById(id int64) (result *entity.Appeal, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Appeal], err error)
}
type appealDao struct {
	*db.DB
}

func NewAppealDao(db *db.DB) (dao AppealDao) {
	dao = &appealDao{db}
	return
}
func (this *appealDao) GetBusinessAppealInitiatorIdStatistics(memberId int64) (result int64, err error) {
	eng := this.DBWrite().Table("Appeal as a").Select("count(a.id) as complainantNum").Where("a.initiatorId = ?", memberId).Find(&result)
	err = eng.Error
	return
}
func (this *appealDao) GetBusinessAppealAssociateIdStatistics(memberId int64) (result int64, err error) {
	eng := this.DBWrite().Table("Appeal as a").Select("count(a.id) as defendantNum").Where("a.associateId = ?", memberId).Find(&result)
	err = eng.Error
	return
}
func (this *appealDao) CountAllByStatus(status *AppealStatus.AppealStatus) (result int64, err error) {
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
func (this *appealDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Appeal], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
