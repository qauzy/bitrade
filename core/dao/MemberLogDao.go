package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type MemberLogDao interface {
	Save(m *entity.MemberLog) (result *entity.MemberLog, err error)
	FindById(id int64) (result *entity.MemberLog, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.MemberLog, err error)
}
type memberLogDao struct {
	*db.DB
}

func NewMemberLogDao(db *db.DB) (dao MemberLogDao) {
	dao = &memberLogDao{db}
	return
}
func (this *memberLogDao) Save(m *entity.MemberLog) (result *entity.MemberLog, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberLogDao) FindById(id int64) (result *entity.MemberLog, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberLogDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberLog{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberLogDao) FindAll(qp *types.QueryParam) (result []*entity.MemberLog, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
