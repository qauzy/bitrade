package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type DepartmentDao interface {
	Save(m *entity.Department) (result *entity.Department, err error)
	FindById(id int64) (result *entity.Department, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Department], err error)
}
type departmentDao struct {
	*db.DB
}

func NewDepartmentDao(db *db.DB) (dao DepartmentDao) {
	dao = &departmentDao{db}
	return
}
func (this *departmentDao) Save(m *entity.Department) (result *entity.Department, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *departmentDao) FindById(id int64) (result *entity.Department, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *departmentDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Department{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *departmentDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Department], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
