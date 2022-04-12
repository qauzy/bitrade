package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type MemberApplicationConfigDao interface {
	Save(m *entity.MemberApplicationConfig) (result *entity.MemberApplicationConfig, err error)
	FindById(id int64) (result *entity.MemberApplicationConfig, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberApplicationConfig], err error)
}
type memberApplicationConfigDao struct {
	*db.DB
}

func NewMemberApplicationConfigDao(db *db.DB) (dao MemberApplicationConfigDao) {
	dao = &memberApplicationConfigDao{db}
	return
}
func (this *memberApplicationConfigDao) Save(m *entity.MemberApplicationConfig) (result *entity.MemberApplicationConfig, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberApplicationConfigDao) FindById(id int64) (result *entity.MemberApplicationConfig, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberApplicationConfigDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberApplicationConfig{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberApplicationConfigDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberApplicationConfig], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
