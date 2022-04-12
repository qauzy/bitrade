package service

func (this *DepartmentService) Save(department *entity.Department) (result *entity.Department, err error) {
	return this.DepartmentDao.Save(department)
}
func (this *DepartmentService) FindOne(departmentId int64) (result *entity.Department, err error) {
	return this.DepartmentDao.FindById(departmentId).Get()
}
func (this *DepartmentService) GetDepartmentDetail(departmentId int64) (result *entity.Department, err error) {
	var department = this.DepartmentDao.FindById(departmentId).Get()
	if department == nil {
		"该部门不存在"
	}
	return department
}
func (this *DepartmentService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.Department], err error) {
	return this.DepartmentDao.FindAll(predicate, pageable)
}
func (this *DepartmentService) Deletes(id int64) (result *util.MessageResult, err error) {
	var department = this.DepartmentDao.FindById(id).Get()
	var list = this.AdminDao.FindAllByDepartment(department)
	if list != nil && list.Size() > 0 {
		var result = MessageResult.Error("请先删除该部门下的所有用户")
		return result
	}
	this.DepartmentDao.DeleteById(id)
	return MessageResult.Success("删除成功")
}
func NewDepartmentService(em *persistence.EntityManager, departmentDao *dao.DepartmentDao, adminDao *dao.AdminDao) (ret *DepartmentService) {
	ret = new(DepartmentService)
	ret.Em = em
	ret.DepartmentDao = departmentDao
	ret.AdminDao = adminDao
	return
}

type DepartmentService struct {
	Em            *persistence.EntityManager
	DepartmentDao dao.DepartmentDao
	AdminDao      dao.AdminDao
	Base.BaseService
}
