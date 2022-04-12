package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type MemberDepositDao interface {
	FindByAddressAndTxid(address string, txid string) (result *entity.MemberDeposit, err error)
	GetDepositStatistics(date string) (result arraylist.List[[]interface {
	}], err error)
	Save(m *entity.MemberDeposit) (result *entity.MemberDeposit, err error)
	FindById(id int64) (result *entity.MemberDeposit, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberDeposit], err error)
}
type memberDepositDao struct {
	*db.DB
}

func NewMemberDepositDao(db *db.DB) (dao MemberDepositDao) {
	dao = &memberDepositDao{db}
	return
}
func (this *memberDepositDao) FindByAddressAndTxid(address string, txid string) (result *entity.MemberDeposit, err error) {
	err = this.DBRead().Where("address = ?", address).Where("txid = ?", txid).First(&result).Error
	return
}
func (this *memberDepositDao) GetDepositStatistics(date string) (result arraylist.List[[]interface {
}], err error) {
	eng := this.DBWrite().Table("member_deposit").Select("unit, sum(amount)").Where("date_format(create_time, '%Y-%m-%d') = ?", date).Group(" group by unit").Find(&result)
	err = eng.Error
	return
}
func (this *memberDepositDao) Save(m *entity.MemberDeposit) (result *entity.MemberDeposit, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberDepositDao) FindById(id int64) (result *entity.MemberDeposit, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberDepositDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberDeposit{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberDepositDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberDeposit], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
