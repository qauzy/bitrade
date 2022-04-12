package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

type MemberWalletDao interface {
	IncreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	DecreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	ThawBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	FreezeBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	DecreaseFrozen(walletId int64, amount math.BigDecimal) (result int, err error)
	FindByCoinAndAddress(coin *entity.Coin, address string) (result *entity.MemberWallet, err error)
	FindByCoinAndMemberId(coin *entity.Coin, memberId int64) (result *entity.MemberWallet, err error)
	FindAllByMemberId(memberId int64) (result arraylist.List[entity.MemberWallet], err error)
	FindAllByCoin(coin *entity.Coin) (result arraylist.List[entity.MemberWallet], err error)
	GetWalletAllBalance(coinName string) (result math.BigDecimal, err error)
	FindALLByCoinIdAndMemberIdList(coin *entity.Coin, memberIdList arraylist.List[int64]) (result arraylist.List[entity.MemberWallet], err error)
	FindCoinIdAndMemberId(coinId string, memberId int64) (result *entity.MemberWallet, err error)
	CreateGiftTable(times int64, coinId string) (result int, err error)
	FindGiftTable(times int64) (result arraylist.List[entity.MemberWallet], err error)
	SumGiftTable(times int64) (result math.BigDecimal, err error)
	ReleaseReisterGiving(amount math.BigDecimal, id int64) (result int, err error)
	Save(m *entity.MemberWallet) (result *entity.MemberWallet, err error)
	FindById(id int64) (result *entity.MemberWallet, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberWallet], err error)
}
type memberWalletDao struct {
	*db.DB
}

func NewMemberWalletDao(db *db.DB) (dao MemberWalletDao) {
	dao = &memberWalletDao{db}
	return
}
func (this *memberWalletDao) IncreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance = wallet.balance + ? where wallet.id = ?", walletId, amount)
	err = eng.Error
	return
}
func (this *memberWalletDao) DecreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance = wallet.balance - ? where wallet.id = ? and wallet.balance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *memberWalletDao) ThawBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance = wallet.balance + ?,wallet.frozenBalance=wallet.frozenBalance - ? where wallet.id = ? and wallet.frozenBalance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *memberWalletDao) FreezeBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance = wallet.balance - ?,wallet.frozenBalance=wallet.frozenBalance + ? where wallet.id = ? and wallet.balance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *memberWalletDao) DecreaseFrozen(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.frozenBalance=wallet.frozenBalance - ? where wallet.id = ? and wallet.frozenBalance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *memberWalletDao) FindByCoinAndAddress(coin *entity.Coin, address string) (result *entity.MemberWallet, err error) {
	err = this.DBRead().Where("coin = ?", coin).Where("address = ?", address).First(&result).Error
	return
}
func (this *memberWalletDao) FindByCoinAndMemberId(coin *entity.Coin, memberId int64) (result *entity.MemberWallet, err error) {
	err = this.DBRead().Where("coin = ?", coin).Where("member_id = ?", memberId).First(&result).Error
	return
}
func (this *memberWalletDao) FindAllByMemberId(memberId int64) (result arraylist.List[entity.MemberWallet], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Find(&result).Error
	return
}
func (this *memberWalletDao) FindAllByCoin(coin *entity.Coin) (result arraylist.List[entity.MemberWallet], err error) {
	err = this.DBRead().Where("coin = ?", coin).Find(&result).Error
	return
}
func (this *memberWalletDao) GetWalletAllBalance(coinName string) (result math.BigDecimal, err error) {
	eng := this.DBWrite().Table("member_wallet as a").Select("ifnull(sum(a.balance), 0) + ifnull(sum(a.frozen_balance), 0) as allBalance").Where("a.coin_id = ?", coinName).Find(&result)
	err = eng.Error
	return
}
func (this *memberWalletDao) FindALLByCoinIdAndMemberIdList(coin *entity.Coin, memberIdList arraylist.List[int64]) (result arraylist.List[entity.MemberWallet], err error) {
	err = this.DBRead().Where("coin_id = ?", coin).Where("member_id_list = ?", memberIdList).First(&result).Error
	return
}
func (this *memberWalletDao) FindCoinIdAndMemberId(coinId string, memberId int64) (result *entity.MemberWallet, err error) {
	eng := this.DBWrite().Table("member_wallet").Select("*").Where("coin_id = ? and member_id = ?", coinId, memberId).Find(&result)
	err = eng.Error
	return
}
func (this *memberWalletDao) CreateGiftTable(times int64, coinId string) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("CREATE TABLE member_wallet_? SELECT * FROM member_wallet where coin_id=? AND balance>0 AND is_lock=0", times, coinId)
	err = eng.Error
	return
}
func (this *memberWalletDao) FindGiftTable(times int64) (result arraylist.List[entity.MemberWallet], err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("SELECT * FROM member_wallet_? ", times)
	err = eng.Error
	return
}
func (this *memberWalletDao) SumGiftTable(times int64) (result math.BigDecimal, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("SELECT sum(balance) FROM member_wallet_? ", times)
	err = eng.Error
	return
}
func (this *memberWalletDao) ReleaseReisterGiving(amount math.BigDecimal, id int64) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance=wallet.balance + ?,wallet.releaseBalance= wallet.releaseBalance -? where wallet.id = :walletId and wallet.releaseBalance >= ?", amount, id)
	err = eng.Error
	return
}
func (this *memberWalletDao) Save(m *entity.MemberWallet) (result *entity.MemberWallet, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberWalletDao) FindById(id int64) (result *entity.MemberWallet, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberWalletDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberWallet{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberWalletDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberWallet], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
