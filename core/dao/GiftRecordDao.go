package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type GiftRecordDao interface {
	Save(m *entity.GiftRecord) (result *entity.GiftRecord, err error)
	FindById(id int64) (result *entity.GiftRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.GiftRecord], err error)
}
type giftRecordDao struct {
	*db.DB
}

func NewGiftRecordDao(db *db.DB) (dao GiftRecordDao) {
	dao = &giftRecordDao{db}
	return
}
func (this *giftRecordDao) Save(m *entity.GiftRecord) (result *entity.GiftRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *giftRecordDao) FindById(id int64) (result *entity.GiftRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *giftRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.GiftRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *giftRecordDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.GiftRecord], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
