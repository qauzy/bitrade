package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type InitPlateDao interface {
	FindInitPlateBySymbol(symbol string) (result entity.InitPlate, err error)
	Save(m *entity.InitPlate) (result *entity.InitPlate, err error)
	FindById(id int64) (result *entity.InitPlate, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.InitPlate, err error)
}
type initPlateDao struct {
	*db.DB
}

func NewInitPlateDao(db *db.DB) (dao InitPlateDao) {
	dao = &initPlateDao{db}
	return
}
func (this *initPlateDao) FindInitPlateBySymbol(symbol string) (result entity.InitPlate, err error) {
	err = this.DBRead().Where("symbol = ?", symbol).First(&result).Error
	return
}
func (this *initPlateDao) Save(m *entity.InitPlate) (result *entity.InitPlate, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *initPlateDao) FindById(id int64) (result *entity.InitPlate, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *initPlateDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.InitPlate{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *initPlateDao) FindAll(qp *types.QueryParam) (result []*entity.InitPlate, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
