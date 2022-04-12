package service

func (this *AppRevisionService) SetDao(dao *dao.AppRevisionDao) (err error) {
	super.SetDao(dao)
}
func (this *AppRevisionService) FindRecentVersion(p *Platform.Platform) (result *entity.AppRevision, err error) {
	return dao.FindAppRevisionByPlatformOrderByIdDesc(p)
}
func NewAppRevisionService() (ret *AppRevisionService) {
	ret = new(AppRevisionService)
	return
}

type AppRevisionService struct {
	Base.TopBaseService[entity.AppRevision, dao.AppRevisionDao]
}
