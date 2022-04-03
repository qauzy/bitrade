package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type MemberSignRecordDao interface {
	Save(m *entity.MemberSignRecord) (result *entity.MemberSignRecord, err error)
	FindById(id int64) (result *entity.MemberSignRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberSignRecord], err error)
}
type memberSignRecordDao struct {
	*db.DB
}

func NewMemberSignRecordDao(db *db.DB) (dao MemberSignRecordDao) {
	dao = &memberSignRecordDao{db}
	return
}
func (this *memberSignRecordDao) Save(m *entity.MemberSignRecord) (result *entity.MemberSignRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberSignRecordDao) FindById(id int64) (result *entity.MemberSignRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberSignRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberSignRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberSignRecordDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberSignRecord], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
