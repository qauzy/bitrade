package dao

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/dto"
	"bitrade/core/entity"
	"github.com/qauzy/math"
)

type CoinDao interface {
	FindByUnit(unit string) (result entity.Coin, err error)
	FindAllByCanWithdrawAndStatusAndHasLegal(is BooleanEnum.BooleanEnum, status CommonStatus.CommonStatus, hasLegal bool) (result []entity.Coin, err error)
	FindCoinByIsPlatformCoin(is BooleanEnum.BooleanEnum) (result entity.Coin, err error)
	FindByHasLegal(hasLegal bool) (result []entity.Coin, err error)
	FindAllByOtc(otcUnits []string) (result []entity.Coin, err error)
	FindAllName() (result []string, err error)
	FindAllNameAndUnit() (result []dto.CoinDTO, err error)
	FindAllCoinNameLegal() (result []string, err error)
	FindAllRpcUnit() (result []string, err error)
	FindAllByIsPlatformCoin(isPlatformCoin BooleanEnum.BooleanEnum) (result []entity.Coin, err error)
	SumBalance(coin entity.Coin) (result math.BigDecimal, err error)
	GetBalanceByMemberIdAndCoinId(coin entity.Coin, memberId int64) (result math.BigDecimal, err error)
	FindAllOrderBySort() (result []entity.Coin, err error)
	FindAllByStatus(status CommonStatus.CommonStatus) (result []entity.Coin, err error)
	FindByStatus(status CommonStatus.CommonStatus) (result []entity.Coin, err error)
	Save(m *entity.Coin) (result *entity.Coin, err error)
	FindById(id int64) (result *entity.Coin, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Coin, err error)
}
type coinDao struct {
	*db.DB
}

func NewCoinDao(db *db.DB) (dao CoinDao) {
	dao = &coinDao{db}
	return
}
func (this *coinDao) FindByUnit(unit string) (result entity.Coin, err error) {
	err = this.DBRead().Where("unit = ?", unit).First(&result).Error
	return
}
func (this *coinDao) FindAllByCanWithdrawAndStatusAndHasLegal(is BooleanEnum.BooleanEnum, status CommonStatus.CommonStatus, hasLegal bool) (result []entity.Coin, err error) {
	err = this.DBRead().Where("can_withdraw = ?", is).Where("status = ?", status).Where("has_legal = ?", hasLegal).Find(&result).Error
	return
}
func (this *coinDao) FindCoinByIsPlatformCoin(is BooleanEnum.BooleanEnum) (result entity.Coin, err error) {
	err = this.DBRead().Where("is_platform_coin = ?", is).First(&result).Error
	return
}
func (this *coinDao) FindByHasLegal(hasLegal bool) (result []entity.Coin, err error) {
	err = this.DBRead().Where("has_legal = ?", hasLegal).First(&result).Error
	return
}
func (this *coinDao) FindAllByOtc(otcUnits []string) (result []entity.Coin, err error) {
	err = this.DBRead().Where("otc = ?", otcUnits).Find(&result).Error
	return
}
func (this *coinDao) FindAllName() (result []string, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select a.name from Coin a")
	err = eng.Error
	return
}
func (this *coinDao) FindAllNameAndUnit() (result []dto.CoinDTO, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select  new cn.ztuo.bitrade.dto.CoinDTO(a.name,a.unit) from Coin a")
	err = eng.Error
	return
}
func (this *coinDao) FindAllCoinNameLegal() (result []string, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select a.name from Coin a where a.hasLegal = true ")
	err = eng.Error
	return
}
func (this *coinDao) FindAllRpcUnit() (result []string, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select a.unit from Coin a where a.enableRpc = 1")
	err = eng.Error
	return
}
func (this *coinDao) FindAllByIsPlatformCoin(isPlatformCoin BooleanEnum.BooleanEnum) (result []entity.Coin, err error) {
	err = this.DBRead().Where("is_platform_coin = ?", isPlatformCoin).Find(&result).Error
	return
}
func (this *coinDao) SumBalance(coin entity.Coin) (result math.BigDecimal, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT sum(a.balance) FROM MemberWallet a WHERE a.coin = ?", coin)
	err = eng.Error
	return
}
func (this *coinDao) GetBalanceByMemberIdAndCoinId(coin entity.Coin, memberId int64) (result math.BigDecimal, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT a.balance FROM MemberWallet a WHERE a.coin = ? AND a.memberId = ?", coin, memberId)
	err = eng.Error
	return
}
func (this *coinDao) FindAllOrderBySort() (result []entity.Coin, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select a from Coin a order by a.sort")
	err = eng.Error
	return
}
func (this *coinDao) FindAllByStatus(status CommonStatus.CommonStatus) (result []entity.Coin, err error) {
	err = this.DBRead().Where("status = ?", status).Find(&result).Error
	return
}
func (this *coinDao) FindByStatus(status CommonStatus.CommonStatus) (result []entity.Coin, err error) {
	err = this.DBRead().Where("status = ?", status).First(&result).Error
	return
}
func (this *coinDao) Save(m *entity.Coin) (result *entity.Coin, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *coinDao) FindById(id int64) (result *entity.Coin, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *coinDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Coin{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *coinDao) FindAll(qp *types.QueryParam) (result []*entity.Coin, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
