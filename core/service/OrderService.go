package service

func (this *OrderService) CancelOrderTask(order *entity.Order) (err error) {
	if order.GetAdvertiseType().Equals(AdvertiseType.BUY) {
		//更改广告
		if !this.AdvertiseService.UpdateAdvertiseAmountForCancel(order.GetAdvertiseId(), order.GetNumber()) {
			return InformationExpiredException("Information Expired")
		}
		//更改钱包
		var memberWallet = this.OtcWalletService.FindByOtcCoinAndMemberId(order.GetCustomerId(), order.GetCoin())
		var result = this.OtcWalletService.ThawBalance(memberWallet, order.GetNumber())
		if result.GetCode() != 0 {
			return InformationExpiredException("Information Expired")
		}
	} else {
		//更改广告
		if !this.AdvertiseService.UpdateAdvertiseAmountForCancel(order.GetAdvertiseId(), add(order.GetNumber(), order.GetCommission())) {
			return InformationExpiredException("Information Expired")
		}
		//更改钱包
		var memberWallet = this.OtcWalletService.FindByOtcCoinAndMemberId(order.GetMemberId(), order.GetCoin())
		var result = this.OtcWalletService.ThawBalance(memberWallet, add(order.GetNumber(), order.GetCommission()))
		if result.GetCode() != 0 {
			return InformationExpiredException("Information Expired")
		}
	}
	//取消订单
	if !(this.CancelNopaymentOrder(order.GetOrderSn()) > 0) {
		return InformationExpiredException("Information Expired")
	}
}
func (this *OrderService) Query(predicateList arraylist.List[types.Predicate], pageNo int, pageSize int) (result pagination.PageResult[entity.Order], err error) {
	var list arraylist.List[entity.Order]
	var jpaQuery = queryFactory.SelectFrom(QOrder.Order)
	if predicateList != nil {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *OrderService) FindOne(id int64) (result *entity.Order, err error) {
	return this.OrderDao.FindById(id).Get()
}
func (this *OrderService) FindOneByOrderSn(orderSn string) (result *entity.Order, err error) {
	return this.OrderDao.GetOrderByOrderSn(orderSn)
}
func (this *OrderService) UpdateOrderAppeal(orderSn string) (result int, err error) {
	return this.OrderDao.UpdateAppealOrder(OrderStatus.APPEAL, orderSn)
}
func (this *OrderService) PayForOrder(orderSn string) (result int, err error) {
	return this.OrderDao.UpdatePayOrder(time.Now(), OrderStatus.PAID, orderSn)
}
func (this *OrderService) CancelOrder(orderSn string) (result int, err error) {
	return this.OrderDao.CancelOrder(time.Now(), OrderStatus.CANCELLED, orderSn)
}
func (this *OrderService) CancelNopaymentOrder(orderSn string) (result int, err error) {
	return this.OrderDao.CancelNopaymentOrder(time.Now(), OrderStatus.CANCELLED, orderSn)
}
func (this *OrderService) ReleaseOrder(orderSn string) (result int, err error) {
	return this.OrderDao.ReleaseOrder(time.Now(), OrderStatus.COMPLETED, orderSn)
}
func (this *OrderService) SaveOrder(order *entity.Order) (result *entity.Order, err error) {
	order.SetOrderSn(String.ValueOf(this.IdWorkByTwitter.NextId()))
	return this.OrderDao.Save(order)
}
func (this *OrderService) PageQuery(pageNo int, pageSize int, status *OrderStatus.OrderStatus, id int64, orderSn string) (result Page[entity.Order], err error) {
	var orders = Criteria.SortStatic("id.desc")
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Or(Restrictions.Eq("memberId", id, false), Restrictions.Eq("customerId", id, false)))
	specification.Add(Restrictions.Eq("status", status, false))
	if StringUtils.IsNotBlank(orderSn) {
		specification.Add(Restrictions.Like("orderSn", orderSn, false))
	}
	return this.OrderDao.FindAll(specification, pageRequest)
}
func (this *OrderService) GetOrderBySn(memberId int64, orderSn string) (result *hashmap.Map[string, interface {
}], err error) {
	var sql = "select o.*,m.real_name from otc_order o  join member m on o.customer_id=m.id and o.member_id=:memberId and o.order_sn =:orderSn "
	var Query = this.Em.CreateNativeQuery(sql)
	//设置结果转成Map类型
	this.Query.Unwrap(hibernate.SQLQuery).SetResultTransformer(Transformers.ALIAS_TO_ENTITY_MAP)
	var object = this.Query.SetParameter("memberId", memberId).SetParameter("orderSn", orderSn).GetSingleResult()
	var oMap = object.(*hashmap.Map[string, interface {
	}])
	return oMap
}
func (this *OrderService) CheckExpiredOrder() (result arraylist.List[entity.Order], err error) {
	return this.OrderDao.FindAllExpiredOrder(time.Now())
}
func (this *OrderService) FindAll() (result arraylist.List[entity.Order], err error) {
	return this.OrderDao.FindAll()
}
func (this *OrderService) Save(order *entity.Order) (result *entity.Order, err error) {
	return this.OrderDao.Save(order)
}
func (this *OrderService) GetOrderNum() (result *util.MessageResult, err error) {
	var predicate = QOrder.Order.Status.Eq(OrderStatus.NONPAYMENT)
	var noPayNum = this.OrderDao.Count(predicate)
	var paidNum = this.OrderDao.Count(QOrder.Order.Status.Eq(OrderStatus.PAID))
	var finishedNum = this.OrderDao.Count(QOrder.Order.Status.Eq(OrderStatus.COMPLETED))
	var cancelNum = this.OrderDao.Count(QOrder.Order.Status.Eq(OrderStatus.CANCELLED))
	var appealNum = this.OrderDao.Count(QOrder.Order.Status.Eq(OrderStatus.APPEAL))
	var oMap = hashset.New[int64]()
	oMap.Put("noPayNum", noPayNum)
	oMap.Put("paidNum", paidNum)
	oMap.Put("finishedNum", finishedNum)
	oMap.Put("cancelNum", cancelNum)
	oMap.Put("appealNum", appealNum)
	return MessageResult.GetSuccessInstance("获取成功", oMap)
}
func (this *OrderService) GetAllOrdering(id int64) (result arraylist.List[entity.Order], err error) {
	return this.OrderDao.FingAllProcessingOrder(id, OrderStatus.APPEAL, OrderStatus.PAID, OrderStatus.NONPAYMENT)
}
func (this *OrderService) FindOneByOrderId(orderId string) (result *entity.Order, err error) {
	return this.OrderDao.GetOrderByOrderSn(orderId)
}
func (this *OrderService) FindAll(predicate *types.Predicate, pageable Pageable) (result Page[entity.Order], err error) {
	return this.OrderDao.FindAll(predicate, pageable)
}
func (this *OrderService) OutExcel(predicates arraylist.List[types.Predicate], pageModel *PageModel.PageModel) (result Page[vo.OtcOrderVO], err error) {
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	var Query = queryFactory.Select(Projections.Fields(vo.OtcOrderVO, QOrder.Order.Id.As("id"), QOrder.Order.OrderSn.As("orderSn"), QOrder.Order.AdvertiseType.As("advertiseType"), QOrder.Order.CreateTime.As("createTime"), QOrder.Order.MemberName.As("memberName"), QOrder.Order.CustomerName.As("customerName"), QOrder.Order.Coin.Unit, QOrder.Order.Money, QOrder.Order.Number, QOrder.Order.Commission.As("fee"), QOrder.Order.PayMode.As("payMode"), QOrder.Order.ReleaseTime.As("releaseTime"), QOrder.Order.CancelTime.As("cancelTime"), QOrder.Order.PayTime.As("payTime"), QOrder.Order.Status.As("status"))).From(QOrder.Order).Where(predicates.ToArray(make([]BooleanExpression, 0)))
	this.Query.OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var list = this.Query.Offset((pageModel.GetPageNo() - 1) * pageModel.GetPageSize()).Limit(pageModel.GetPageSize()).Fetch()
	var total = this.Query.FetchCount()
	return PageImpl(list, pageModel.GetPageable(), total)
}
func (this *OrderService) GetOtcOrderStatistics(date string) (result arraylist.List[[]interface {
}], err error) {
	return this.OrderDao.GetOtcTurnoverAmount(date)
}
func (this *OrderService) CountByMemberProcessing(memberId int64) (result int64, err error) {
	return this.OrderDao.Count(QOrder.Order.Status.Eq(OrderStatus.PAID).Or(QOrder.Order.Status.Eq(OrderStatus.NONPAYMENT)).And(QOrder.Order.MemberId.Eq(memberId).Or(QOrder.Order.CustomerId.Eq(memberId))))
}
func (this *OrderService) CountOrderByMemberIdAndCreateTime(startTime xtime.Xtime, endTime xtime.Xtime) (result int, err error) {
	var objectList = this.OrderDao.CountOrdersByMemberIdAndCreateTime(startTime, endTime)
	if objectList != nil && objectList.Size() > 0 {
		return objectList.Size()
	} else {
		return 0
	}
}
func NewOrderService(em *persistence.EntityManager, orderDao *dao.OrderDao, idWorkByTwitter *util.IdWorkByTwitter, advertiseService AdvertiseService, otcWalletService OtcWalletService) (ret *OrderService) {
	ret = new(OrderService)
	ret.Em = em
	ret.OrderDao = orderDao
	ret.IdWorkByTwitter = idWorkByTwitter
	ret.AdvertiseService = advertiseService
	ret.OtcWalletService = otcWalletService
	return
}

type OrderService struct {
	Em               *persistence.EntityManager
	OrderDao         dao.OrderDao
	IdWorkByTwitter  *util.IdWorkByTwitter
	AdvertiseService *AdvertiseService
	OtcWalletService *OtcWalletService
	Base.BaseService
}
