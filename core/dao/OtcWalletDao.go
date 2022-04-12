package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

type OtcWalletDao interface {
	FindOtcWalletByMemberId(memberId int64) (result arraylist.List[entity.OtcWallet], err error)
	FindByMemberIdAndCoin(memberId int64, coin *entity.Coin) (result *entity.OtcWallet, err error)
	AddWallet(id int64, amount math.BigDecimal) (result int, err error)
	SubWallet(id int64, amount math.BigDecimal) (result int, err error)
	ThawBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	DecreaseFrozen(walletId int64, amount math.BigDecimal) (result int, err error)
	IncreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	FreezeBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	Save(m *entity.OtcWallet) (result *entity.OtcWallet, err error)
	FindById(id int64) (result *entity.OtcWallet, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.OtcWallet], err error)
}
type otcWalletDao struct {
	*db.DB
}

func NewOtcWalletDao(db *db.DB) (dao OtcWalletDao) {
	dao = &otcWalletDao{db}
	return
}
func (this *otcWalletDao) FindOtcWalletByMemberId(memberId int64) (result arraylist.List[entity.OtcWallet], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).First(&result).Error
	return
}
func (this *otcWalletDao) FindByMemberIdAndCoin(memberId int64, coin *entity.Coin) (result *entity.OtcWallet, err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Where("coin = ?", coin).First(&result).Error
	return
}
func (this *otcWalletDao) AddWallet(id int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update OtcWallet set balance = balance+? where id=?", id, amount)
	err = eng.Error
	return
}
func (this *otcWalletDao) SubWallet(id int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update OtcWallet set balance = balance-? where id=? and balance>=?", id, amount)
	err = eng.Error
	return
}
func (this *otcWalletDao) ThawBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update OtcWallet wallet set wallet.balance = wallet.balance + ?,wallet.frozenBalance=wallet.frozenBalance - ? where wallet.id = ? and wallet.frozenBalance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *otcWalletDao) DecreaseFrozen(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update OtcWallet wallet set wallet.frozenBalance=wallet.frozenBalance - ? where wallet.id = ? and wallet.frozenBalance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *otcWalletDao) IncreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update OtcWallet wallet set wallet.balance = wallet.balance + ? where wallet.id = ?", walletId, amount)
	err = eng.Error
	return
}
func (this *otcWalletDao) FreezeBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update OtcWallet wallet set wallet.balance = wallet.balance - ?,wallet.frozenBalance=wallet.frozenBalance + ? where wallet.id = ? and wallet.balance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *otcWalletDao) Save(m *entity.OtcWallet) (result *entity.OtcWallet, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *otcWalletDao) FindById(id int64) (result *entity.OtcWallet, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *otcWalletDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.OtcWallet{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *otcWalletDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.OtcWallet], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
