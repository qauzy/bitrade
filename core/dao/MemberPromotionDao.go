package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type MemberPromotionDao interface {
	Save(m *entity.MemberPromotion) (result *entity.MemberPromotion, err error)
	FindById(id int64) (result *entity.MemberPromotion, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberPromotion], err error)
}
type memberPromotionDao struct {
	*db.DB
}

func NewMemberPromotionDao(db *db.DB) (dao MemberPromotionDao) {
	dao = &memberPromotionDao{db}
	return
}
func (this *memberPromotionDao) Save(m *entity.MemberPromotion) (result *entity.MemberPromotion, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberPromotionDao) FindById(id int64) (result *entity.MemberPromotion, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberPromotionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberPromotion{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberPromotionDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberPromotion], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
