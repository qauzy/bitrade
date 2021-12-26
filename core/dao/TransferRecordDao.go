package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type TransferRecordDao interface {
	Save(m *entity.TransferRecord) (result *entity.TransferRecord, err error)
	FindById(id int64) (result *entity.TransferRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.TransferRecord, err error)
}
type transferRecordDao struct {
	*db.DB
}

func NewTransferRecordDao(db *db.DB) (dao TransferRecordDao) {
	dao = &transferRecordDao{db}
	return
}
func (this *transferRecordDao) Save(m *entity.TransferRecord) (result *entity.TransferRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *transferRecordDao) FindById(id int64) (result *entity.TransferRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *transferRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.TransferRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *transferRecordDao) FindAll(qp *types.QueryParam) (result []*entity.TransferRecord, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
