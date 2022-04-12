package service

func (this *PoundageConvertEthService) Save(poundageConvertEth *entity.PoundageConvertEth) (result *entity.PoundageConvertEth, err error) {
	return this.PoundageConvertEthDao.Save(poundageConvertEth)
}
func (this *PoundageConvertEthService) GetAmountEthByDate(startDate string, endDate string) (result *math.BigDecimal, err error) {
	return this.PoundageConvertEthDao.GetAmountEthByDate(startDate, endDate)
}
func (this *PoundageConvertEthService) GetAmountEth(staTime string, arivedTime string) (result *math.BigDecimal, err error) {
	var startTime = staTime + " 00:00:00"
	var endTime = arivedTime + " 23:59:59"
	return this.PoundageConvertEthDao.GetAmountEth(startTime, endTime)
}
func NewPoundageConvertEthService(poundageConvertEthDao *dao.PoundageConvertEthDao) (ret *PoundageConvertEthService) {
	ret = new(PoundageConvertEthService)
	ret.PoundageConvertEthDao = poundageConvertEthDao
	return
}

type PoundageConvertEthService struct {
	PoundageConvertEthDao dao.PoundageConvertEthDao
	Base.BaseService[entity.PoundageConvertEth]
}
