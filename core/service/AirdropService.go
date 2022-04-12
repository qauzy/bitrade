package service

import (
	"bitrade/core/constant/PageModel"
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
)

func (this *AirdropService) Save(airdrop *entity.Airdrop) (result *entity.Airdrop, err error) {
	return this.AirdropDao.Save(airdrop)
}
func (this *AirdropService) FindById(id int64) (result *entity.Airdrop, err error) {
	return this.AirdropDao.FindById(id)
}
func (this *AirdropService) FindAll(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.Airdrop], err error) {
	return this.AirdropDao.FindAll(predicate, pageModel.GetPageable())
}
func NewAirdropService(airdropDao *dao.AirdropDao) (ret *AirdropService) {
	ret = new(AirdropService)
	ret.AirdropDao = airdropDao
	return
}

type AirdropService struct {
	AirdropDao dao.AirdropDao
	Base.BaseService
}
