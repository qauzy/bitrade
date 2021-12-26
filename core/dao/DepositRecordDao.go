package dao

import (
	"bitrade/core/constant/DepositStatusEnum"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type DepositRecordDao interface {
	FindByMemberAndStatus(member entity.Member, status DepositStatusEnum.DepositStatusEnum) (result []entity.DepositRecord, err error)
	Save(m *entity.DepositRecord) (result *entity.DepositRecord, err error)
	FindById(id int64) (result *entity.DepositRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.DepositRecord, err error)
}
type depositRecordDao struct {
	*db.DB
}

func NewDepositRecordDao(db *db.DB) (dao DepositRecordDao) {
	dao = &depositRecordDao{db}
	return
}
func (this *depositRecordDao) FindByMemberAndStatus(member entity.Member, status DepositStatusEnum.DepositStatusEnum) (result []entity.DepositRecord, err error) {
	err = this.DBRead().Where("member = ?", member).Where("status = ?", status).First(&result).Error
	return
}
func (this *depositRecordDao) Save(m *entity.DepositRecord) (result *entity.DepositRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *depositRecordDao) FindById(id int64) (result *entity.DepositRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *depositRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.DepositRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *depositRecordDao) FindAll(qp *types.QueryParam) (result []*entity.DepositRecord, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
