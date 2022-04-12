package service

func (this *DividendStartRecordService) SetDao(dao *dao.DividendStartRecordDao) (err error) {
	super.SetDao(dao)
}
func (this *DividendStartRecordService) MatchRecord(start int64, end int64, unit string) (result arraylist.List[entity.DividendStartRecord], err error) {
	return dao.FindAllByTimeAndUnit(start, end, unit)
}
func (this *DividendStartRecordService) Save(dividendStartRecord *entity.DividendStartRecord) (result *entity.DividendStartRecord, err error) {
	return dao.Save(dividendStartRecord)
}
func NewDividendStartRecordService() (ret *DividendStartRecordService) {
	ret = new(DividendStartRecordService)
	return
}

type DividendStartRecordService struct {
	Base.TopBaseService[entity.DividendStartRecord, dao.DividendStartRecordDao]
}
