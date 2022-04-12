package base

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type BaseDao[T any] interface {
	Save(m T) (result T, err error)
	FindById(id int64) (result T, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[T], err error)
}
type baseDao[T any] struct {
	*db.DB
}

func NewBaseDao[T any](db *db.DB) (dao BaseDao[T]) {
	dao = &baseDao[T]{db}
	return
}
func (this *baseDao[T]) Save(m T) (result T, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *baseDao[T]) FindById(id int64) (result T, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *baseDao[T]) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(T{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *baseDao[T]) FindAll(qp *types.QueryParam) (result arraylist.List[T], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
