package service

import (
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
)

func (this *MemberSignRecordService) SetDao(dao *dao.MemberSignRecordDao) (err error) {
	super.SetDao(dao)
}
func NewMemberSignRecordService() (ret *MemberSignRecordService) {
	ret = new(MemberSignRecordService)
	return
}

type MemberSignRecordService struct {
	Base.TopBaseService[entity.MemberSignRecord, dao.MemberSignRecordDao]
}
