package service

func (this *DepositRecordService) GetAll() (result arraylist.List[entity.DepositRecord], err error) {
	return this.DepositRecordDao.FindAll()
}
func (this *DepositRecordService) List(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.DepositRecord], err error) {
	return this.DepositRecordDao.FindAll(predicate, pageModel.GetPageable())
}
func (this *DepositRecordService) FindAll(predicate *types.Predicate) (result arraylist.List[entity.DepositRecord], err error) {
	return this.DepositRecordDao.FindAll(predicate).(arraylist.List[entity.DepositRecord])
}
func (this *DepositRecordService) FindOne(id string) (result *entity.DepositRecord, err error) {
	return this.DepositRecordDao.FindById(id)
}
func (this *DepositRecordService) Update(depositRecord *entity.DepositRecord) (err error) {
	this.DepositRecordDao.Save(depositRecord)
}
func (this *DepositRecordService) Create(depositRecord *entity.DepositRecord) (err error) {
	this.DepositRecordDao.Save(depositRecord)
}
func (this *DepositRecordService) FindByMemberAndStatus(member *entity.Member, status *DepositStatusEnum.DepositStatusEnum) (result arraylist.List[entity.DepositRecord], err error) {
	return this.DepositRecordDao.FindByMemberAndStatus(member, status)
}
func (this *DepositRecordService) Save(depositRecord *entity.DepositRecord) (result *entity.DepositRecord, err error) {
	return this.DepositRecordDao.Save(depositRecord)
}
func NewDepositRecordService(depositRecordDao *dao.DepositRecordDao) (ret *DepositRecordService) {
	ret = new(DepositRecordService)
	ret.DepositRecordDao = depositRecordDao
	return
}

type DepositRecordService struct {
	DepositRecordDao dao.DepositRecordDao
	Base.BaseService
}
