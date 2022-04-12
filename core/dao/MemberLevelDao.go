package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type MemberLevelDao interface {
	FindOneByIsDefault(isDefault bool) (result *entity.MemberLevel, err error)
	UpdateDefault() (result int, err error)
	Save(m *entity.MemberLevel) (result *entity.MemberLevel, err error)
	FindById(id int64) (result *entity.MemberLevel, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberLevel], err error)
}
type memberLevelDao struct {
	*db.DB
}

func NewMemberLevelDao(db *db.DB) (dao MemberLevelDao) {
	dao = &memberLevelDao{db}
	return
}
func (this *memberLevelDao) FindOneByIsDefault(isDefault bool) (result *entity.MemberLevel, err error) {
	err = this.DBRead().Where("is_default = ?", isDefault).First(&result).Error
	return
}
func (this *memberLevelDao) UpdateDefault() (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberLevel set isDefault = false  where isDefault = true ")
	err = eng.Error
	return
}
func (this *memberLevelDao) Save(m *entity.MemberLevel) (result *entity.MemberLevel, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberLevelDao) FindById(id int64) (result *entity.MemberLevel, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberLevelDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberLevel{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberLevelDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberLevel], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
