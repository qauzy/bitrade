package service

import (
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *TransferRecordService) Save(transferRecord *entity.TransferRecord) (result *entity.TransferRecord, err error) {
	return this.TransferRecordDao.Save(transferRecord)
}
func (this *TransferRecordService) FindAllByMemberId(memberId int64, page int, pageSize int) (result arraylist.List[entity.TransferRecord], err error) {
	var orders = Criteria.SortStatic("id.desc")
	var pageRequest = PageRequest.Of(page, pageSize, orders)
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("memberId", memberId, false))
	return this.TransferRecordDao.FindAll(specification, pageRequest)
}
func NewTransferRecordService(transferRecordDao dao.TransferRecordDao) (ret *TransferRecordService) {
	ret = new(TransferRecordService)
	ret.TransferRecordDao = transferRecordDao
	return
}

type TransferRecordService struct {
	TransferRecordDao dao.TransferRecordDao
	Base.BaseService
}
