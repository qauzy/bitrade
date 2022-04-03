package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type DataDictionaryDao interface {
	FindByBond(bond string) (result *entity.DataDictionary, err error)
	Save(m *entity.DataDictionary) (result *entity.DataDictionary, err error)
	FindById(id int64) (result *entity.DataDictionary, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.DataDictionary], err error)
}
type dataDictionaryDao struct {
	*db.DB
}

func NewDataDictionaryDao(db *db.DB) (dao DataDictionaryDao) {
	dao = &dataDictionaryDao{db}
	return
}
func (this *dataDictionaryDao) FindByBond(bond string) (result *entity.DataDictionary, err error) {
	err = this.DBRead().Where("bond = ?", bond).First(&result).Error
	return
}
func (this *dataDictionaryDao) Save(m *entity.DataDictionary) (result *entity.DataDictionary, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *dataDictionaryDao) FindById(id int64) (result *entity.DataDictionary, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *dataDictionaryDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.DataDictionary{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *dataDictionaryDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.DataDictionary], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
