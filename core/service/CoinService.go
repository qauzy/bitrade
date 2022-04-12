package service

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/PageModel"
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/dto"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

func (this *CoinService) Query(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageNo int, pageSize int) (result pagination.PageResult[entity.Coin], err error) {
	var list arraylist.List[entity.Coin]
	var jpaQuery = queryFactory.SelectFrom(QCoin.Coin)
	if booleanExpressionList != nil {
		jpaQuery.Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
	//添加总条数
}
func (this *CoinService) FindOne(name string) (result *entity.Coin, err error) {
	return this.CoinDao.FindByUnit(name)
}
func (this *CoinService) FindByUnit(unit string) (result *entity.Coin, err error) {
	return this.CoinDao.FindByUnit(unit)
}
func (this *CoinService) Save(coin *entity.Coin) (result *entity.Coin, err error) {
	return this.CoinDao.Save(coin)
}
func (this *CoinService) FindAll() (result arraylist.List[entity.Coin], err error) {
	return this.CoinDao.FindAllOrderBySort()
}
func (this *CoinService) FindAllCoinByOtc() (result arraylist.List[entity.Coin], err error) {
	var supportUnits = this.OtcCoinDao.FindAll().Stream().Map(func(x interface {
	}) {
		x.GetUnit()
	}).Collect(Collectors.ToList())
	if supportUnits.Size() > 0 {
		return this.CoinDao.FindAllByOtc(supportUnits)
	} else {
		return nil
	}
}
func (this *CoinService) PageQuery(pageNo int, pageSize int) (result domain.Page[entity.Coin], err error) {
	//排序方式 (需要倒序 这样    Criteria.sort("id","createTime.desc") ) //参数实体类为字段名
	var orders = Criteria.SortStatic("sort")
	//分页参数
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	//查询条件
	var specification = new(pagination.Criteria)
	return this.CoinDao.FindAll(specification, pageRequest)
}
func (this *CoinService) FindAllCanWithDraw() (result arraylist.List[entity.Coin], err error) {
	return this.CoinDao.FindAllByCanWithdrawAndStatusAndHasLegal(IS_TRUE, CommonStatus.NORMAL, false)
}
func (this *CoinService) DeleteOne(name string) (err error) {
	this.CoinDao.DeleteById(name)
}
func (this *CoinService) SetPlatformCoin(coin *entity.Coin) (err error) {
	var list = this.CoinDao.FindAll()
	list.Stream().Filter(func(x *entity.Coin) bool {
		!x.GetName().Equals(coin.GetName())
	}).ForEach(func(x *entity.Coin) {
		x.SetIsPlatformCoin(BooleanEnum.IS_FALSE)
		this.CoinDao.Save(x)
	})
	coin.SetIsPlatformCoin(IS_TRUE)
	this.CoinDao.SaveAndFlush(coin)
	var otcCoin = this.OtcCoinDao.FindOtcCoinByUnit(coin.GetUnit())
	if otcCoin != nil {
		otcCoin.SetIsPlatformCoin(IS_TRUE)
		this.OtcCoinDao.SaveAndFlush(otcCoin)
	}
	var list1 = this.OtcCoinDao.FindAll()
	list1.Stream().Filter(func(x *entity.Coin) {
		!x.GetUnit().Equals(coin.GetUnit())
	}).ForEach(func(x *entity.Coin) {
		x.SetIsPlatformCoin(BooleanEnum.IS_FALSE)
		this.OtcCoinDao.Save(x)
	})
}
func (this *CoinService) QueryPlatformCoin() (result *entity.Coin, err error) {
	return this.CoinDao.FindCoinByIsPlatformCoin(IS_TRUE)
}
func (this *CoinService) FindLegalAll() (result arraylist.List[entity.Coin], err error) {
	return this.CoinDao.FindAll(QCoin.Coin.HasLegal.Eq(true)).(arraylist.List[entity.Coin])
}
func (this *CoinService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.Coin], err error) {
	return this.CoinDao.FindAll(predicate, pageable)
}
func (this *CoinService) FindLegalCoinPage(pageModel *PageModel.PageModel) (result *domain.Page, err error) {
	var eq = QCoin.Coin.HasLegal.Eq(true)
	return this.CoinDao.FindAll(eq, pageModel.GetPageable())
}
func (this *CoinService) GetAllCoinName() (result arraylist.List[string], err error) {
	var list = this.CoinDao.FindAllName()
	return list
}
func (this *CoinService) GetAllCoinNameAndUnit() (result arraylist.List[dto.CoinDTO], err error) {
	var allNameAndUnit = this.CoinDao.FindAllNameAndUnit()
	return allNameAndUnit
}
func (this *CoinService) GetAllCoinNameLegal() (result arraylist.List[string], err error) {
	return this.CoinDao.FindAllCoinNameLegal()
}
func (this *CoinService) FindAllRpcUnit() (result arraylist.List[string], err error) {
	return this.CoinDao.FindAllRpcUnit()
}
func (this *CoinService) SetPlatform(coin *entity.Coin) (err error) {
	//取消其他平台币
	var coins = this.CoinDao.FindAllByIsPlatformCoin(IS_TRUE)
	coins.ForEach(func(x *entity.Coin) {
		x.SetIsPlatformCoin(IS_FALSE)
	})
	//设置传入币为平台币
	coin.SetIsPlatformCoin(IS_TRUE)
}
func (this *CoinService) FindAllByStatus(status *CommonStatus.CommonStatus) (result arraylist.List[entity.Coin], err error) {
	return this.CoinDao.FindAllByStatus(status)
}
func (this *CoinService) SumBalance(coin *entity.Coin) (result math.BigDecimal, err error) {
	return this.CoinDao.SumBalance(coin)
}
func (this *CoinService) GetBalanceByMemberIdAndCoinId(coin *entity.Coin, memberId int64) (result math.BigDecimal, err error) {
	return this.CoinDao.GetBalanceByMemberIdAndCoinId(coin, memberId)
}
func (this *CoinService) FindByStatus(status *CommonStatus.CommonStatus) (result arraylist.List[entity.Coin], err error) {
	return this.CoinDao.FindByStatus(status)
}
func NewCoinService(coinDao dao.CoinDao, otcCoinDao dao.OtcCoinDao) (ret *CoinService) {
	ret = new(CoinService)
	ret.CoinDao = coinDao
	ret.OtcCoinDao = otcCoinDao
	return
}

type CoinService struct {
	CoinDao    dao.CoinDao
	OtcCoinDao dao.OtcCoinDao
	Base.BaseService
}
