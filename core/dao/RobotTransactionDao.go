package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type RobotTransactionDao interface {
	Save(m *entity.RobotTransaction) (result *entity.RobotTransaction, err error)
	FindById(id int64) (result *entity.RobotTransaction, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.RobotTransaction], err error)
}
type robotTransactionDao struct {
	*db.DB
}

func NewRobotTransactionDao(db *db.DB) (dao RobotTransactionDao) {
	dao = &robotTransactionDao{db}
	return
}
func (this *robotTransactionDao) Save(m *entity.RobotTransaction) (result *entity.RobotTransaction, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *robotTransactionDao) FindById(id int64) (result *entity.RobotTransaction, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *robotTransactionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.RobotTransaction{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *robotTransactionDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.RobotTransaction], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
