package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type ReleaseBalanceDao interface {
	Save(m *entity.ReleaseBalance) (result *entity.ReleaseBalance, err error)
	FindById(id int64) (result *entity.ReleaseBalance, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.ReleaseBalance], err error)
}
type releaseBalanceDao struct {
	*db.DB
}

func NewReleaseBalanceDao(db *db.DB) (dao ReleaseBalanceDao) {
	dao = &releaseBalanceDao{db}
	return
}
func (this *releaseBalanceDao) Save(m *entity.ReleaseBalance) (result *entity.ReleaseBalance, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *releaseBalanceDao) FindById(id int64) (result *entity.ReleaseBalance, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *releaseBalanceDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.ReleaseBalance{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *releaseBalanceDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.ReleaseBalance], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
