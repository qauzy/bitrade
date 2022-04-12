package dao

import (
	"bitrade/core/constant/SignStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type SignDao interface {
	FindByStatus(status *SignStatus.SignStatus) (result *entity.Sign, err error)
	Save(m *entity.Sign) (result *entity.Sign, err error)
	FindById(id int64) (result *entity.Sign, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Sign], err error)
}
type signDao struct {
	*db.DB
}

func NewSignDao(db *db.DB) (dao SignDao) {
	dao = &signDao{db}
	return
}
func (this *signDao) FindByStatus(status *SignStatus.SignStatus) (result *entity.Sign, err error) {
	err = this.DBRead().Where("status = ?", status).First(&result).Error
	return
}
func (this *signDao) Save(m *entity.Sign) (result *entity.Sign, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *signDao) FindById(id int64) (result *entity.Sign, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *signDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Sign{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *signDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Sign], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
