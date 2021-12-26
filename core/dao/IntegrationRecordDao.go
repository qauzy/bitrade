package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type IntegrationRecordDao interface {
	Save(m *entity.IntegrationRecord) (result *entity.IntegrationRecord, err error)
	FindById(id int64) (result *entity.IntegrationRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.IntegrationRecord, err error)
}
type integrationRecordDao struct {
	*db.DB
}

func NewIntegrationRecordDao(db *db.DB) (dao IntegrationRecordDao) {
	dao = &integrationRecordDao{db}
	return
}
func (this *integrationRecordDao) Save(m *entity.IntegrationRecord) (result *entity.IntegrationRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *integrationRecordDao) FindById(id int64) (result *entity.IntegrationRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *integrationRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.IntegrationRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *integrationRecordDao) FindAll(qp *types.QueryParam) (result []*entity.IntegrationRecord, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
