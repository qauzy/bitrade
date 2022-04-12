package service

func (this *MemberLevelService) FindAll() (result arraylist.List[entity.MemberLevel], err error) {
	return this.MemberLevelDao.FindAll()
}
func (this *MemberLevelService) FindOne(id int64) (result *entity.MemberLevel, err error) {
	return this.MemberLevelDao.FindById(id).Get()
}
func (this *MemberLevelService) FindDefault() (result *entity.MemberLevel, err error) {
	return this.MemberLevelDao.FindOneByIsDefault(true)
}
func (this *MemberLevelService) UpdateDefault() (result int, err error) {
	return this.MemberLevelDao.UpdateDefault()
}
func (this *MemberLevelService) Save(memberLevel *entity.MemberLevel) (result *entity.MemberLevel, err error) {
	return this.MemberLevelDao.Save(memberLevel)
}
func NewMemberLevelService(memberLevelDao *dao.MemberLevelDao) (ret *MemberLevelService) {
	ret = new(MemberLevelService)
	ret.MemberLevelDao = memberLevelDao
	return
}

type MemberLevelService struct {
	MemberLevelDao dao.MemberLevelDao
	Base.BaseService
}
