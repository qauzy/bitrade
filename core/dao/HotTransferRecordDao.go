package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type HotTransferRecordDao interface {
	Save(m *entity.HotTransferRecord) (result *entity.HotTransferRecord, err error)
	FindById(id int64) (result *entity.HotTransferRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.HotTransferRecord, err error)
}
type hotTransferRecordDao struct {
	*db.DB
}

func NewHotTransferRecordDao(db *db.DB) (dao HotTransferRecordDao) {
	dao = &hotTransferRecordDao{db}
	return
}
func (this *hotTransferRecordDao) Save(m *entity.HotTransferRecord) (result *entity.HotTransferRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *hotTransferRecordDao) FindById(id int64) (result *entity.HotTransferRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *hotTransferRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.HotTransferRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *hotTransferRecordDao) FindAll(qp *types.QueryParam) (result []*entity.HotTransferRecord, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
