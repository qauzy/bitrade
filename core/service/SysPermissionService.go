package service

func (this *SysPermissionService) SetDao(dao *dao.SysPermissionDao) (err error) {
	super.SetDao(dao)
}
func (this *SysPermissionService) FindOne(id int64) (result *entity.SysPermission, err error) {
	return this.SysPermissionDao.FindById(id).Get()
}
func (this *SysPermissionService) FindAll() (result arraylist.List[entity.SysPermission], err error) {
	return this.SysPermissionDao.FindAll()
}
func (this *SysPermissionService) Save(sysPermission *entity.SysPermission) (result *entity.SysPermission, err error) {
	return this.SysPermissionDao.Save(sysPermission)
}
func (this *SysPermissionService) Deletes(ids []int64) (err error) {
	for _, id := range ids {
		this.SysPermissionDao.DeletePermission(id)
		this.SysPermissionDao.DeleteById(id)
	}
}
func (this *SysPermissionService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.SysPermission], err error) {
	return this.SysPermissionDao.FindAll(predicate, pageable)
}
func (this *SysPermissionService) FindByPermissionName(name string) (result *entity.SysPermission, err error) {
	return this.SysPermissionDao.FindSysPermissionByName(name)
}
func NewSysPermissionService(sysPermissionDao *dao.SysPermissionDao) (ret *SysPermissionService) {
	ret = new(SysPermissionService)
	ret.SysPermissionDao = sysPermissionDao
	return
}

type SysPermissionService struct {
	SysPermissionDao dao.SysPermissionDao
	Base.BaseService[entity.SysPermission]
}
