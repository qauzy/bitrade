package service

func (this *RobotTransactionService) Save(transaction *entity.RobotTransaction) (result *entity.RobotTransaction, err error) {
	return this.Dao.Save(transaction)
}
func NewRobotTransactionService(dao *dao.RobotTransactionDao) (ret *RobotTransactionService) {
	ret = new(RobotTransactionService)
	ret.Dao = dao
	return
}

type RobotTransactionService struct {
	Dao dao.RobotTransactionDao
	Base.BaseService[entity.RobotTransaction]
}
