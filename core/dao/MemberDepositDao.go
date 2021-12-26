package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type MemberDepositDao interface {
	FindByAddressAndTxid(address string, txid string) (result entity.MemberDeposit, err error)
	GetDepositStatistics(date string) (result [][]interface{}, err error)
	Save(m *entity.MemberDeposit) (result *entity.MemberDeposit, err error)
	FindById(id int64) (result *entity.MemberDeposit, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.MemberDeposit, err error)
}
type memberDepositDao struct {
	*db.DB
}

func NewMemberDepositDao(db *db.DB) (dao MemberDepositDao) {
	dao = &memberDepositDao{db}
	return
}
func (this *memberDepositDao) FindByAddressAndTxid(address string, txid string) (result entity.MemberDeposit, err error) {
	err = this.DBRead().Where("address = ?", address).Where("txid = ?", txid).First(&result).Error
	return
}
func (this *memberDepositDao) GetDepositStatistics(date string) (result [][]interface{}, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select unit ,sum(amount) from member_deposit where date_format(create_time,'%Y-%m-%d') = ? group by unit", date)
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
func (this *memberDepositDao) FindAll(qp *types.QueryParam) (result []*entity.MemberDeposit, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
