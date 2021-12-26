package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type DividendStartRecordDao interface {
	FindAllByTimeAndUnit(start int64, end int64, unit string) (result []entity.DividendStartRecord, err error)
	Save(m *entity.DividendStartRecord) (result *entity.DividendStartRecord, err error)
	FindById(id int64) (result *entity.DividendStartRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.DividendStartRecord, err error)
}
type dividendStartRecordDao struct {
	*db.DB
}

func NewDividendStartRecordDao(db *db.DB) (dao DividendStartRecordDao) {
	dao = &dividendStartRecordDao{db}
	return
}
func (this *dividendStartRecordDao) FindAllByTimeAndUnit(start int64, end int64, unit string) (result []entity.DividendStartRecord, err error) {
	err = this.DBRead().Where("time = ?", start).Where("unit = ?", end).Find(&result).Error
	return
}
func (this *dividendStartRecordDao) Save(m *entity.DividendStartRecord) (result *entity.DividendStartRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *dividendStartRecordDao) FindById(id int64) (result *entity.DividendStartRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *dividendStartRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.DividendStartRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *dividendStartRecordDao) FindAll(qp *types.QueryParam) (result []*entity.DividendStartRecord, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
