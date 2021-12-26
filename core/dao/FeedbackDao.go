package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type FeedbackDao interface {
	Save(m *entity.Feedback) (result *entity.Feedback, err error)
	FindById(id int64) (result *entity.Feedback, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Feedback, err error)
}
type feedbackDao struct {
	*db.DB
}

func NewFeedbackDao(db *db.DB) (dao FeedbackDao) {
	dao = &feedbackDao{db}
	return
}
func (this *feedbackDao) Save(m *entity.Feedback) (result *entity.Feedback, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *feedbackDao) FindById(id int64) (result *entity.Feedback, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *feedbackDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Feedback{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *feedbackDao) FindAll(qp *types.QueryParam) (result []*entity.Feedback, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
