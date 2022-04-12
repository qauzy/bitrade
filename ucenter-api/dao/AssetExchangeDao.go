package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/ucenter-api/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type AssetExchangeDao interface {
	FindAllByToUnit(unit string) (result arraylist.List[entity.AssetExchangeCoin], err error)
	FindByFromUnitAndToUnit(fromUnit string, toUnit string) (result *entity.AssetExchangeCoin, err error)
	Save(m *entity.AssetExchangeCoin) (result *entity.AssetExchangeCoin, err error)
	FindById(id int64) (result *entity.AssetExchangeCoin, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.AssetExchangeCoin], err error)
}
type assetExchangeDao struct {
	*db.DB
}

func NewAssetExchangeDao(db *db.DB) (dao AssetExchangeDao) {
	dao = &assetExchangeDao{db}
	return
}
func (this *assetExchangeDao) FindAllByToUnit(unit string) (result arraylist.List[entity.AssetExchangeCoin], err error) {
	err = this.DBRead().Where("to_unit = ?", unit).Find(&result).Error
	return
}
func (this *assetExchangeDao) FindByFromUnitAndToUnit(fromUnit string, toUnit string) (result *entity.AssetExchangeCoin, err error) {
	err = this.DBRead().Where("from_unit = ?", fromUnit).Where("to_unit = ?", toUnit).First(&result).Error
	return
}
func (this *assetExchangeDao) Save(m *entity.AssetExchangeCoin) (result *entity.AssetExchangeCoin, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *assetExchangeDao) FindById(id int64) (result *entity.AssetExchangeCoin, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *assetExchangeDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.AssetExchangeCoin{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *assetExchangeDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.AssetExchangeCoin], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
