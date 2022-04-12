package service

func (this *SmsService) GetByStatus() (result *dto.SmsDTO, err error) {
	return this.SmsDao.FindBySmsStatus()
}
func NewSmsService(smsDao *dao.SmsDao) (ret *SmsService) {
	ret = new(SmsService)
	ret.SmsDao = smsDao
	return
}

type SmsService struct {
	SmsDao dao.SmsDao
	Base.BaseService
}
