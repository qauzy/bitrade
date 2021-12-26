package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type SysPermissionDao interface {
	DeletePermission(permissionId int64) (result int, err error)
	FindSysPermissionByName(name string) (result entity.SysPermission, err error)
	Save(m *entity.SysPermission) (result *entity.SysPermission, err error)
	FindById(id int64) (result *entity.SysPermission, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.SysPermission, err error)
}
type sysPermissionDao struct {
	*db.DB
}

func NewSysPermissionDao(db *db.DB) (dao SysPermissionDao) {
	dao = &sysPermissionDao{db}
	return
}
func (this *sysPermissionDao) DeletePermission(permissionId int64) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("delete from admin_role_permission where rule_id = ?", permissionId)
	err = eng.Error
	return
}
func (this *sysPermissionDao) FindSysPermissionByName(name string) (result entity.SysPermission, err error) {
	err = this.DBRead().Where("name = ?", name).First(&result).Error
	return
}
func (this *sysPermissionDao) Save(m *entity.SysPermission) (result *entity.SysPermission, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *sysPermissionDao) FindById(id int64) (result *entity.SysPermission, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *sysPermissionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.SysPermission{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *sysPermissionDao) FindAll(qp *types.QueryParam) (result []*entity.SysPermission, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
