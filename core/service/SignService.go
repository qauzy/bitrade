package service

import (
	"bitrade/core/constant/SignStatus"
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
)

func (this *SignService) SetDao(dao *dao.SignDao) (err error) {
	super.SetDao(dao)
}
func (this *SignService) FetchUnderway() (result *entity.Sign, err error) {
	return this.Dao.FindByStatus(SignStatus.UNDERWAY)
}
func (this *SignService) EarlyClosing(sign *entity.Sign) (err error) {
	sign.SetStatus(SignStatus.FINISH)
	dao.Save(sign)
}
func NewSignService() (ret *SignService) {
	ret = new(SignService)
	return
}

type SignService struct {
	Base.TopBaseService[entity.Sign, dao.SignDao]
}
