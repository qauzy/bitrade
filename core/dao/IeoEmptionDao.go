package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/math"
	"github.com/qauzy/util/lists/arraylist"
)

type IeoEmptionDao interface {
	FindbyCondition(id int64, times string) (result *entity.IeoEmption, err error)
	SubAmount(id int64, receAmount *math.BigDecimal) (result int, err error)
	Save(m *entity.IeoEmption) (result *entity.IeoEmption, err error)
	FindById(id int64) (result *entity.IeoEmption, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.IeoEmption], err error)
}
type ieoEmptionDao struct {
	*db.DB
}

func NewIeoEmptionDao(db *db.DB) (dao IeoEmptionDao) {
	dao = &ieoEmptionDao{db}
	return
}
func (this *ieoEmptionDao) FindbyCondition(id int64, times string) (result *entity.IeoEmption, err error) {
	eng := this.DBWrite().Table("ieo_emption").Select("*").Where("id = ? and start_time <= ? and end_time > ?", id, times).Find(&result)
	err = eng.Error
	return
}
func (this *ieoEmptionDao) SubAmount(id int64, receAmount *math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("UPDATE ieo_emption set surplus_amount=surplus_amount-? where id=? AND surplus_amount>=?", id, receAmount)
	err = eng.Error
	return
}
func (this *ieoEmptionDao) Save(m *entity.IeoEmption) (result *entity.IeoEmption, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *ieoEmptionDao) FindById(id int64) (result *entity.IeoEmption, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *ieoEmptionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.IeoEmption{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *ieoEmptionDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.IeoEmption], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
