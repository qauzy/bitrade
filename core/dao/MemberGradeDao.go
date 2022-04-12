package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type MemberGradeDao interface {
	Save(m *entity.MemberGrade) (result *entity.MemberGrade, err error)
	FindById(id int64) (result *entity.MemberGrade, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberGrade], err error)
}
type memberGradeDao struct {
	*db.DB
}

func NewMemberGradeDao(db *db.DB) (dao MemberGradeDao) {
	dao = &memberGradeDao{db}
	return
}
func (this *memberGradeDao) Save(m *entity.MemberGrade) (result *entity.MemberGrade, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberGradeDao) FindById(id int64) (result *entity.MemberGrade, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberGradeDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberGrade{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberGradeDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberGrade], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
