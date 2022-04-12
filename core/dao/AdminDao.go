package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"time"
)

type AdminDao interface {
	FindAdminByUsernameAndPassword(username string, password string) (result *entity.Admin, err error)
	UpdateAdminLastTimeAndIp(date time.Time, ip string, memberId int64) (result int, err error)
	DeleteBatch(roleId int64) (result int, err error)
	FindAllByDepartment(department *entity.Department) (result arraylist.List[entity.Admin], err error)
	FindAllByRoleId(id int64) (result arraylist.List[entity.Admin], err error)
	FindByUsername(username string) (result *entity.Admin, err error)
	FindByMobilePhone(mobilePhone string) (result arraylist.List[entity.Admin], err error)
	Save(m *entity.Admin) (result *entity.Admin, err error)
	FindById(id int64) (result *entity.Admin, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Admin], err error)
}
type adminDao struct {
	*db.DB
}

func NewAdminDao(db *db.DB) (dao AdminDao) {
	dao = &adminDao{db}
	return
}
func (this *adminDao) FindAdminByUsernameAndPassword(username string, password string) (result *entity.Admin, err error) {
	err = this.DBRead().Where("username = ?", username).Where("password = ?", password).First(&result).Error
	return
}
func (this *adminDao) UpdateAdminLastTimeAndIp(date time.Time, ip string, memberId int64) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update Admin a set a.lastLoginTime=?,a.lastLoginIp=? where a.id=?", date, ip, memberId)
	err = eng.Error
	return
}
func (this *adminDao) DeleteBatch(roleId int64) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("delete from Admin a where a.roleId = ?", roleId)
	err = eng.Error
	return
}
func (this *adminDao) FindAllByDepartment(department *entity.Department) (result arraylist.List[entity.Admin], err error) {
	err = this.DBRead().Where("department = ?", department).Find(&result).Error
	return
}
func (this *adminDao) FindAllByRoleId(id int64) (result arraylist.List[entity.Admin], err error) {
	err = this.DBRead().Where("role_id = ?", id).Find(&result).Error
	return
}
func (this *adminDao) FindByUsername(username string) (result *entity.Admin, err error) {
	err = this.DBRead().Where("username = ?", username).First(&result).Error
	return
}
func (this *adminDao) FindByMobilePhone(mobilePhone string) (result arraylist.List[entity.Admin], err error) {
	err = this.DBRead().Where("mobile_phone = ?", mobilePhone).First(&result).Error
	return
}
func (this *adminDao) Save(m *entity.Admin) (result *entity.Admin, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *adminDao) FindById(id int64) (result *entity.Admin, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *adminDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Admin{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *adminDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Admin], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
