package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type CountryDao interface {
	FindAllOrderBySort() (result []entity.Country, err error)
	FindByZhName(zhname string) (result entity.Country, err error)
	FindByLocalCurrency(localCurrency string) (result []entity.Country, err error)
	Save(m *entity.Country) (result *entity.Country, err error)
	FindById(id int64) (result *entity.Country, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Country, err error)
}
type countryDao struct {
	*db.DB
}

func NewCountryDao(db *db.DB) (dao CountryDao) {
	dao = &countryDao{db}
	return
}
func (this *countryDao) FindAllOrderBySort() (result []entity.Country, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select a from Country a order by a.sort")
	err = eng.Error
	return
}
func (this *countryDao) FindByZhName(zhname string) (result entity.Country, err error) {
	err = this.DBRead().Where("zh_name = ?", zhname).First(&result).Error
	return
}
func (this *countryDao) FindByLocalCurrency(localCurrency string) (result []entity.Country, err error) {
	err = this.DBRead().Where("local_currency = ?", localCurrency).First(&result).Error
	return
}
func (this *countryDao) Save(m *entity.Country) (result *entity.Country, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *countryDao) FindById(id int64) (result *entity.Country, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *countryDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Country{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *countryDao) FindAll(qp *types.QueryParam) (result []*entity.Country, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}