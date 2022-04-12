package service

func (this *TestQueryDSLService) InitFactory() (err error) {
	this.QueryFactory = JPAQueryFactory(this.EntityManager)
}
func (this *TestQueryDSLService) Test(page int) (result domain.Page[entity.Advertise], err error) {
	//动态条件
	var qAdvertise = QAdvertise.Advertise
	var predicate = qAdvertise.Id.Gt(1)
	//分页排序
	var orders = Criteria.SortStatic("id")
	var pageRequest = PageRequest.Of(page-1, 10, orders)
	//Page<Advertise> all = advertiseDao.findAll(predicate, pageRequest);
	return this.AdvertiseDao.FindAll(predicate, pageRequest)
}
func (this *TestQueryDSLService) Test2() (result arraylist.List[entity.Advertise], err error) {
	var quser = QAdvertise.Advertise
	return this.QueryFactory.SelectFrom(quser).Fetch()
}
func NewTestQueryDSLService(advertiseDao *dao.AdvertiseDao, entityManager *persistence.EntityManager, queryFactory *impl.JPAQueryFactory) (ret *TestQueryDSLService) {
	ret = new(TestQueryDSLService)
	ret.AdvertiseDao = advertiseDao
	ret.EntityManager = entityManager
	ret.QueryFactory = queryFactory
	return
}

type TestQueryDSLService struct {
	AdvertiseDao  dao.AdvertiseDao
	EntityManager *persistence.EntityManager
	QueryFactory  *impl.JPAQueryFactory
	Base.BaseService
}
