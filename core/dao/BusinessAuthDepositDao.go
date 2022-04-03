package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type BusinessAuthDepositDao interface {
	FindAllByStatus(status *CommonStatus.CommonStatus) (result arraylist.List[entity.BusinessAuthDeposit], err error)
	Save(m *entity.BusinessAuthDeposit) (result *entity.BusinessAuthDeposit, err error)
	FindById(id int64) (result *entity.BusinessAuthDeposit, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.BusinessAuthDeposit], err error)
}
type businessAuthDepositDao struct {
	*db.DB
}

func NewBusinessAuthDepositDao(db *db.DB) (dao BusinessAuthDepositDao) {
	dao = &businessAuthDepositDao{db}
	return
}
func (this *businessAuthDepositDao) FindAllByStatus(status *CommonStatus.CommonStatus) (result arraylist.List[entity.BusinessAuthDeposit], err error) {
	err = this.DBRead().Where("status = ?", status).Find(&result).Error
	return
}
func (this *businessAuthDepositDao) Save(m *entity.BusinessAuthDeposit) (result *entity.BusinessAuthDeposit, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *businessAuthDepositDao) FindById(id int64) (result *entity.BusinessAuthDeposit, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *businessAuthDepositDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.BusinessAuthDeposit{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *businessAuthDepositDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.BusinessAuthDeposit], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
