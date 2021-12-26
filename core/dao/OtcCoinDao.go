package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type OtcCoinDao interface {
	FindOtcCoinByUnitAndStatus(unit string, status CommonStatus.CommonStatus) (result entity.OtcCoin, err error)
	FindAllByStatus(status CommonStatus.CommonStatus) (result []entity.OtcCoin, err error)
	FindOtcCoinByUnit(unit string) (result entity.OtcCoin, err error)
	FindAllUnits() (result []string, err error)
	Save(m *entity.OtcCoin) (result *entity.OtcCoin, err error)
	FindById(id int64) (result *entity.OtcCoin, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.OtcCoin, err error)
}
type otcCoinDao struct {
	*db.DB
}

func NewOtcCoinDao(db *db.DB) (dao OtcCoinDao) {
	dao = &otcCoinDao{db}
	return
}
func (this *otcCoinDao) FindOtcCoinByUnitAndStatus(unit string, status CommonStatus.CommonStatus) (result entity.OtcCoin, err error) {
	err = this.DBRead().Where("unit = ?", unit).Where("status = ?", status).First(&result).Error
	return
}
func (this *otcCoinDao) FindAllByStatus(status CommonStatus.CommonStatus) (result []entity.OtcCoin, err error) {
	err = this.DBRead().Where("status = ?", status).Find(&result).Error
	return
}
func (this *otcCoinDao) FindOtcCoinByUnit(unit string) (result entity.OtcCoin, err error) {
	err = this.DBRead().Where("unit = ?", unit).First(&result).Error
	return
}
func (this *otcCoinDao) FindAllUnits() (result []string, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select distinct a.unit from OtcCoin a where a.status = 0")
	err = eng.Error
	return
}
func (this *otcCoinDao) Save(m *entity.OtcCoin) (result *entity.OtcCoin, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *otcCoinDao) FindById(id int64) (result *entity.OtcCoin, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *otcCoinDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.OtcCoin{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *otcCoinDao) FindAll(qp *types.QueryParam) (result []*entity.OtcCoin, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
