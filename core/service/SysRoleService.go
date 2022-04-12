
package service

func (this *SysRoleService) SetDao(dao *dao.SysRoleDao) (err error) {
	super.SetDao(dao)
}
func (this *SysRoleService) FindOne(id int64) (result *entity.SysRole, err error) {
	var role = this.SysRoleDao.FindById(id).Get()
	return role
}
func (this *SysRoleService) GetPermissions(roleId int64) (result arraylist.List[entity.SysPermission], err error) {
	var sysRole = this.FindOne(roleId)
	var list = sysRole.GetPermissions()
	return list
}
func (this *SysRoleService) Deletes(id int64) (result *util.MessageResult, err error) {
	var list = this.AdminDao.FindAllByRoleId(id)
	if list != nil && list.Size() > 0 {
		return MessageResult.Error("删除失败，请先删除该角色下的所有用户")
	}
	this.SysRoleDao.DeleteById(id)
	return MessageResult.Success("删除成功")
}
func (this *SysRoleService) ToMenus(sysPermissions arraylist.List[entity.SysPermission], parentId int64) (result arraylist.List[core.Menu], err error) {
	return sysPermissions.Stream().Filter(func(x interface {
	}) {
		x.GetParentId().Equals(parentId)
	}).Sorted(Comparator.Comparing(SysPermission : GetSort)).Map(func(x interface {
	}) {
		new(Menu).SetId(x.GetId()).SetName(x.GetName()).SetParentId(x.GetParentId()).SetSort(x.GetSort()).SetTitle(x.GetTitle()).SetDescription(x.GetDescription()).SetSubMenu(this.ToMenus(sysPermissions, x.GetId()))
	}).Collect(Collectors.ToList())
}
func (this *SysRoleService) Save(sysRole *entity.SysRole) (result *entity.SysRole, err error) {
	return this.SysRoleDao.Save(sysRole)
}
func (this *SysRoleService) UpdateDetail(sysRole *entity.SysRole) (result int, err error) {
	return this.SysRoleDao.UpdateSysRole(sysRole.GetDescription(), sysRole.GetRole(), sysRole.GetId())
}
func (this *SysRoleService) GetAllPermission() (result arraylist.List[entity.SysPermission], err error) {
	return this.SysPermissionDao.FindAll()
}
func (this *SysRoleService) GetAllSysRole() (result arraylist.List[entity.SysRole], err error) {
	return this.SysRoleDao.FindAllSysRole()
}
func (this *SysRoleService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.SysRole], err error) {
	return this.SysRoleDao.FindAll(predicate, pageable)
}
func NewSysRoleService(adminService AdminService, sysRoleDao *dao.SysRoleDao, adminDao *dao.AdminDao, sysPermissionDao *dao.SysPermissionDao) (ret *SysRoleService) {
	ret = new(SysRoleService)
	ret.AdminService = adminService
	ret.SysRoleDao = sysRoleDao
	ret.AdminDao = adminDao
	ret.SysPermissionDao = sysPermissionDao
	return
}

type SysRoleService struct {
	AdminService     *AdminService
	SysRoleDao       dao.SysRoleDao
	AdminDao         dao.AdminDao
	SysPermissionDao dao.SysPermissionDao
	Base.TopBaseService[entity.SysRole, dao.SysRoleDao]
}

