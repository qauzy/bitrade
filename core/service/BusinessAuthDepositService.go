package service

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/PageModel"
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *BusinessAuthDepositService) FindAll(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.BusinessAuthDeposit], err error) {
	return this.BusinessAuthDepositDao.FindAll(predicate, pageModel.GetPageable())
}
func (this *BusinessAuthDepositService) FindAllByStatus(status *CommonStatus.CommonStatus) (result arraylist.List[entity.BusinessAuthDeposit], err error) {
	return this.BusinessAuthDepositDao.FindAllByStatus(status)
}
func (this *BusinessAuthDepositService) FindById(id int64) (result *entity.BusinessAuthDeposit, err error) {
	return this.BusinessAuthDepositDao.FindById(id).Get()
}
func (this *BusinessAuthDepositService) Save(businessAuthDeposit *entity.BusinessAuthDeposit) (err error) {
	this.BusinessAuthDepositDao.Save(businessAuthDeposit)
}
func (this *BusinessAuthDepositService) Update(businessAuthDeposit *entity.BusinessAuthDeposit) (err error) {
	this.BusinessAuthDepositDao.Save(businessAuthDeposit)
}
func NewBusinessAuthDepositService(businessAuthDepositDao *dao.BusinessAuthDepositDao) (ret *BusinessAuthDepositService) {
	ret = new(BusinessAuthDepositService)
	ret.BusinessAuthDepositDao = businessAuthDepositDao
	return
}

type BusinessAuthDepositService struct {
	BusinessAuthDepositDao dao.BusinessAuthDepositDao
	Base.BaseService
}
