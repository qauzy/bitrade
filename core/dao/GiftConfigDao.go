package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type GiftConfigDao interface {
	Save(m *entity.GiftConfig) (result *entity.GiftConfig, err error)
	FindById(id int64) (result *entity.GiftConfig, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.GiftConfig], err error)
}
type giftConfigDao struct {
	*db.DB
}

func NewGiftConfigDao(db *db.DB) (dao GiftConfigDao) {
	dao = &giftConfigDao{db}
	return
}
func (this *giftConfigDao) Save(m *entity.GiftConfig) (result *entity.GiftConfig, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *giftConfigDao) FindById(id int64) (result *entity.GiftConfig, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *giftConfigDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.GiftConfig{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *giftConfigDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.GiftConfig], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
