package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type AdminAccessLogDao interface {
	Save(m *entity.AdminAccessLog) (result *entity.AdminAccessLog, err error)
	FindById(id int64) (result *entity.AdminAccessLog, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.AdminAccessLog], err error)
}
type adminAccessLogDao struct {
	*db.DB
}

func NewAdminAccessLogDao(db *db.DB) (dao AdminAccessLogDao) {
	dao = &adminAccessLogDao{db}
	return
}
func (this *adminAccessLogDao) Save(m *entity.AdminAccessLog) (result *entity.AdminAccessLog, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *adminAccessLogDao) FindById(id int64) (result *entity.AdminAccessLog, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *adminAccessLogDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.AdminAccessLog{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *adminAccessLogDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.AdminAccessLog], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
