package service

import (
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/maps/hashmap"
	"github.com/qauzy/chocolate/xtime"
)

func (this *AdminService) SetDao(dao *dao.AdminDao) (err error) {
	super.SetDao(dao)
}
func (this *AdminService) Login(username string, password string) (result *entity.Admin, err error) {
	var admin = dao.FindAdminByUsernameAndPassword(username, password)
	return admin
}
func (this *AdminService) Update(date xtime.Xtime, ip string, memberId int64) (result int, err error) {
	return dao.UpdateAdminLastTimeAndIp(date, ip, memberId)
}
func (this *AdminService) FindOne(id int64) (result *entity.Admin, err error) {
	return dao.FindById(id).Get()
}
func (this *AdminService) FindAdminDetail(id int64) (result *hashmap.Map[string, interface {
}], err error) {
	var sql = "select a.id,a.role_id roleId,a.department_id departmentId,a.real_name realName,a.avatar,a.email,a.enable,a.mobile_phone mobilePhone,a.qq,a.username, " + "d.name as 'departmentName',r.role from admin a LEFT join department d on a.department_id=d.id LEFT JOIN admin_role r on a.role_id=r.id WHERE a.id=:adminId "
	var query = this.Em.CreateNativeQuery(sql)
	//设置结果转成Map类型
	query.Unwrap(hibernate.SQLQuery).SetResultTransformer(Transformers.ALIAS_TO_ENTITY_MAP)
	var object = query.SetParameter("adminId", id).GetSingleResult()
	var oMap = object.(*hashmap.Map[string, interface {
	}])
	return oMap
}
func (this *AdminService) SaveAdmin(admin *entity.Admin) (result *entity.Admin, err error) {
	return dao.Save(admin)
}
func (this *AdminService) FindByUsername(username string) (result *entity.Admin, err error) {
	return dao.FindByUsername(username)
}
func (this *AdminService) DeleteBatch(roleId int64) (err error) {
	dao.DeleteBatch(roleId)
}
func (this *AdminService) DeleteBatch(ids []int64) (err error) {
	for _, id := range ids {
		delete(id)
	}
}
func (this *AdminService) FindAll(predicate predicate, pageable *domain.Pageable) (result domain.Page[entity.Admin], err error) {
	return dao.FindAll(predicate, pageable)
}
func (this *AdminService) FindByMobilePhone(mobilePhone string) (result arraylist.List[entity.Admin], err error) {
	return dao.FindByMobilePhone(mobilePhone)
}
func NewAdminService(em *persistence.EntityManager) (ret *AdminService) {
	ret = new(AdminService)
	ret.Em = em
	return
}

type AdminService struct {
	Em *persistence.EntityManager
	Base.TopBaseService[entity.Admin, dao.AdminDao]
}
