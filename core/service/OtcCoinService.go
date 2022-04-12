package service

func (this *OtcCoinService) QueryWhereOrPage(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageNo int, pageSize int) (result pagination.PageResult[entity.OtcCoin], err error) {
	var list arraylist.List[entity.OtcCoin]
	var jpaQuery = queryFactory.SelectFrom(QOtcCoin.OtcCoin)
	if booleanExpressionList != nil {
		jpaQuery.Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *OtcCoinService) Save(otcCoin *entity.OtcCoin) (result *entity.OtcCoin, err error) {
	return this.OtcCoinDao.Save(otcCoin)
}
func (this *OtcCoinService) FindAll() (result arraylist.List[entity.OtcCoin], err error) {
	return this.OtcCoinDao.FindAll()
}
func (this *OtcCoinService) FindOne(id int64) (result *entity.OtcCoin, err error) {
	return this.OtcCoinDao.FindById(id).Get()
}
func (this *OtcCoinService) FindByUnit(unit string) (result *entity.OtcCoin, err error) {
	return this.OtcCoinDao.FindOtcCoinByUnit(unit)
}
func (this *OtcCoinService) GetAllNormalCoin() (result arraylist.List[*hashmap.Map[string, string]], err error) {
	return Model("otc_coin").Field("id,name,name_cn as nameCn,unit,sell_min_amount,sort,buy_min_amount").Where("status=?", CommonStatus.NORMAL.Ordinal()).Order("sort asc").Select()
}
func (this *OtcCoinService) GetNormalCoin() (result arraylist.List[entity.OtcCoin], err error) {
	return this.OtcCoinDao.FindAllByStatus(CommonStatus.NORMAL)
}
func (this *OtcCoinService) PageQuery(pageNo int, pageSize int, name string, nameCn string) (result domain.Page[entity.OtcCoin], err error) {
	//排序方式
	var orders = Criteria.SortStatic("sort")
	//分页参数
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	//查询条件
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Like("name", name, false))
	specification.Add(Restrictions.Like("nameCn", nameCn, false))
	return this.OtcCoinDao.FindAll(specification, pageRequest)
}
func (this *OtcCoinService) Deletes(ids []int64) (err error) {
	for _, id := range ids {
		this.OtcCoinDao.DeleteById(id)
	}
}
func (this *OtcCoinService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.OtcCoin], err error) {
	return this.OtcCoinDao.FindAll(predicate, pageable)
}
func (this *OtcCoinService) FindAllUnits() (result arraylist.List[string], err error) {
	var list = this.OtcCoinDao.FindAllUnits()
	return list
}
func (this *OtcCoinService) FindOtcCoinByUnitAndStatus(coinUnit string, commonStatus *CommonStatus.CommonStatus) (result *entity.OtcCoin, err error) {
	return this.OtcCoinDao.FindOtcCoinByUnitAndStatus(coinUnit, commonStatus)
}
func NewOtcCoinService(otcCoinDao *dao.OtcCoinDao) (ret *OtcCoinService) {
	ret = new(OtcCoinService)
	ret.OtcCoinDao = otcCoinDao
	return
}

type OtcCoinService struct {
	OtcCoinDao dao.OtcCoinDao
	Base.BaseService
}
