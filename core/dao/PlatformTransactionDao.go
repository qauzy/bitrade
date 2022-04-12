package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type PlatformTransactionDao interface {
	Save(m *entity.PlatformTransaction) (result *entity.PlatformTransaction, err error)
	FindById(id int64) (result *entity.PlatformTransaction, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.PlatformTransaction], err error)
}
type platformTransactionDao struct {
	*db.DB
}

func NewPlatformTransactionDao(db *db.DB) (dao PlatformTransactionDao) {
	dao = &platformTransactionDao{db}
	return
}
func (this *platformTransactionDao) Save(m *entity.PlatformTransaction) (result *entity.PlatformTransaction, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *platformTransactionDao) FindById(id int64) (result *entity.PlatformTransaction, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *platformTransactionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.PlatformTransaction{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *platformTransactionDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.PlatformTransaction], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
