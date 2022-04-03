package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type EmptionRecordDao interface {
	Save(m *entity.EmptionRecord) (result *entity.EmptionRecord, err error)
	FindById(id int64) (result *entity.EmptionRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.EmptionRecord], err error)
}
type emptionRecordDao struct {
	*db.DB
}

func NewEmptionRecordDao(db *db.DB) (dao EmptionRecordDao) {
	dao = &emptionRecordDao{db}
	return
}
func (this *emptionRecordDao) Save(m *entity.EmptionRecord) (result *entity.EmptionRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *emptionRecordDao) FindById(id int64) (result *entity.EmptionRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *emptionRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.EmptionRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *emptionRecordDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.EmptionRecord], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
