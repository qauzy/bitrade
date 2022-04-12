package service

func (this *HotTransferRecordService) SetDao(dao *dao.HotTransferRecordDao) (err error) {
	super.SetDao(dao)
}
func NewHotTransferRecordService() (ret *HotTransferRecordService) {
	ret = new(HotTransferRecordService)
	return
}

type HotTransferRecordService struct {
	Base.TopBaseService[entity.HotTransferRecord, dao.HotTransferRecordDao]
}
