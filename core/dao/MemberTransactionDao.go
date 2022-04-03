
package dao

import (
	"bitrade/core/constant/TransactionType"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/maps/hashmap"
	"time"
)

type MemberTransactionDao interface {
	FindTransactionSum(memberId int64, types arraylist.List[TransactionType.TransactionType]) (result *hashmap.Map[string, interface {
	}], err error)
	FindMatchTransactionSum(memberId int64, symbol string, oType *TransactionType.TransactionType, startTime time.Time, endTime time.Time) (result *hashmap.Map[string, interface {
	}], err error)
	FindMatchTransactionSum(symbol string, oType *TransactionType.TransactionType, startTime time.Time, endTime time.Time) (result *hashmap.Map[string, interface {
	}], err error)
	FindAllByCreateTime(beginDate string, endDate string) (result arraylist.List[entity.MemberTransaction], err error)
	Save(m *entity.MemberTransaction) (result *entity.MemberTransaction, err error)
	FindById(id int64) (result *entity.MemberTransaction, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberTransaction], err error)
}
type memberTransactionDao struct {
	 *db.DB
}

func NewMemberTransactionDao(db *db.DB) (dao MemberTransactionDao) {
	dao = &memberTransactionDao{db}
	return
}
func (this *memberTransactionDao) FindTransactionSum(memberId int64, types arraylist.List[TransactionType.TransactionType]) (result *hashmap.Map[string, interface {
}], err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("select sum(t.amount)  as amount from MemberTransaction t where  t.amount > 0 and t.memberId = ? and t.type in ?",memberId,types)
	err = eng.Error
	return
}
func (this *memberTransactionDao) FindMatchTransactionSum(memberId int64, symbol string, oType *TransactionType.TransactionType, startTime time.Time, endTime time.Time) (result *hashmap.Map[string, interface {
}], err error) {
	eng := this.DBWrite().Table("MemberTransaction as t").Select("sum(t.amount) as amount").Where("t.flag = 0 and t.memberId = ? and t.symbol = ? and t.type = ? and t.createTime >= ? and t.createTime <= ?",memberId,symbol,type,startTime,endTime).Find(&result)
	err = eng.Error
	return
}
func (this *memberTransactionDao) FindMatchTransactionSum(symbol string, oType *TransactionType.TransactionType, startTime time.Time, endTime time.Time) (result *hashmap.Map[string, interface {
}], err error) {
	eng := this.DBWrite().Table("MemberTransaction as t").Select("sum(t.amount) as amount").Where("t.flag = 0 and t.symbol = ? and t.type = ? and t.createTime >= ? and t.createTime <= ?",symbol,type,startTime,endTime).Find(&result)
	err = eng.Error
	return
}
func (this *memberTransactionDao) FindAllByCreateTime(beginDate string, endDate string) (result arraylist.List[entity.MemberTransaction], err error) {
	err = this.DBRead().Where("create_time = ?",beginDate).Find(&result).Error
	return
}
func (this *memberTransactionDao) Save(m *entity.MemberTransaction) (result *entity.MemberTransaction, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberTransactionDao) FindById(id int64) (result *entity.MemberTransaction, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberTransactionDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberTransaction{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberTransactionDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberTransaction], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}

