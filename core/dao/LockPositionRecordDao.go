package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"time"
)

type LockPositionRecordDao interface {
	FindByStatusAndUnlockTime(status *CommonStatus.CommonStatus, unlockTime time.Time) (result arraylist.List[entity.LockPositionRecord], err error)
	UnlockById(id int64, status *CommonStatus.CommonStatus) (err error)
	UnlockByIds(ids arraylist.List[int64], status *CommonStatus.CommonStatus) (err error)
	FindByMemberIdAndCoinAndStatus(memberId int64, coin *entity.Coin, status *CommonStatus.CommonStatus) (result arraylist.List[entity.LockPositionRecord], err error)
	Save(m *entity.LockPositionRecord) (result *entity.LockPositionRecord, err error)
	FindById(id int64) (result *entity.LockPositionRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.LockPositionRecord], err error)
}
type lockPositionRecordDao struct {
	*db.DB
}

func NewLockPositionRecordDao(db *db.DB) (dao LockPositionRecordDao) {
	dao = &lockPositionRecordDao{db}
	return
}
func (this *lockPositionRecordDao) FindByStatusAndUnlockTime(status *CommonStatus.CommonStatus, unlockTime time.Time) (result arraylist.List[entity.LockPositionRecord], err error) {
	err = this.DBRead().Where("status = ?", status).Where("unlock_time = ?", unlockTime).First(&result).Error
	return
}
func (this *lockPositionRecordDao) UnlockById(id int64, status *CommonStatus.CommonStatus) (err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update LockPositionRecord l set l.status = ? where l.id= ?", id, status)
	err = eng.Error
	return
}
func (this *lockPositionRecordDao) UnlockByIds(ids arraylist.List[int64], status *CommonStatus.CommonStatus) (err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update LockPositionRecord l set l.status = ? where l.id in (?)", ids, status)
	err = eng.Error
	return
}
func (this *lockPositionRecordDao) FindByMemberIdAndCoinAndStatus(memberId int64, coin *entity.Coin, status *CommonStatus.CommonStatus) (result arraylist.List[entity.LockPositionRecord], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Where("coin = ?", coin).Where("status = ?", status).First(&result).Error
	return
}
func (this *lockPositionRecordDao) Save(m *entity.LockPositionRecord) (result *entity.LockPositionRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *lockPositionRecordDao) FindById(id int64) (result *entity.LockPositionRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *lockPositionRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.LockPositionRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *lockPositionRecordDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.LockPositionRecord], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
