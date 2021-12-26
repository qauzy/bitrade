package dao

import (
	"bitrade/core/constant/WithdrawStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"time"
)

type WithdrawRecordDao interface {
	GetWithdrawStatistics(date string) (result [][]interface{}, err error)
	CountAllByStatus(status WithdrawStatus.WithdrawStatus) (result int64, err error)
	FindByIds(ids []int) (result []entity.WithdrawRecord, err error)
	CountWithdrawAmountByTimeAndMemberIdAndCoin(startTime time.Time, endTime time.Time, coin entity.Coin) (result float64, err error)
	Save(m *entity.WithdrawRecord) (result *entity.WithdrawRecord, err error)
	FindById(id int64) (result *entity.WithdrawRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.WithdrawRecord, err error)
}
type withdrawRecordDao struct {
	*db.DB
}

func NewWithdrawRecordDao(db *db.DB) (dao WithdrawRecordDao) {
	dao = &withdrawRecordDao{db}
	return
}
func (this *withdrawRecordDao) GetWithdrawStatistics(date string) (result [][]interface{}, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select a.unit , sum(b.total_amount) amount,sum(b.fee) from withdraw_record b,coin a where a.name = b.coin_id and date_format(b.deal_time,'%Y-%m-%d') = ? and b.status = 3 group by a.unit", date)
	err = eng.Error
	return
}
func (this *withdrawRecordDao) CountAllByStatus(status WithdrawStatus.WithdrawStatus) (result int64, err error) {
	return
}
func (this *withdrawRecordDao) FindByIds(ids []int) (result []entity.WithdrawRecord, err error) {
	err = this.DBRead().Where("ids = ?", ids).First(&result).Error
	return
}
func (this *withdrawRecordDao) CountWithdrawAmountByTimeAndMemberIdAndCoin(startTime time.Time, endTime time.Time, coin entity.Coin) (result float64, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select sum(w.totalAmount) from WithdrawRecord w where w.status <> 2 and w.coin= ? and w.createTime between ? and ?", startTime, endTime, coin)
	err = eng.Error
	return
}
func (this *withdrawRecordDao) Save(m *entity.WithdrawRecord) (result *entity.WithdrawRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *withdrawRecordDao) FindById(id int64) (result *entity.WithdrawRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *withdrawRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.WithdrawRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *withdrawRecordDao) FindAll(qp *types.QueryParam) (result []*entity.WithdrawRecord, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
