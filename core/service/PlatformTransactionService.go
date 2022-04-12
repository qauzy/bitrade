package service

func (this *PlatformTransactionService) Add(amount *math.BigDecimal, direction int, oType int, bizOrderId string) (err error) {
	var transaction = new(entity.PlatformTransaction)
	transaction.SetAmount(amount)
	transaction.SetDirection(direction)
	transaction.SetType(oType)
	transaction.SetBizOrderId(bizOrderId)
	transaction.SetDay(DateUtil.GetDate())
	transaction.SetTime(DateUtil.GetCurrentDate())
	this.Save(transaction)
}
func (this *PlatformTransactionService) Save(transaction *entity.PlatformTransaction) (err error) {
	this.PlatformTransactionDao.Save(transaction)
}
func NewPlatformTransactionService(platformTransactionDao *dao.PlatformTransactionDao) (ret *PlatformTransactionService) {
	ret = new(PlatformTransactionService)
	ret.PlatformTransactionDao = platformTransactionDao
	return
}

type PlatformTransactionService struct {
	PlatformTransactionDao dao.PlatformTransactionDao
}
