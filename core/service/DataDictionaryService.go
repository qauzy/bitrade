package service

func (this *DataDictionaryService) SetDao(dao *dao.DataDictionaryDao) (err error) {
	super.SetDao(dao)
}
func (this *DataDictionaryService) FindByBond(bond string) (result *entity.DataDictionary, err error) {
	return this.DataDictionaryDao.FindByBond(bond)
}
func NewDataDictionaryService(dataDictionaryDao *dao.DataDictionaryDao) (ret *DataDictionaryService) {
	ret = new(DataDictionaryService)
	ret.DataDictionaryDao = dataDictionaryDao
	return
}

type DataDictionaryService struct {
	DataDictionaryDao dao.DataDictionaryDao
	Base.TopBaseService[entity.DataDictionary, dao.DataDictionaryDao]
}
