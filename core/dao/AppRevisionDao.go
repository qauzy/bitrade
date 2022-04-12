package dao

import (
	"bitrade/core/constant/Platform"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type AppRevisionDao interface {
	FindAppRevisionByPlatformOrderByIdDesc(platform *Platform.Platform) (result *entity.AppRevision, err error)
	Save(m *entity.AppRevision) (result *entity.AppRevision, err error)
	FindById(id int64) (result *entity.AppRevision, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.AppRevision], err error)
}
type appRevisionDao struct {
	*db.DB
}

func NewAppRevisionDao(db *db.DB) (dao AppRevisionDao) {
	dao = &appRevisionDao{db}
	return
}
func (this *appRevisionDao) FindAppRevisionByPlatformOrderByIdDesc(platform *Platform.Platform) (result *entity.AppRevision, err error) {
	err = this.DBRead().Where("id_desc = ?", platform).First(&result).Error
	return
}
func (this *appRevisionDao) Save(m *entity.AppRevision) (result *entity.AppRevision, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *appRevisionDao) FindById(id int64) (result *entity.AppRevision, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *appRevisionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.AppRevision{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *appRevisionDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.AppRevision], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
