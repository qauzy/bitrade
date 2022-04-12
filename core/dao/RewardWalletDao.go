package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

type RewardWalletDao interface {
	FindRewardWalletByMemberId(memberId int64) (result *entity.RewardWallet, err error)
	FindRewardWalletByMemberIdAndCoinUnit(memberId int64, unit string) (result *entity.RewardWallet, err error)
	Save(rewardWallet *entity.RewardWallet) (result *entity.RewardWallet, err error)
	IncreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	DecreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error)
	FindAllByMemberId(memberId int64) (result arraylist.List[entity.RewardWallet], err error)
	FindById(id int64) (result *entity.RewardWallet, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.RewardWallet], err error)
}
type rewardWalletDao struct {
	*db.DB
}

func NewRewardWalletDao(db *db.DB) (dao RewardWalletDao) {
	dao = &rewardWalletDao{db}
	return
}
func (this *rewardWalletDao) FindRewardWalletByMemberId(memberId int64) (result *entity.RewardWallet, err error) {
	err = this.DBRead().Where("member_id = ?", memberId).First(&result).Error
	return
}
func (this *rewardWalletDao) FindRewardWalletByMemberIdAndCoinUnit(memberId int64, unit string) (result *entity.RewardWallet, err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Where("coin_unit = ?", unit).First(&result).Error
	return
}
func (this *rewardWalletDao) Save(rewardWallet *entity.RewardWallet) (result *entity.RewardWallet, err error) {
	return
}
func (this *rewardWalletDao) IncreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance = wallet.balance + ? where wallet.id = ?", walletId, amount)
	err = eng.Error
	return
}
func (this *rewardWalletDao) DecreaseBalance(walletId int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberWallet wallet set wallet.balance = wallet.balance - ? where wallet.id = ? and wallet.balance >= ?", walletId, amount)
	err = eng.Error
	return
}
func (this *rewardWalletDao) FindAllByMemberId(memberId int64) (result arraylist.List[entity.RewardWallet], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Find(&result).Error
	return
}

func (this *rewardWalletDao) FindById(id int64) (result *entity.RewardWallet, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *rewardWalletDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.RewardWallet{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *rewardWalletDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.RewardWallet], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
