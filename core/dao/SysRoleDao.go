package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type SysRoleDao interface {
	UpdateSysRole(description string, role string, id int64) (result int, err error)
	FindAllSysRole() (result []entity.SysRole, err error)
	Save(m *entity.SysRole) (result *entity.SysRole, err error)
	FindById(id int64) (result *entity.SysRole, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.SysRole, err error)
}
type sysRoleDao struct {
	*db.DB
}

func NewSysRoleDao(db *db.DB) (dao SysRoleDao) {
	dao = &sysRoleDao{db}
	return
}
func (this *sysRoleDao) UpdateSysRole(description string, role string, id int64) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update SysRole s set s.description=?,s.role=? where s.id=?", description, role, id)
	err = eng.Error
	return
}
func (this *sysRoleDao) FindAllSysRole() (result []entity.SysRole, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT new SysRole(s.id,s.role,s.description) FROM SysRole s")
	err = eng.Error
	return
}
func (this *sysRoleDao) Save(m *entity.SysRole) (result *entity.SysRole, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *sysRoleDao) FindById(id int64) (result *entity.SysRole, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *sysRoleDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.SysRole{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *sysRoleDao) FindAll(qp *types.QueryParam) (result []*entity.SysRole, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
