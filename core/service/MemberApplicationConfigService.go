package service

func (this *MemberApplicationConfigService) SetDao(dao *dao.MemberApplicationConfigDao) (err error) {
	super.SetDao(dao)
}
func (this *MemberApplicationConfigService) Get() (result *entity.MemberApplicationConfig, err error) {
	var list = dao.FindAll()
	if list != nil && list.Size() > 0 {
		return list.Get(0)
	} else {
		return nil
	}
}
func NewMemberApplicationConfigService() (ret *MemberApplicationConfigService) {
	ret = new(MemberApplicationConfigService)
	return
}

type MemberApplicationConfigService struct {
	Base.TopBaseService[entity.MemberApplicationConfig, dao.MemberApplicationConfigDao]
}
